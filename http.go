package fun

import (
	"compress/flate"
	"compress/gzip"
	"crypto/tls"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	HttpDefaultTimeOut        = 10000
	HttpDefaultUserAgent      = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36"
	HttpDefaultAcceptEncoding = "gzip, deflate"
)

type HttpReq struct {
	// UserAgent 优先于请求头 Headers 中的 User-Agent 字段
	UserAgent string

	// 请求头
	Headers map[string]string

	// 限制最大返回大小
	MaxContentLength int64

	// 限制允许访问 ContentType 列表, 前缀匹配
	AllowedContentTypes []string

	// 最大 Redirect 次数, 范围 [0,10), 否则采用默认的跳转策略 (最大 10 次)
	MaxRedirect int

	// 禁止跳转
	DisableRedirect bool

	// http.Transport
	Transport http.RoundTripper
}

type HttpResp struct {
	// 是否成功 (200-299), 成功仍可能返回 err
	Success bool

	// Http 状态码
	StatusCode int

	// 响应体
	Body []byte

	// ContentLength (字节数)
	ContentLength int64

	// 响应头
	Headers *http.Header
}

// HttpDefaultTransport 默认全局使用的 http.Transport
var HttpDefaultTransport = &http.Transport{
	DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
	ForceAttemptHTTP2:     true,
	MaxIdleConns:          150,
	MaxIdleConnsPerHost:   3,
	IdleConnTimeout:       60 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
	TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
}

// HttpGet 参数为请求地址 (HttpReq, 超时时间)
// HttpGet(url)、HttpGet(url, HttpReq)、HttpGet(url, timeout)、HttpGet(url, HttpReq, timeout)
// 返回 body, 错误信息
func HttpGet(urlStr string, args ...any) ([]byte, error) {
	l := len(args)

	switch l {
	case 0:
		return HttpGetDo(urlStr, nil, 0)
	case 1:
		switch v := args[0].(type) {
		case int:
			timeout := ToInt(args[0])
			return HttpGetDo(urlStr, nil, timeout)
		case *HttpReq:
			return HttpGetDo(urlStr, v, 0)
		}
	case 2:
		timeout := ToInt(args[1])
		switch v := args[0].(type) {
		case *HttpReq:
			return HttpGetDo(urlStr, v, timeout)
		}
	}

	return nil, errors.New("http get params error")
}

// HttpDelete 参数为请求地址 (HttpReq, 超时时间)
// HttpDelete(url)、HttpDelete(url, HttpReq)、HttpDelete(url, timeout)、HttpDelete(url, HttpReq, timeout)
// 返回 body, 错误信息
func HttpDelete(urlStr string, args ...any) ([]byte, error) {
	l := len(args)

	switch l {
	case 0:
		return HttpDeleteDo(urlStr, nil, 0)
	case 1:
		switch v := args[0].(type) {
		case int:
			timeout := ToInt(args[0])
			return HttpDeleteDo(urlStr, nil, timeout)
		case *HttpReq:
			return HttpDeleteDo(urlStr, v, 0)
		}
	case 2:
		timeout := ToInt(args[1])
		switch v := args[0].(type) {
		case *HttpReq:
			return HttpDeleteDo(urlStr, v, timeout)
		}
	}

	return nil, errors.New("http delete params error")
}

// HttpPost 参数为请求地址 (body io.Reader, HttpReq, 超时时间)
// HttpPost(url)、HttpPost(url, timeout)、HttpPost(url, body)、HttpPost(url, body, timeout)、HttpPost(url, body, HttpReq)、HttpPostForm(url, body, HttpReq, timeout)
// 返回 body, 错误信息
func HttpPost(urlStr string, args ...any) ([]byte, error) {
	l := len(args)

	switch l {
	case 0:
		return HttpPostDo(urlStr, nil, nil, 0)
	case 1:
		switch v := args[0].(type) {
		case int:
			timeout := ToInt(args[0])
			return HttpPostDo(urlStr, nil, nil, timeout)
		case io.Reader:
			return HttpPostDo(urlStr, v, nil, 0)
		}
	case 2:
		switch v := args[0].(type) {
		case io.Reader:
			switch h := args[1].(type) {
			case int:
				timeout := ToInt(args[1])
				return HttpPostDo(urlStr, v, nil, timeout)
			case *HttpReq:
				return HttpPostDo(urlStr, v, h, 0)
			}
		}
	case 3:
		switch v := args[0].(type) {
		case io.Reader:
			switch h := args[1].(type) {
			case *HttpReq:
				timeout := ToInt(args[2])
				return HttpPostDo(urlStr, v, h, timeout)
			}
		}
	}

	return nil, errors.New("http post params error")
}

