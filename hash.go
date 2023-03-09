package fun

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

// Md5 返回字符串 Md5 值
func Md5(str string) string {
	hexStr := md5.Sum(Bytes(str))

	return hex.EncodeToString(hexStr[:])
}

// Md5Bit16 返回 16位 字符串 Md5 值
func Md5Bit16(str string) string {
	s := Md5(str)
	return SubString(s, 8, 16)
}

// Sha1 返回字符串 Sha1 值
func Sha1(str string) string {
	hexStr := sha1.Sum(Bytes(str))

	return hex.EncodeToString(hexStr[:])
}

// Sha256 返回字符串 Sha256 值
func Sha256(str string) string {
	hexStr := sha256.Sum256(Bytes(str))

	return hex.EncodeToString(hexStr[:])
}

// Sha384 返回字符串 Sha384 值
func Sha384(str string) string {
	hexStr := sha512.Sum384(Bytes(str))

	return hex.EncodeToString(hexStr[:])
}

// Sha512 返回字符串 Sha512 值
func Sha512(str string) string {
	hexStr := sha512.Sum512(Bytes(str))

	return hex.EncodeToString(hexStr[:])
}

// Base64Encode 返回字符串 Base64 值
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString(Bytes(str))
}

// Base64Decode 返回 Base64 值对应的字符串
func Base64Decode(str string) string {
	decode, _ := base64.StdEncoding.DecodeString(str)

	return String(decode)
}

// Base64UrlEncode 返回字符串 Url Safe Base64 值
func Base64UrlEncode(str string) string {
	return base64.URLEncoding.EncodeToString(Bytes(str))
}

// Base64UrlDecode 返回 Url Safe Base64 值对应的字符串
func Base64UrlDecode(str string) string {
	decode, _ := base64.URLEncoding.DecodeString(str)

	return String(decode)
}
