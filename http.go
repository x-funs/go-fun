package fun

import (
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	HttpDefaultTimeOut   = 5000
	HttpDefaultUserAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36"
)

type HttpReq struct {
	Headers     map[string]string
	ProxyString string
}

type HttpResp struct {
	Success       bool
	StatusCode    int
	Body          string
	ContentLength int64
	Headers       http.Header
}

// HttpGet 参数为请求地址（HttpReq, 超时时间）
// HttpGet(url)、HttpGet(url, timeout)、HttpGet(url, httpReq, timeout)
// 返回值为请求内容 String，错误信息
func HttpGet(urlStr string, args ...any) (string, error) {
	l := len(args)

	switch l {
	case 0:
		return HttpGetBody(urlStr, nil, 0)
	case 1:
		timeout := ToInt(args[0])
		return HttpGetBody(urlStr, nil, timeout)
	case 2:
		timeout := ToInt(args[1])
		switch v := args[0].(type) {
		case *HttpReq:
			return HttpGetBody(urlStr, v, timeout)
		}
	}

	return "", errors.New("HttpGet() 参数错误")
}

// HttpPost 参数为请求地址（Form 数据 map[string]string，HttpReq，超时时间）
// HttpPost(url)、HttpPost(url, timeout)、HttpPost(url, posts)、HttpPost(url, posts, timeout)、HttpPost(url, posts, httpReq, timeout)
// 返回值为请求内容 String，错误信息
func HttpPost(urlStr string, args ...any) (string, error) {
	l := len(args)

	switch l {
	case 0:
		return HttpPostBody(urlStr, nil, nil, 0)
	case 1:
		switch v := args[0].(type) {
		case int:
			timeout := ToInt(args[0])
			return HttpPostBody(urlStr, nil, nil, timeout)
		case map[string]string:
			return HttpPostBody(urlStr, v, nil, 0)
		}
	case 2:
		switch v := args[0].(type) {
		case map[string]string:
			timeout := ToInt(args[1])
			return HttpPostBody(urlStr, v, nil, timeout)
		}
	case 3:
		switch v := args[0].(type) {
		case map[string]string:
			switch h := args[1].(type) {
			case *HttpReq:
				timeout := ToInt(args[2])
				return HttpPostBody(urlStr, v, h, timeout)
			}
		}
	}

	return "", errors.New("HttpPost() 参数错误")
}

// HttpPostJson 参数为请求地址（Json 数据 string，HttpReq, 超时时间）
// HttpPostJson(url)、HttpPostJson(url, timeout)、HttpPostJson(url, json)、HttpPost(url, json, timeout)、HttpPost(url, json, httpReq, timeout)
// 返回值为请求内容 String，错误信息
func HttpPostJson(urlStr string, args ...any) (string, error) {
	l := len(args)
	switch l {
	case 0:
		return HttpPostJsonBody(urlStr, "{}", nil, 0)
	case 1:
		switch v := args[0].(type) {
		case int:
			timeout := ToInt(args[0])
			return HttpPostJsonBody(urlStr, "{}", nil, timeout)
		case string:
			return HttpPostJsonBody(urlStr, v, nil, 0)
		}
	case 2:
		switch v := args[0].(type) {
		case string:
			timeout := ToInt(args[1])
			return HttpPostJsonBody(urlStr, v, nil, timeout)
		}
	case 3:
		switch v := args[0].(type) {
		case string:
			switch h := args[1].(type) {
			case *HttpReq:
				timeout := ToInt(args[2])
				return HttpPostJsonBody(urlStr, v, h, timeout)
			}
		}
	}

	return "", errors.New("HttpPostJson() 参数错误")
}

// HttpGetBody Http Get 请求，参数为请求地址，HttpReq，超时时间(毫秒)
// 返回请求内容 String，错误信息
func HttpGetBody(urlStr string, r *HttpReq, timeout int) (string, error) {
	resp, err := HttpGetResp(urlStr, r, timeout)
	if err != nil {
		return "", err
	} else {
		return resp.Body, nil
	}
}

// HttpPostBody Http Post Form，参数为请求地址，Form 数据 map[string]string，HttpReq，超时时间(毫秒)
// 返回请求内容 String，错误信息
func HttpPostBody(urlStr string, posts map[string]string, r *HttpReq, timeout int) (string, error) {
	resp, err := HttpPostResp(urlStr, posts, r, timeout)
	if err != nil {
		return "", err
	} else {
		return resp.Body, nil
	}
}

// HttpPostJsonBody Http Post Json 请求，参数为请求地址，Json 数据 string，HttpReq，超时时间(毫秒)
// 返回请求内容 String，错误信息
func HttpPostJsonBody(urlStr string, json string, r *HttpReq, timeout int) (string, error) {
	resp, err := HttpPostJsonResp(urlStr, json, r, timeout)
	if err != nil {
		return "", err
	} else {
		return resp.Body, nil
	}
}

