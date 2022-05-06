package fun

import (
	"fmt"
	"testing"
)

func TestHttpGetBody(t *testing.T) {
	headers := map[string]string{
		"X-Test": "123",
	}
	body, _ := HttpGetBody("", 0, headers)

	fmt.Println(body)
}
