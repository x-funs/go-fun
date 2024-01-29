package fun

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNumber(t *testing.T) {
	assert.Equal(t, false, IsNumber(""))
	assert.Equal(t, true, IsNumber("1"))
	assert.Equal(t, true, IsNumber("012345"))
}

func TestIsLetter(t *testing.T) {
	assert.Equal(t, false, IsLetter(""))
	assert.Equal(t, false, IsLetter("   "))
	assert.Equal(t, true, IsLetter("a"))
	assert.Equal(t, true, IsLetter("abc"))
	assert.Equal(t, false, IsLetter("abc123"))
	assert.Equal(t, false, IsLetter("123"))
	assert.Equal(t, true, IsLetter("上"))
	assert.Equal(t, true, IsLetter("海"))
}

func BenchmarkIsLetter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsLetter("上")
	}

	fmt.Printf("%+v\n", IsLetter("上"))
}

func TestIsASCIILetter(t *testing.T) {
	assert.Equal(t, false, IsASCIILetter(""))
	assert.Equal(t, false, IsASCIILetter("   "))
	assert.Equal(t, true, IsASCIILetter("a"))
	assert.Equal(t, true, IsASCIILetter("abc"))
	assert.Equal(t, false, IsASCIILetter("abc123"))
	assert.Equal(t, false, IsASCIILetter("123"))
	assert.Equal(t, false, IsASCIILetter("上"))
	assert.Equal(t, false, IsASCIILetter("海"))
}

func BenchmarkIsASCIILetter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsASCIILetter("上")
	}

	fmt.Printf("%+v\n", IsASCIILetter("上"))
}

func TestIsEmail(t *testing.T) {
	assert.Equal(t, false, IsEmail(""))
	assert.Equal(t, true, IsEmail("aaa@aa.aa"))
}

func TestIsASCII(t *testing.T) {
	assert.Equal(t, true, IsASCII(""))
	assert.Equal(t, true, IsASCII("#"))
	assert.Equal(t, false, IsASCII("中文"))
}

func TestIsDir(t *testing.T) {
	t.Log(IsDir("/tmp"))
	t.Log(IsDir("/tmp/test"))
}

func TestIsExist(t *testing.T) {
	t.Log(IsExist("/tmp"))
	t.Log(IsExist("/tmp/test"))
}

func TestIsUtf8(t *testing.T) {
	assert.Equal(t, true, IsUtf8(Bytes("中文")))
}

func TestIsIp(t *testing.T) {
	assert.Equal(t, true, IsIp("192.168.0.100"))
	assert.Equal(t, true, IsIpV4("36.112.24.10"))
	assert.Equal(t, false, IsIpV4("2001:db8:4006:812::200e"))
	assert.Equal(t, true, IsIpV6("2001:db8:4006:812::200e"))
	assert.Equal(t, true, IsIpV6("::1"))
}
