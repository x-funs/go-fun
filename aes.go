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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	textBytes := Bytes(text)
	keyBytes := Bytes(key)
	ivBytes := Bytes(iv)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", errors.New(err.Error())
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

	return cipherHex, nil
}

// AesCBCDecrypt Aes CBC 对称加密
func AesCBCDecrypt(cipherStr string, key string, iv string) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	cipherBytes, err := hex.DecodeString(cipherStr)
	if err != nil {
		return "", errors.New(err.Error())
	}

	keyBytes := Bytes(key)
	ivBytes := Bytes(iv)

	// 创建解密器
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", errors.New(err.Error())
	}

	// 创建解密块链
	mode := cipher.NewCBCDecrypter(block, ivBytes)

	// 解密数据
	textBytes := make([]byte, len(cipherBytes))
	mode.CryptBlocks(textBytes, cipherBytes)

	textBytes = pKCS7UnPadding(textBytes)

	return String(textBytes), nil
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
