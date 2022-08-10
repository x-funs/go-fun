package fun

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"testing"
)

const (
	TestUrl = "http://localhost:8080"
)

func TestHttpGet(t *testing.T) {
	urlStr := TestUrl + "/get"

	body, _ := HttpGet(urlStr)
	t.Log(String(body))

	body, _ = HttpGet(urlStr + "?query1=value1&query2=value2")
	t.Log(String(body))

	// 超时时间
	body, _ = HttpGet(urlStr, 1000)
	t.Log(String(body))

	// Headers 与超时时间
	r := &HttpReq{
		UserAgent: "test-ua",
		Headers: map[string]string{
			"X-Header": "test-header",
		},
	}
	body, _ = HttpGet(urlStr+"?query1=value1", r, 1000)
	t.Log(String(body))

	// 错误的 Headers
	body, err := HttpGet(urlStr, "error header", 1000)
	t.Log(String(body))
	t.Log(err)

}

func TestHttpPostForm(t *testing.T) {
	urlStr := TestUrl + "/postForm"

	// Post 数据
	posts := map[string]string{
		"post1": "post1",
	}
	body, _ := HttpPostForm(urlStr, posts)
	t.Log(String(body))

	// Post 数据与超时时间
	body, _ = HttpPostForm(urlStr, posts, 1000)
	t.Log(String(body))

	r := &HttpReq{
		UserAgent: "test-ua",
		Headers: map[string]string{
			"X-Header": "test-header",
		},
	}
	body, _ = HttpPostForm(urlStr, posts, r, 1000)
	t.Log(String(body))
}

func TestHttpPutForm(t *testing.T) {
	urlStr := TestUrl + "/putForm"

	// Post 数据
	posts := map[string]string{
		"post1": "post1",
	}
	body, _ := HttpPutForm(urlStr, posts)
	t.Log(String(body))

	// Post 数据与超时时间
	body, _ = HttpPutForm(urlStr, posts, 1000)
	t.Log(String(body))

	r := &HttpReq{
		UserAgent: "test-ua",
		Headers: map[string]string{
			"X-Header": "test-header",
		},
	}
	body, _ = HttpPutForm(urlStr, posts, r, 1000)
	t.Log(String(body))
}

func TestHttpPostJson(t *testing.T) {
	urlStr := TestUrl + "/bindPostJson"

	body, _ := HttpPostJson(urlStr)
	t.Log(String(body))

	body, _ = HttpPostJson(urlStr, 1000)
	t.Log(String(body))

	json := `{"username":"admin","email":"admin@admin.com", "joinTime":"2006-01-02 15:04:05", "isVip":true}`
	body, _ = HttpPostJson(urlStr, json)
	t.Log(String(body))

	body, _ = HttpPostJson(urlStr, json, 1000)
	t.Log(String(body))

	r := &HttpReq{
		UserAgent: "test-ua",
		Headers: map[string]string{
			"X-Header": "test-header",
		},
	}
	body, _ = HttpPostJson(urlStr, json, r, 1000)
	t.Log(String(body))
}

func TestHttpPutJson(t *testing.T) {
	urlStr := TestUrl + "/bindPutJson"

	body, _ := HttpPutJson(urlStr)
	t.Log(String(body))

	body, _ = HttpPutJson(urlStr, 1000)
	t.Log(String(body))

	json := `{"username":"admin","email":"admin@admin.com", "joinTime":"2006-01-02 15:04:05", "isVip":true}`
	body, _ = HttpPutJson(urlStr, json)
	t.Log(String(body))

	body, _ = HttpPutJson(urlStr, json, 1000)
	t.Log(String(body))

	r := &HttpReq{
		UserAgent: "test-ua",
		Headers: map[string]string{
			"X-Header": "test-header",
		},
	}
	body, _ = HttpPutJson(urlStr, json, r, 1000)
	t.Log(String(body))
}

func TestHttpPost(t *testing.T) {
	urlStr := TestUrl + "/postForm"

	data := url.Values{}
	data.Set("post1", "post1")
	data.Set("post2", "post2")
	b := strings.NewReader(data.Encode())

	r := &HttpReq{
		UserAgent: "test-ua",
		Headers: map[string]string{
			"Content-Type": MimePostForm,
		},
	}

	body, err := HttpPost(urlStr, b, r, 1000)
	t.Log(String(body))
	t.Log(err)
}

func TestHttpPut(t *testing.T) {
	urlStr := TestUrl + "/putForm"

	data := url.Values{}
	data.Set("post1", "post1")
	data.Set("post2", "post2")
	b := strings.NewReader(data.Encode())

	r := &HttpReq{
		UserAgent: "test-ua",
		Headers: map[string]string{
			"Content-Type": MimePostForm,
		},
	}

	body, err := HttpPut(urlStr, b, r, 1000)
	t.Log(String(body))
	t.Log(err)
}