// HttpPostForm 参数为请求地址 (Form 数据 map[string]string, HttpReq, 超时时间)
// HttpPostForm(url)、HttpPostForm(url, timeout)、HttpPostForm(url, posts)、HttpPostForm(url, posts, timeout)、HttpPostForm(url, posts, HttpReq)、HttpPostForm(url, posts, HttpReq, timeout)
// 返回 body, 错误信息
func HttpPostForm(urlStr string, args ...any) ([]byte, error) {
	l := len(args)

	switch l {
	case 0:
		return HttpPostFormDo(urlStr, nil, nil, 0)
	case 1:
		switch v := args[0].(type) {
		case int:
			timeout := ToInt(args[0])
			return HttpPostFormDo(urlStr, nil, nil, timeout)
		case map[string]string:
			return HttpPostFormDo(urlStr, v, nil, 0)
		}
	case 2:
		switch v := args[0].(type) {
		case map[string]string:
			switch h := args[1].(type) {
			case int:
				timeout := ToInt(args[1])
				return HttpPostFormDo(urlStr, v, nil, timeout)
			case *HttpReq:
				return HttpPostFormDo(urlStr, v, h, 0)
			}
		}
	case 3:
		switch v := args[0].(type) {
		case map[string]string:
			switch h := args[1].(type) {
			case *HttpReq:
				timeout := ToInt(args[2])
				return HttpPostFormDo(urlStr, v, h, timeout)
			}
		}
	}

	return nil, errors.New("http post Form params error")
}

// HttpPostJson 参数为请求地址 (Json 数据 string, HttpReq, 超时时间)
// HttpPostJson(url)、HttpPostJson(url, timeout)、HttpPostJson(url, json)、HttpPost(url, json, timeout)、HttpPost(url, json, HttpReq)、HttpPost(url, json, HttpReq, timeout)
// 返回 body, 错误信息
func HttpPostJson(urlStr string, args ...any) ([]byte, error) {
	l := len(args)
	switch l {
	case 0:
		return HttpPostJsonDo(urlStr, "{}", nil, 0)
	case 1:
		switch v := args[0].(type) {
		case int:
			timeout := ToInt(args[0])
			return HttpPostJsonDo(urlStr, "{}", nil, timeout)
		case string:
			return HttpPostJsonDo(urlStr, v, nil, 0)
		}
	case 2:
		switch v := args[0].(type) {
		case string:
			switch h := args[1].(type) {
			case int:
				timeout := ToInt(args[1])
				return HttpPostJsonDo(urlStr, v, nil, timeout)
			case *HttpReq:
				return HttpPostJsonDo(urlStr, v, h, 0)
			}
		}
	case 3:
		switch v := args[0].(type) {
		case string:
			switch h := args[1].(type) {
			case *HttpReq:
				timeout := ToInt(args[2])
				return HttpPostJsonDo(urlStr, v, h, timeout)
			}
		}
	}

	return nil, errors.New("http post json params error")
}

