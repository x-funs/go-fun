package fun

import (
	"fmt"
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

func TestEncryptPanic(t *testing.T) {
	encrypt1, err := AesCBCEncrypt("Hello, world!", "123", "0")
	fmt.Println(encrypt1)
	fmt.Println(err)
}

func TestDecryptPanic(t *testing.T) {
	decrypt1, err := AesCBCDecrypt("f87cd942", "0123456789abcdef", "0123456789abcdef")
	fmt.Println(decrypt1)
	fmt.Println(err)
}

func TestAesCBCEncryptWithInvalidKey(t *testing.T) {
	// 测试错误的 key 长度
	_, err := AesCBCEncrypt("Hello, world!", "123", "0123456789abcdef")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "key length must be")
}

func TestAesCBCEncryptWithInvalidIV(t *testing.T) {
	// 测试错误的 iv 长度
	_, err := AesCBCEncrypt("Hello, world!", "0123456789abcdef", "0")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "iv length must be 16 bytes")
}

func TestAesCBCDecryptWithInvalidKey(t *testing.T) {
	// 测试错误的 key 长度
	_, err := AesCBCDecrypt("f87cd9421d03a38d8a8353d0b1d85d73", "123", "0123456789abcdef")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "key length must be")
}

func TestAesCBCDecryptWithInvalidIV(t *testing.T) {
	// 测试错误的 iv 长度
	_, err := AesCBCDecrypt("f87cd9421d03a38d8a8353d0b1d85d73", "0123456789abcdef", "0")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "iv length must be 16 bytes")
}

func TestAesCBCDecryptWithEmptyCipher(t *testing.T) {
	// 测试空的密文字符串
	_, err := AesCBCDecrypt("", "0123456789abcdef", "0123456789abcdef")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cipher string cannot be empty")
}

func TestAesCBCDecryptWithInvalidCipherFormat(t *testing.T) {
	// 测试错误格式的密文字符串
	_, err := AesCBCDecrypt("invalid-hex", "0123456789abcdef", "0123456789abcdef")
	assert.Error(t, err)
}

func TestAesCBCDecryptWithInvalidCipherLength(t *testing.T) {
	// 测试密文长度不是块大小倍数的情况
	// 12 字节的密文，不是 16 的倍数
	_, err := AesCBCDecrypt("f87cd9421d03", "0123456789abcdef", "0123456789abcdef")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ciphertext length must be a multiple of the block size")
}

func TestAesGCMEncryptDecrypt(t *testing.T) {
	// 测试正常的 GCM 模式加密解密
	key := "0123456789abcdef"
	nonce := "0123456789ab" // 12 字节
	plaintext := "Hello, world!"

	// 加密
	cipherHex, tagHex, err := AesGCMEncrypt(plaintext, key, nonce)
	assert.NoError(t, err)
	assert.NotEmpty(t, cipherHex)
	assert.NotEmpty(t, tagHex)

	t.Log(cipherHex)
	t.Log(tagHex)

	// 解密
	decrypted, err := AesGCMDecrypt(cipherHex, tagHex, key, nonce)
	assert.NoError(t, err)
	assert.Equal(t, plaintext, decrypted)

	t.Log(decrypted)

	// 测试中文
	chinesePlaintext := "Hello, 你好，中国！"
	cipherHex2, tagHex2, err := AesGCMEncrypt(chinesePlaintext, key, nonce)
	assert.NoError(t, err)
	assert.NotEmpty(t, cipherHex2)
	assert.NotEmpty(t, tagHex2)

	t.Log(cipherHex2)
	t.Log(tagHex2)

	decrypted2, err := AesGCMDecrypt(cipherHex2, tagHex2, key, nonce)
	assert.NoError(t, err)
	assert.Equal(t, chinesePlaintext, decrypted2)

	t.Log(chinesePlaintext)
}

func TestAesGCMEncryptWithInvalidKey(t *testing.T) {
	// 测试错误的 key 长度
	_, _, err := AesGCMEncrypt("Hello, world!", "123", "0123456789ab")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "key length must be")
}

func TestAesGCMEncryptWithInvalidNonce(t *testing.T) {
	// 测试错误的 nonce 长度
	_, _, err := AesGCMEncrypt("Hello, world!", "0123456789abcdef", "0")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "nonce length should be 12 bytes")
}

func TestAesGCMDecryptWithInvalidKey(t *testing.T) {
	// 测试错误的 key 长度
	_, err := AesGCMDecrypt("f87cd942", "tag123", "123", "0123456789ab")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "key length must be")
}

func TestAesGCMDecryptWithInvalidNonce(t *testing.T) {
	// 测试错误的 nonce 长度
	_, err := AesGCMDecrypt("f87cd942", "tag123", "0123456789abcdef", "0")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "nonce length should be 12 bytes")
}

func TestAesGCMDecryptWithEmptyCipher(t *testing.T) {
	// 测试空的密文字符串
	_, err := AesGCMDecrypt("", "tag123", "0123456789abcdef", "0123456789ab")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cipher string cannot be empty")
}

func TestAesGCMDecryptWithEmptyTag(t *testing.T) {
	// 测试空的标签字符串
	_, err := AesGCMDecrypt("f87cd942", "", "0123456789abcdef", "0123456789ab")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "tag string cannot be empty")
}
