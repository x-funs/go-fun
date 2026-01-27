package fun

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
)

// AesCBCEncrypt Aes CBC 对称加密, key 的长度决定 AES-128, AES-192, or AES-256
func AesCBCEncrypt(text string, key string, iv string) (string, error) {
	var result string
	var err error

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v", r)
		}
	}()

	// 参数验证
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", errors.New("key length must be 16, 24, or 32 bytes for AES-128, AES-192, or AES-256")
	}

	if len(iv) != 16 {
		return "", errors.New("iv length must be 16 bytes for CBC mode")
	}

	textBytes := Bytes(text)
	keyBytes := Bytes(key)
	ivBytes := Bytes(iv)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	// 对数据进行填充，使其满足加密块大小，加密块大小为 16 字节
	blockSize := block.BlockSize()
	paddingText := pKCS7Padding(textBytes, blockSize)

	// 创建加密块链，使用 CBC 加密模式，iv 的长度需要和 block.BlockSize() 一致
	mode := cipher.NewCBCEncrypter(block, ivBytes)

	// 加密数据
	cipherText := make([]byte, len(paddingText))
	mode.CryptBlocks(cipherText, paddingText)
	cipherHex := hex.EncodeToString(cipherText)

	result = cipherHex
	return result, err
}

// AesCBCDecrypt Aes CBC 对称加密
func AesCBCDecrypt(cipherStr string, key string, iv string) (string, error) {
	var result string
	var err error

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v", r)
		}
	}()

	// 参数验证
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", errors.New("key length must be 16, 24, or 32 bytes for AES-128, AES-192, or AES-256")
	}

	if len(iv) != 16 {
		return "", errors.New("iv length must be 16 bytes for CBC mode")
	}

	if len(cipherStr) == 0 {
		return "", errors.New("cipher string cannot be empty")
	}

	cipherBytes, err := hex.DecodeString(cipherStr)
	if err != nil {
		return "", err
	}

	keyBytes := Bytes(key)
	ivBytes := Bytes(iv)

	// 创建解密器
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	// 验证密文长度是否为块大小的倍数
	if len(cipherBytes)%block.BlockSize() != 0 {
		return "", errors.New("ciphertext length must be a multiple of the block size")
	}

	// 创建解密块链
	mode := cipher.NewCBCDecrypter(block, ivBytes)

	// 解密数据
	textBytes := make([]byte, len(cipherBytes))
	mode.CryptBlocks(textBytes, cipherBytes)

	textBytes = pKCS7UnPadding(textBytes)

	result = String(textBytes)
	return result, err
}

// pKCS7Padding 对数据进行填充，满足加密块大小
func pKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(data, padText...)
}

// pKCS7UnPadding 去除填充的数据
func pKCS7UnPadding(data []byte) []byte {
	length := len(data)
	unPadding := int(data[length-1])

	return data[:(length - unPadding)]
}

// AesGCMEncrypt Aes GCM 对称加密, key 的长度决定 AES-128, AES-192, or AES-256，返回密文和认证标签
// GCM 模式不需要手动填充，且会自动生成认证标签，解密需要同时提供密文和认证标签
// 安全性更高、并行加密解密性能更好
func AesGCMEncrypt(text string, key string, nonce string) (string, string, error) {
	var cipherHex string
	var tagHex string
	var err error

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v", r)
		}
	}()

	// 参数验证
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", "", errors.New("key length must be 16, 24, or 32 bytes for AES-128, AES-192, or AES-256")
	}

	// nonce 长度建议为 12 字节（96 位），这是 GCM 模式的最佳实践
	if len(nonce) != 12 {
		return "", "", errors.New("nonce length should be 12 bytes for optimal GCM performance and security")
	}

	textBytes := Bytes(text)
	keyBytes := Bytes(key)
	nonceBytes := Bytes(nonce)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", "", err
	}

	// 创建 GCM 模式的加密器
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", "", err
	}

	// 加密数据并生成认证标签
	cipherText := aesGCM.Seal(nil, nonceBytes, textBytes, nil)

	// 分离密文和认证标签
	tagSize := aesGCM.Overhead()
	cipherData := cipherText[:len(cipherText)-tagSize]
	tag := cipherText[len(cipherText)-tagSize:]

	cipherHex = hex.EncodeToString(cipherData)
	tagHex = hex.EncodeToString(tag)

	return cipherHex, tagHex, err
}

// AesGCMDecrypt Aes GCM 对称加密
func AesGCMDecrypt(cipherStr string, tagStr string, key string, nonce string) (string, error) {
	var result string
	var err error

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v", r)
		}
	}()

	// 参数验证
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", errors.New("key length must be 16, 24, or 32 bytes for AES-128, AES-192, or AES-256")
	}

	if len(nonce) != 12 {
		return "", errors.New("nonce length should be 12 bytes for optimal GCM performance and security")
	}

	if len(cipherStr) == 0 {
		return "", errors.New("cipher string cannot be empty")
	}

	if len(tagStr) == 0 {
		return "", errors.New("tag string cannot be empty")
	}

	cipherBytes, err := hex.DecodeString(cipherStr)
	if err != nil {
		return "", err
	}

	tagBytes, err := hex.DecodeString(tagStr)
	if err != nil {
		return "", err
	}

	keyBytes := Bytes(key)
	nonceBytes := Bytes(nonce)

	// 创建解密器
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	// 创建 GCM 模式的解密器
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// 组合密文和认证标签
	cipherText := append(cipherBytes, tagBytes...)

	// 解密数据并验证认证标签
	textBytes, err := aesGCM.Open(nil, nonceBytes, cipherText, nil)
	if err != nil {
		return "", err
	}

	result = String(textBytes)
	return result, err
}