// HttpPut 参数为请求地址 (body io.Reader, HttpReq, 超时时间)
// HttpPut(url)、HttpPut(url, timeout)、HttpPut(url, body)、HttpPut(url, body, timeout)、HttpPut(url, body, HttpReq)、HttpPut(url, body, HttpReq, timeout)
// 返回 body, 错误信息
func HttpPut(urlStr string, args ...any) ([]byte, error) {
	l := len(args)

	switch l {
	case 0:
		return HttpPutDo(urlStr, nil, nil, 0)
	case 1:
		switch v := args[0].(type) {
		case int:
			timeout := ToInt(args[0])
			return HttpPutDo(urlStr, nil, nil, timeout)
		case io.Reader:
			return HttpPutDo(urlStr, v, nil, 0)
		}
	case 2:
		switch v := args[0].(type) {
		case io.Reader:
			switch h := args[1].(type) {
			case int:
				timeout := ToInt(args[1])
				return HttpPutDo(urlStr, v, nil, timeout)
			case *HttpReq:
				return HttpPutDo(urlStr, v, h, 0)
			}
		}
	case 3:
		switch v := args[0].(type) {
		case io.Reader:
			switch h := args[1].(type) {
			case *HttpReq:
				timeout := ToInt(args[2])
				return HttpPutDo(urlStr, v, h, timeout)
			}
		}
	}

	return nil, errors.New("http post params error")
}

// HttpPutForm 参数为请求地址 (Form 数据 map[string]string, HttpReq, 超时时间)
// HttpPutForm(url)、HttpPutForm(url, timeout)、HttpPutForm(url, posts)、HttpPutForm(url, posts, timeout)、HttpPutForm(url, posts, HttpReq)、HttpPutForm(url, posts, HttpReq, timeout)
// 返回 body, 错误信息
func HttpPutForm(urlStr string, args ...any) ([]byte, error) {
	l := len(args)

	switch l {
	case 0:
		return HttpPutFormDo(urlStr, nil, nil, 0)
	case 1:
		switch v := args[0].(type) {
		case int:
			timeout := ToInt(args[0])
			return HttpPutFormDo(urlStr, nil, nil, timeout)
		case map[string]string:
			return HttpPutFormDo(urlStr, v, nil, 0)
		}
	case 2:
		switch v := args[0].(type) {
		case map[string]string:
			switch h := args[1].(type) {
			case int:
				timeout := ToInt(args[1])
				return HttpPutFormDo(urlStr, v, nil, timeout)
			case *HttpReq:
				return HttpPutFormDo(urlStr, v, h, 0)
			}
		}
	case 3:
		switch v := args[0].(type) {
		case map[string]string:
			switch h := args[1].(type) {
			case *HttpReq:
				timeout := ToInt(args[2])
				return HttpPutFormDo(urlStr, v, h, timeout)
			}
		}
	}

	return nil, errors.New("http post form params error")
}

// HttpPutJson 参数为请求地址 (Json 数据 string, HttpReq, 超时时间)
// HttpPutJson(url)、HttpPutJson(url, timeout)、HttpPutJson(url, json)、HttpPutJson(url, json, timeout)、HttpPutJson(url, json, httpReq)、HttpPutJson(url, json, httpReq, timeout)
// 返回 body, 错误信息
func HttpPutJson(urlStr string, args ...any) ([]byte, error) {
	l := len(args)
	switch l {
	case 0:
		return HttpPutJsonDo(urlStr, "{}", nil, 0)
	case 1:
		switch v := args[0].(type) {
		case int:
			timeout := ToInt(args[0])
			return HttpPutJsonDo(urlStr, "{}", nil, timeout)
		case string:
			return HttpPutJsonDo(urlStr, v, nil, 0)
		}
	case 2:
		switch v := args[0].(type) {
		case string:
			switch h := args[1].(type) {
			case int:
				timeout := ToInt(args[1])
				return HttpPutJsonDo(urlStr, v, nil, timeout)
			case *HttpReq:
				return HttpPutJsonDo(urlStr, v, h, 0)
			}
		}
	case 3:
		switch v := args[0].(type) {
		case string:
			switch h := args[1].(type) {
			case *HttpReq:
				timeout := ToInt(args[2])
				return HttpPutJsonDo(urlStr, v, h, timeout)
			}
		}
	}

	return nil, errors.New("http post json params error")
}

