package fun

import (
	"net/url"
	"strings"
	"testing"
)

const (
	TestUrl = "http://localhost:8080"
)

func TestHttpGet(t *testing.T) {
	urlStr := TestUrl + "/get"

	body, _ := HttpGet(urlStr)
	t.Log(BytesToString(body))

	body, _ = HttpGet(urlStr + "?query1=value1&query2=value2")
	t.Log(BytesToString(body))

	// 超时时间
	body, _ = HttpGet(urlStr, 1000)
	t.Log(BytesToString(body))

	// Headers 与超时时间
	req := &HttpReq{
		UserAgent: "test-ua",
		Headers: map[string]string{
			"X-Header": "test-header",
		},
	}
	body, _ = HttpGet(urlStr+"?query1=value1", req, 1000)
	t.Log(BytesToString(body))

	// 错误的 Headers
	body, err := HttpGet(urlStr, "error header", 1000)
	t.Log(BytesToString(body))
	t.Log(err)

}

func TestHttpPostForm(t *testing.T) {
	urlStr := TestUrl + "/postForm"

	// Post 数据
	posts := map[string]string{
		"post1": "post1",
	}
	body, _ := HttpPostForm(urlStr, posts)
	t.Log(BytesToString(body))

	// Post 数据与超时时间
	body, _ = HttpPostForm(urlStr, posts, 1000)
	t.Log(BytesToString(body))

	req := &HttpReq{
		UserAgent: "test-ua",
		Headers: map[string]string{
			"X-Header": "test-header",
		},
	}
	body, _ = HttpPostForm(urlStr, posts, req, 1000)
	t.Log(BytesToString(body))
}

func TestHttpPutForm(t *testing.T) {
	urlStr := TestUrl + "/putForm"

	// Post 数据
	posts := map[string]string{
		"post1": "post1",
	}
	body, _ := HttpPutForm(urlStr, posts)
	t.Log(BytesToString(body))

	// Post 数据与超时时间
	body, _ = HttpPutForm(urlStr, posts, 1000)
	t.Log(BytesToString(body))

	req := &HttpReq{
		UserAgent: "test-ua",
		Headers: map[string]string{
			"X-Header": "test-header",
		},
	}
	body, _ = HttpPutForm(urlStr, posts, req, 1000)
	t.Log(BytesToString(body))
}

func TestHttpPostJson(t *testing.T) {
	urlStr := TestUrl + "/bindPostJson"

	body, _ := HttpPostJson(urlStr)
	t.Log(BytesToString(body))

	body, _ = HttpPostJson(urlStr, 1000)
	t.Log(BytesToString(body))

	json := `{"username":"admin","email":"admin@admin.com", "joinTime":"2006-01-02 15:04:05", "isVip":true}`
	body, _ = HttpPostJson(urlStr, json)
	t.Log(BytesToString(body))

	body, _ = HttpPostJson(urlStr, json, 1000)
	t.Log(BytesToString(body))

	req := &HttpReq{
		UserAgent: "test-ua",
		Headers: map[string]string{
			"X-Header": "test-header",
		},
	}
	body, _ = HttpPostJson(urlStr, json, req, 1000)
	t.Log(BytesToString(body))
}

func TestHttpPutJson(t *testing.T) {
	urlStr := TestUrl + "/bindPutJson"

	body, _ := HttpPutJson(urlStr)
	t.Log(BytesToString(body))

	body, _ = HttpPutJson(urlStr, 1000)
	t.Log(BytesToString(body))

	json := `{"username":"admin","email":"admin@admin.com", "joinTime":"2006-01-02 15:04:05", "isVip":true}`
	body, _ = HttpPutJson(urlStr, json)
	t.Log(BytesToString(body))

	body, _ = HttpPutJson(urlStr, json, 1000)
	t.Log(BytesToString(body))

	req := &HttpReq{
		UserAgent: "test-ua",
		Headers: map[string]string{
			"X-Header": "test-header",
		},
	}
	body, _ = HttpPutJson(urlStr, json, req, 1000)
	t.Log(BytesToString(body))
}

func TestHttpPost(t *testing.T) {
	urlStr := TestUrl + "/postForm"

	data := url.Values{}
	data.Set("post1", "post1")
	data.Set("post2", "post2")
	b := strings.NewReader(data.Encode())

	req := &HttpReq{
		UserAgent: "test-ua",
		Headers: map[string]string{
			"Content-Type": MimePostForm,
		},
	}

	body, err := HttpPost(urlStr, b, req, 1000)
	t.Log(BytesToString(body))
	t.Log(err)
}

func TestHttpPut(t *testing.T) {
	urlStr := TestUrl + "/putForm"

	data := url.Values{}
	data.Set("post1", "post1")
	data.Set("post2", "post2")
	b := strings.NewReader(data.Encode())

	req := &HttpReq{
		UserAgent: "test-ua",
		Headers: map[string]string{
			"Content-Type": MimePostForm,
		},
	}

	body, err := HttpPut(urlStr, b, req, 1000)
	t.Log(BytesToString(body))
	t.Log(err)
}

func TestHttpDelete(t *testing.T) {
	urlStr := TestUrl + "/delete/path"

	// Headers 与超时时间
	req := &HttpReq{
		UserAgent: "test-ua",
		Headers: map[string]string{
			"X-Header": "test-header",
		},
	}
	body, _ := HttpDelete(urlStr+"?query1=value1", req, 1000)
	t.Log(BytesToString(body))

	// 错误的 Headers
	body, err := HttpDelete(urlStr, "error header", 1000)
	t.Log(BytesToString(body))
	t.Log(err)
}