func TestHttpDelete(t *testing.T) {
	urlStr := TestUrl + "/delete/path"

	// Headers 与超时时间
	r := &HttpReq{
		UserAgent: "test-ua",
		Headers: map[string]string{
			"X-Header": "test-header",
		},
	}
	body, _ := HttpDelete(urlStr+"?query1=value1", r, 1000)
	t.Log(String(body))

	// 错误的 Headers
	body, err := HttpDelete(urlStr, "error header", 1000)
	t.Log(String(body))
	t.Log(err)
}

func TestHttpDo(t *testing.T) {
	urlStr := TestUrl + "/get?query1=value1&query2=value2"

	req, err := http.NewRequest(http.MethodGet, urlStr, nil)

	// Headers 与超时时间
	r := &HttpReq{
		UserAgent: "test-ua",
		Headers: map[string]string{
			"X-Header": "test-header",
		},
	}
	body, err := HttpDo(req, r, 1000)
	t.Log(String(body))
	t.Log(err)
}

func TestHttpDoResp(t *testing.T) {
	var urlStr string

	urlStr = TestUrl + "/get?query1=value1&query2=value2"

	req, err := http.NewRequest(http.MethodGet, urlStr, nil)

	// Headers 与超时时间
	r := &HttpReq{
		UserAgent: "test-ua",
		Headers: map[string]string{
			"X-Header": "test-header",
			// "Accept-Encoding": "gzip",
		},
	}
	resp, err := HttpDoResp(req, r, 1000)
	t.Log(resp.Success)
	t.Log(resp.StatusCode)
	t.Log(resp.Headers)
	t.Log(String(resp.Body))
	t.Log(err)
}

func TestHttpGetWithProxy(t *testing.T) {
	transport := &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DisableKeepAlives: true,
	}
	proxyString := "http://username:password@host:port"
	proxy, _ := url.Parse(proxyString)
	transport.Proxy = http.ProxyURL(proxy)

	r := &HttpReq{
		Transport: transport,
	}

	body, _ := HttpGet("http://test/ip.php", r)
	t.Log(String(body))
}

func TestHttpSharedTransport(t *testing.T) {
	urlStr := TestUrl + "/get"

	transport := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		DisableKeepAlives:   true,
		MaxIdleConnsPerHost: RandomInt(10, 100),
	}

	var wg sync.WaitGroup

	// 使用了共享的 Transport
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			r := &HttpReq{
				Transport: transport,
			}
			body, _ := HttpGet(urlStr, r)
			t.Log(String(body))

			defer wg.Done()
		}()
	}

	wg.Wait()
}

func TestHttpTransport(t *testing.T) {
	urlStr := TestUrl + "/get"

	var wg sync.WaitGroup

	// 使用了不同的 Transport
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			transport := &http.Transport{
				TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
				DisableKeepAlives:   true,
				MaxIdleConnsPerHost: RandomInt(10, 100),
			}

			r := &HttpReq{
				Transport: transport,
			}
			body, _ := HttpGet(urlStr, r)
			t.Log(String(body))

			defer wg.Done()
		}()
	}

	wg.Wait()
}

func TestHttpGetPublic(t *testing.T) {
	var urlStr string

	// urlStr = "http://www.163.com"
	urlStr = "http://www.qq.com"

	// 默认不会进行编码探测和转换, gbk 会输出乱码
	resp, err := HttpGetResp(urlStr, nil, 10000)

	t.Log(err)
	t.Log(resp.Success)
	t.Log(resp.StatusCode)
	t.Log(resp.Headers)
	t.Log(String(resp.Body))
}

func TestHttpGetMaxContentLength(t *testing.T) {
	var urlStr string

	// urlStr = "https://www.163.com"
	urlStr = "https://mirrors.163.com/mysql/Downloads/MySQL-8.0/mysql-8.0.27-macos11-x86_64.tar"

	req := &HttpReq{
		MaxContentLength: 10000,
	}
	resp, err := HttpGetResp(urlStr, req, 10000)

	t.Log(err)
	t.Log(resp.Success)
	t.Log(resp.ContentLength)
	t.Log(resp.Headers)
	t.Log(String(resp.Body))
}

func TestHttpGetContentType(t *testing.T) {
	var urlStr string

	urlStr = "https://mirrors.163.com/mysql/Downloads/MySQL-8.0/mysql-8.0.27-macos11-x86_64.tar"

	req := &HttpReq{
		AllowedContentTypes: []string{"text/html"},
	}
	resp, err := HttpGetResp(urlStr, req, 10000)

	t.Log(err)
	t.Log(resp.Success)
	t.Log(resp.ContentLength)
	t.Log(resp.Headers)
	t.Log(String(resp.Body))
}

func TestHttpGetRedirect(t *testing.T) {
	var urlStr string

	// http 默认会 30x 跳转到 https
	urlStr = "http://www.sohu.com"

	req := &HttpReq{
		MaxRedirect: 2,
	}
	resp, err := HttpGetResp(urlStr, req, 10000)

	t.Log(err)
	t.Log(resp.Success)
	t.Log(resp.StatusCode)
	t.Log(resp.ContentLength)
	t.Log(resp.Headers)
	t.Log(resp.RequestURL)
	t.Log(String(resp.Body))
}
