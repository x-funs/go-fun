package fun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToInt(t *testing.T) {
	assert.Equal(t, 0, ToInt(""))
	assert.Equal(t, 0, ToInt(" "))
	assert.Equal(t, 0, ToInt(" 123 "))
	assert.Equal(t, 123, ToInt("123"))
	assert.Equal(t, 123, ToInt("0123"))
	assert.Equal(t, 0, ToInt("1.1"))
	assert.Equal(t, -123, ToInt("-123"))
}

func TestToInt64(t *testing.T) {
	assert.Equal(t, int64(0), ToInt64(""))
	assert.Equal(t, int64(0), ToInt64(" "))
	assert.Equal(t, int64(0), ToInt64(" 123 "))
	assert.Equal(t, int64(123), ToInt64("123"))
	assert.Equal(t, int64(123), ToInt64("0123"))
	assert.Equal(t, int64(0), ToInt64("1.1"))
	assert.Equal(t, int64(0), ToLong("1.1"))
	assert.Equal(t, int64(-123), ToLong("-123"))
}

func TestToUnit(t *testing.T) {
	assert.Equal(t, uint(0), ToUint(""))
	assert.Equal(t, uint(123), ToUint("0123"))
	assert.Equal(t, uint8(0), ToUint8("-1"))
}

func TestToUtf8AndCharset(t *testing.T) {
	s := []byte{0xd7, 0xd4}
	t.Log(string(s))

	result, err := ToUtf8(s, "gbk")
	if err != nil {
		t.Log(err)
	} else {
		t.Log(string(result))
	}

	ss := "Hello，世界"
	gbk, _ := Utf8To([]byte(ss), "gbk")
	t.Log(string(gbk))

	utf8, _ := ToUtf8(gbk, "gbk")
	t.Log(string(utf8))
}

func TestToBool(t *testing.T) {
	assert.Equal(t, false, ToBool(""))
	assert.Equal(t, true, ToBool("true"))
	assert.Equal(t, false, ToBool("false"))
	assert.Equal(t, false, ToBool(" "))
	assert.Equal(t, false, ToBool("a"))
}

func TestIP2Long(t *testing.T) {
	assert.Equal(t, uint32(2130706433), Ip2Long("127.0.0.1"))
	assert.Equal(t, uint32(3221234342), Ip2Long("192.0.34.166"))
	assert.Equal(t, uint32(659439616), Ip2Long("39.78.64.0"))
	assert.Equal(t, uint32(659439617), Ip2Long("39.78.64.1"))
	assert.Equal(t, uint32(659439870), Ip2Long("39.78.64.254"))
	assert.Equal(t, uint32(659439871), Ip2Long("39.78.64.255"))

	assert.Equal(t, "39.78.64.255", Long2Ip(659439871))
}

func TestToString(t *testing.T) {
	assert.Equal(t, "1", ToString(1))
	assert.Equal(t, "0.123", ToString(0.123))
	assert.Equal(t, "<nil>", ToString(nil))
	assert.Equal(t, "[1 2 3]", ToString([]int{1, 2, 3}))
}

func TestToJson(t *testing.T) {
	assert.Equal(t, `null`, ToJson(nil))
	assert.Equal(t, `""`, ToJson(""))
	assert.Equal(t, `"abc"`, ToJson("abc"))
	assert.Equal(t, `123`, ToJson(123))
	assert.Equal(t, `["a","1","b","2"]`, ToJson([]string{"a", "1", "b", "2"}))
	assert.Equal(t, `{"a":"1","b":"2"}`, ToJson(map[string]string{"a": "1", "b": "2"}))
}
