package fun

import (
	"testing"
)

const (
	TestUrl = "http://localhost:8080"
)

func TestHttpGet(t *testing.T) {
	urlStr := TestUrl + "/get"

	body, _ := HttpGet(urlStr)
	t.Log(body)

	body, _ = HttpGet(urlStr + "?query1=abc1&query2=abc2")
	t.Log(body)

	// 超时时间
	body, _ = HttpGet(urlStr, 1000)
	t.Log(body)

	// Headers 与超时时间
	req := &HttpReq{
		Headers: map[string]string{
			"User-Agent": "test-ua",
			"X-Header":   "test-header",
		},
	}
	body, _ = HttpGet(urlStr, req, 1000)
	t.Log(body)

	// 错误的 Headers
	body, err := HttpGet(urlStr, "error header", 1000)
	t.Log(body)
	t.Log(err)

}

func TestHttpGetProxy(t *testing.T) {
	urlStr := "https://www.ip138.com/"
	req := &HttpReq{
		Headers: map[string]string{
			"User-Agent": "test-ua",
			"X-Header":   "test-header",
		},
		ProxyString: "http://username:password@localhost:8080",
	}
	body, _ := HttpGet(urlStr, req, 5000)
	t.Log(body)
}

func TestHttpPost(t *testing.T) {
	urlStr := TestUrl + "/post"

	body, _ := HttpPost(urlStr)
	t.Log(body)

	// 超时时间
	body, _ = HttpPost(urlStr, 1000)
	t.Log(body)

	// Post 数据
	posts := map[string]string{
		"post1": "post1",
		"post2": "post2",
	}
	body, _ = HttpPost(urlStr, posts)
	t.Log(body)

	// Post 数据与超时时间
	body, _ = HttpPost(urlStr, posts, 1000)
	t.Log(body)

	req := &HttpReq{
		Headers: map[string]string{
			"User-Agent": "test-ua",
			"X-Header":   "test-header",
		},
	}
	body, _ = HttpPost(urlStr, posts, req, 1000)
	t.Log(body)
}

func TestHttpPostJson(t *testing.T) {
	urlStr := TestUrl + "/postJson"

	body, _ := HttpPostJson(urlStr)
	t.Log(body)

	body, _ = HttpPostJson(urlStr, 1000)
	t.Log(body)

	json := `{"param1":"json1","param2":"json2"}`
	body, _ = HttpPostJson(urlStr, json)
	t.Log(body)

	body, _ = HttpPostJson(urlStr, json, 1000)
	t.Log(body)

	req := &HttpReq{
		Headers: map[string]string{
			"User-Agent": "test-ua",
			"X-Header":   "test-header",
		},
	}
	body, _ = HttpPostJson(urlStr, json, req, 1000)
	t.Log(body)
}