// HttpGetResp Http Get 请求，参数为请求地址，HttpReq，超时时间(毫秒)
// 返回 HttpResp，错误信息
func HttpGetResp(urlStr string, r *HttpReq, timeout int) (*HttpResp, error) {
	httpResp := &HttpResp{
		Success:       false,
		StatusCode:    0,
		Body:          "",
		ContentLength: 0,
		Headers:       nil,
	}

	if timeout == 0 {
		timeout = HttpDefaultTimeOut
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// 设置 Proxy 代理
	if r != nil && r.ProxyString != "" {
		proxy, _ := url.Parse(r.ProxyString)
		transport.Proxy = http.ProxyURL(proxy)
	}

	client := &http.Client{
		Timeout:   time.Duration(timeout) * time.Millisecond,
		Transport: transport,
	}

	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return httpResp, err
	}

	// 处理请求头
	headers := make(map[string]string)
	if r != nil && r.Headers != nil && len(r.Headers) > 0 {
		headers = r.Headers
		if _, exist := headers["User-Agent"]; !exist {
			headers["User-Agent"] = HttpDefaultUserAgent
		}
	} else {
		headers["User-Agent"] = HttpDefaultUserAgent
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return httpResp, err
	}

	httpResp.StatusCode = resp.StatusCode
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		httpResp.Success = true
	}
	httpResp.Headers = resp.Header
	httpResp.ContentLength = resp.ContentLength

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return httpResp, err
	} else {
		httpResp.Body = BytesToString(body)
	}

	return httpResp, nil
}

// HttpPostResp Http Post Form，参数为请求地址，Form 数据 map[string]string，HttpReq，超时时间(毫秒)
// 返回 HttpResp，错误信息
func HttpPostResp(urlStr string, posts map[string]string, r *HttpReq, timeout int) (*HttpResp, error) {
	httpResp := &HttpResp{
		Success:       false,
		StatusCode:    0,
		Body:          "",
		ContentLength: 0,
		Headers:       nil,
	}

	if timeout == 0 {
		timeout = HttpDefaultTimeOut
	}

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Millisecond,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	data := url.Values{}
	if posts != nil && len(posts) > 0 {
		for k, v := range posts {
			data.Set(k, v)
		}
	}

	// 处理 Form
	req, err := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))
	if err != nil {
		return httpResp, err
	}

	// 处理请求头
	headers := make(map[string]string)
	if r != nil && r.Headers != nil && len(r.Headers) > 0 {
		headers = r.Headers
		if _, exist := headers["User-Agent"]; !exist {
			headers["User-Agent"] = HttpDefaultUserAgent
		}
	} else {
		headers["User-Agent"] = HttpDefaultUserAgent
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	req.Header.Set("Content-Type", MimeMultipartPostForm)

	resp, err := client.Do(req)
	if err != nil {
		return httpResp, err
	}

	httpResp.StatusCode = resp.StatusCode
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		httpResp.Success = true
	}
	httpResp.Headers = resp.Header
	httpResp.ContentLength = resp.ContentLength

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return httpResp, err
	} else {
		httpResp.Body = BytesToString(body)
	}

	return httpResp, nil
}

// HttpPostJsonResp Http Post Json 请求，参数为请求地址，Json 数据 string，HttpReq，超时时间(毫秒)
// 返回 HttpResp，错误信息
func HttpPostJsonResp(urlStr string, json string, r *HttpReq, timeout int) (*HttpResp, error) {
	httpResp := &HttpResp{
		Success:       false,
		StatusCode:    0,
		Body:          "",
		ContentLength: 0,
		Headers:       nil,
	}

	if timeout == 0 {
		timeout = HttpDefaultTimeOut
	}

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// 处理 Json
	req, err := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(json))
	if err != nil {
		return httpResp, err
	}

	// 处理请求头
	headers := make(map[string]string)
	if r != nil && r.Headers != nil && len(r.Headers) > 0 {
		headers = r.Headers
		if _, exist := headers["User-Agent"]; !exist {
			headers["User-Agent"] = HttpDefaultUserAgent
		}
	} else {
		headers["User-Agent"] = HttpDefaultUserAgent
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	req.Header.Set("Content-Type", MimeJson)

	resp, err := client.Do(req)
	if err != nil {
		return httpResp, err
	}

	httpResp.StatusCode = resp.StatusCode
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		httpResp.Success = true
	}
	httpResp.Headers = resp.Header
	httpResp.ContentLength = resp.ContentLength

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return httpResp, err
	} else {
		httpResp.Body = BytesToString(body)
	}

	return httpResp, nil
}