// HttpGetDo Http Get 请求, 参数为请求地址, HttpReq, 超时时间(毫秒)
// 返回 body, 错误信息
func HttpGetDo(urlStr string, r *HttpReq, timeout int) ([]byte, error) {
	resp, err := HttpGetResp(urlStr, r, timeout)
	if err != nil {
		return nil, err
	} else {
		return resp.Body, nil
	}
}

// HttpDeleteDo Http Delete 请求, 参数为请求地址, HttpReq, 超时时间(毫秒)
// 返回 body, 错误信息
func HttpDeleteDo(urlStr string, r *HttpReq, timeout int) ([]byte, error) {
	resp, err := HttpDeleteResp(urlStr, r, timeout)
	if err != nil {
		return nil, err
	} else {
		return resp.Body, nil
	}
}

// HttpPostDo Http Post, 参数为请求地址, body io.Reader, HttpReq, 超时时间(毫秒)
// 返回 body, 错误信息
func HttpPostDo(urlStr string, body io.Reader, r *HttpReq, timeout int) ([]byte, error) {
	resp, err := HttpPostResp(urlStr, body, r, timeout)
	if err != nil {
		return nil, err
	} else {
		return resp.Body, nil
	}
}

// HttpPostFormDo Http Post Form, 参数为请求地址, Form 数据 map[string]string, HttpReq, 超时时间(毫秒)
// 返回 body, 错误信息
func HttpPostFormDo(urlStr string, posts map[string]string, r *HttpReq, timeout int) ([]byte, error) {
	resp, err := HttpPostFormResp(urlStr, posts, r, timeout)
	if err != nil {
		return nil, err
	} else {
		return resp.Body, nil
	}
}

// HttpPostJsonDo Http Post Json 请求, 参数为请求地址, Json 数据 string, HttpReq, 超时时间(毫秒)
// 返回 body, 错误信息
func HttpPostJsonDo(urlStr string, json string, r *HttpReq, timeout int) ([]byte, error) {
	resp, err := HttpPostJsonResp(urlStr, json, r, timeout)
	if err != nil {
		return nil, err
	} else {
		return resp.Body, nil
	}
}

// HttpPutDo Http Put, 参数为请求地址, body io.Reader, HttpReq, 超时时间(毫秒)
// 返回 body, 错误信息
func HttpPutDo(urlStr string, body io.Reader, r *HttpReq, timeout int) ([]byte, error) {
	resp, err := HttpPutResp(urlStr, body, r, timeout)
	if err != nil {
		return nil, err
	} else {
		return resp.Body, nil
	}
}

// HttpPutFormDo Http Put Form, 参数为请求地址, Form 数据 map[string]string, HttpReq, 超时时间(毫秒)
// 返回 body, 错误信息
func HttpPutFormDo(urlStr string, posts map[string]string, r *HttpReq, timeout int) ([]byte, error) {
	resp, err := HttpPutFormResp(urlStr, posts, r, timeout)
	if err != nil {
		return nil, err
	} else {
		return resp.Body, nil
	}
}

// HttpPutJsonDo Http Put Json 请求, 参数为请求地址, Json 数据 string, HttpReq, 超时时间(毫秒)
// 返回 body, 错误信息
func HttpPutJsonDo(urlStr string, json string, r *HttpReq, timeout int) ([]byte, error) {
	resp, err := HttpPutJsonResp(urlStr, json, r, timeout)
	if err != nil {
		return nil, err
	} else {
		return resp.Body, nil
	}
}

// HttpGetResp Http Get 请求, 参数为请求地址, HttpReq, 超时时间(毫秒)
// 返回 HttpResp, 错误信息
func HttpGetResp(urlStr string, r *HttpReq, timeout int) (*HttpResp, error) {
	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}

	return HttpDoResp(req, r, timeout)
}

// HttpDeleteResp Http Delete 请求, 参数为请求地址, HttpReq, 超时时间(毫秒)
// 返回 HttpResp, 错误信息
func HttpDeleteResp(urlStr string, r *HttpReq, timeout int) (*HttpResp, error) {
	req, err := http.NewRequest(http.MethodDelete, urlStr, nil)
	if err != nil {
		return nil, err
	}

	return HttpDoResp(req, r, timeout)
}

