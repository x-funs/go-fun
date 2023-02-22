package fun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAesCBCEncrypt(t *testing.T) {
	encrypt1, err := AesCBCEncrypt("Hello, world!", "0123456789abcdef", "0123456789abcdef")
	t.Log(err)
	assert.Equal(t, "f87cd9421d03a38d8a8353d0b1d85d73", encrypt1)

	bit16 := Md5Bit16("")
	t.Log(bit16)
	encrypt2, _ := AesCBCEncrypt("Hello, 你好，中国！", bit16, bit16)
	assert.Equal(t, "2ba4f416d2f6dcaa13661933cf56db41a02fdeef1d210b1cda643cd71957ecd8", encrypt2)
}

func TestAesCBCDecrypt(t *testing.T) {
	decrypt1, err := AesCBCDecrypt("f87cd9421d03a38d8a8353d0b1d85d73", "0123456789abcdef", "0123456789abcdef")
	t.Log(err)
	assert.Equal(t, "Hello, world!", decrypt1)

	bit16 := Md5Bit16("")
	t.Log(bit16)
	decrypt2, _ := AesCBCDecrypt("2ba4f416d2f6dcaa13661933cf56db41a02fdeef1d210b1cda643cd71957ecd8", bit16, bit16)
	assert.Equal(t, "Hello, 你好，中国！", decrypt2)
}
