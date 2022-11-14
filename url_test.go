package fun

import (
	"net/url"
	"testing"
)

func BenchmarkUrlParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = url.Parse("http://www.baidu.com")
	}
}