// HttpPostResp Http Post, 参数为请求地址, body io.Reader, HttpReq, 超时时间(毫秒)
// 返回 HttpResp, 错误信息
func HttpPostResp(urlStr string, body io.Reader, r *HttpReq, timeout int) (*HttpResp, error) {
	req, err := http.NewRequest(http.MethodPost, urlStr, body)
	if err != nil {
		return nil, err
	}

	return HttpDoResp(req, r, timeout)
}

// HttpPostFormResp Http Post Form, 参数为请求地址, Form 数据 map[string]string, HttpReq, 超时时间(毫秒)
// 返回 HttpResp, 错误信息
func HttpPostFormResp(urlStr string, posts map[string]string, r *HttpReq, timeout int) (*HttpResp, error) {
	data := url.Values{}
	if posts != nil && len(posts) > 0 {
		for k, v := range posts {
			data.Set(k, v)
		}
	}

	req, err := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", MimePostForm)

	return HttpDoResp(req, r, timeout)
}

// HttpPostJsonResp Http Post Json 请求, 参数为请求地址, Json 数据 string, HttpReq, 超时时间(毫秒)
// 返回 HttpResp, 错误信息
func HttpPostJsonResp(urlStr string, json string, r *HttpReq, timeout int) (*HttpResp, error) {
	req, err := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(json))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", MimeJson)

	return HttpDoResp(req, r, timeout)
}

// HttpPutResp Http Put, 参数为请求地址, body io.Reader, HttpReq, 超时时间(毫秒)
// 返回 HttpResp, 错误信息
func HttpPutResp(urlStr string, body io.Reader, r *HttpReq, timeout int) (*HttpResp, error) {
	req, err := http.NewRequest(http.MethodPut, urlStr, body)
	if err != nil {
		return nil, err
	}

	return HttpDoResp(req, r, timeout)
}

