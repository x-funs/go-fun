package fun

import (
	"fmt"
	"testing"
)

func TestHttpGet(t *testing.T) {
	body, _ := HttpGet("http://localhost:9090/get")
	t.Log(body)

	body, _ = HttpGet("http://localhost:9090/get?param1=abc1&param2=abc2")
	t.Log(body)

	body, _ = HttpGet("http://localhost:9090/get", 1)
	t.Log(body)

	headers := map[string]string{
		"User-Agent": "test-ua",
		"X-Header":   "test-header",
	}
	body, _ = HttpGet("http://localhost:9090/get", 1, headers)
	t.Log(body)

	body, err := HttpGet("http://localhost:9090/get", 1, "error header")
	t.Log(body)
	fmt.Println(err)
}

func TestHttpGetBody(t *testing.T) {
	headers := map[string]string{
		"X-Test": "123",
	}
	body, _ := HttpGetBody("", 0, headers)

	fmt.Println(body)
}