// HttpPutFormResp Http Put Form, 参数为请求地址, Form 数据 map[string]string, HttpReq, 超时时间(毫秒)
// 返回 HttpResp, 错误信息
func HttpPutFormResp(urlStr string, posts map[string]string, r *HttpReq, timeout int) (*HttpResp, error) {
	data := url.Values{}
	if posts != nil && len(posts) > 0 {
		for k, v := range posts {
			data.Set(k, v)
		}
	}

	req, err := http.NewRequest(http.MethodPut, urlStr, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", MimePostForm)

	return HttpDoResp(req, r, timeout)
}

// HttpPutJsonResp Http Put Json 请求, 参数为请求地址, Json 数据 string, HttpReq, 超时时间(毫秒)
// 返回 HttpResp, 错误信息
func HttpPutJsonResp(urlStr string, json string, r *HttpReq, timeout int) (*HttpResp, error) {
	req, err := http.NewRequest(http.MethodPut, urlStr, strings.NewReader(json))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", MimeJson)

	return HttpDoResp(req, r, timeout)
}

// HttpDo Http 请求, 参数为 http.Request, HttpReq, 超时时间(毫秒)
// 返回 body, 错误信息
func HttpDo(req *http.Request, r *HttpReq, timeout int) ([]byte, error) {
	resp, err := HttpDoResp(req, r, timeout)
	if err != nil {
		return nil, err
	} else {
		return resp.Body, nil
	}
}

// HttpDoResp Http 请求, 参数为 http.Request, HttpReq, 超时时间(毫秒)
// 返回 HttpResp, 错误信息
func HttpDoResp(req *http.Request, r *HttpReq, timeout int) (*HttpResp, error) {
	if timeout == 0 {
		timeout = HttpDefaultTimeOut
	}

	// NewClient
	var client *http.Client
	if r != nil && r.Transport != nil {
		client = &http.Client{
			Timeout:   time.Duration(timeout) * time.Millisecond,
			Transport: r.Transport,
		}
	} else {
		client = &http.Client{
			Timeout:   time.Duration(timeout) * time.Millisecond,
			Transport: HttpDefaultTransport,
		}
	}

	// Redirect 策略
	if r == nil || r.DisableRedirect {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	} else if r.MaxRedirect > 0 && r.MaxRedirect < 10 {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			if len(via) > r.MaxRedirect {
				return http.ErrUseLastResponse
			}
			return nil
		}
	}

	// 处理请求头
	headers := make(map[string]string)
	if r != nil && r.UserAgent != "" {
		r.Headers["User-Agent"] = r.UserAgent
	}
	if r != nil && r.Headers != nil && len(r.Headers) > 0 {
		headers = r.Headers
		if _, exist := headers["User-Agent"]; !exist {
			headers["User-Agent"] = HttpDefaultUserAgent
		}
		if _, exist := headers["Accept-Encoding"]; !exist {
			headers["Accept-Encoding"] = HttpDefaultAcceptEncoding
		}
	} else {
		headers["User-Agent"] = HttpDefaultUserAgent
		headers["Accept-Encoding"] = HttpDefaultAcceptEncoding
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// HttpResp
	httpResp := &HttpResp{
		Success:       false,
		StatusCode:    0,
		Body:          nil,
		ContentLength: 0,
		Headers:       nil,
	}

	// Do
	resp, err := client.Do(req)
	if err != nil {
		return httpResp, err
	}
	defer resp.Body.Close()

	// 响应
	httpResp.StatusCode = resp.StatusCode
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		httpResp.Success = true
	} else {
		return httpResp, errors.New("error with http Status code")
	}

	httpResp.Headers = &resp.Header
	httpResp.ContentLength = resp.ContentLength

	// http.Transport 定义了当请求头不包含 Accept-Encoding 或为空时, 默认会发送 Accept-Encoding=gzip
	// 它会自动判断服务端是否是gzip 然后在接受响应时自动 uncompress, 并会自动移除响应头中的 Content-Encoding、Content-Length
	// 为了获取 Content-Length, 我们需要手动设置不为空的 Accept-Encoding (默认是 HttpDefaultAcceptEncoding), 并且手动 uncompress
	var body []byte
	var reader io.ReadCloser
	switch strings.ToLower(resp.Header.Get("Content-Encoding")) {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return httpResp, errors.New("error with gzip NewReader error")
		}
	case "deflate":
		reader = flate.NewReader(resp.Body)
	default:
		reader = resp.Body
	}
	defer reader.Close()

	// ContentType 限制
	if _, err := allowContentTypes(r, httpResp.Headers); err != nil {
		return httpResp, err
	}

	// ContentLength 限制
	if r != nil && r.MaxContentLength > 0 {
		if resp.ContentLength != -1 {
			if resp.ContentLength > r.MaxContentLength {
				return httpResp, errors.New("error with ContentLength > MaxContentLength ")
			}
			body, err = ioutil.ReadAll(reader)
		} else {
			// 只读取到最大长度
			body, err = ioutil.ReadAll(io.LimitReader(reader, r.MaxContentLength))
		}
	} else {
		body, err = ioutil.ReadAll(reader)
	}

	if err != nil {
		httpResp.Success = false
		return httpResp, err
	} else {
		httpResp.Body = body
	}

	return httpResp, nil
}

// allowContentTypes 判断 Content-Type 是否在允许列表中
func allowContentTypes(r *HttpReq, headers *http.Header) (bool, error) {
	if r == nil {
		return true, nil
	}

	if r.AllowedContentTypes != nil && len(r.AllowedContentTypes) > 0 {
		valid := false

		ct := strings.TrimSpace(strings.ToLower(headers.Get("Content-Type")))
		for _, t := range r.AllowedContentTypes {
			if strings.HasPrefix(ct, t) {
				valid = true
				break
			}
		}

		if valid {
			return valid, nil
		} else {
			return valid, errors.New("error with AllowedContentTypes invalid")
		}
	}

	return true, nil
}
