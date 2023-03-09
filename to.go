package fun

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"

	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
)

// Ip2Long 字符串 IP 转整型
func Ip2Long(ipStr string) uint32 {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	if ip == nil {
		return 0
	}

	return binary.BigEndian.Uint32(ip)
}

// Long2Ip 整型转字符串 IP
func Long2Ip(long uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, long)
	ip := net.IP(ipByte)

	return ip.String()
}

// ToString 将任意一个类型转换为字符串
func ToString(value any) string {
	return fmt.Sprintf("%v", value)
}

// ToInt 数字或字符串转 int 类型
func ToInt(value any) int {
	switch v := value.(type) {
	case int8:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int:
		return v
	case string:
		i, _ := strconv.Atoi(v)
		return i
	}

	return 0
}

// ToLong ToInt64 别名, 数字或字符串转 int64
func ToLong(value any) int64 {
	return ToInt64(value)
}

// ToBool 字符串转 bool 类型
func ToBool(str string) bool {
	b, err := strconv.ParseBool(str)
	if err != nil {
		return false
	} else {
		return b
	}
}

// ToUint 数字或字符串转 uint
func ToUint(value any) uint {
	switch v := value.(type) {
	case int8:
		return uint(v)
	case uint8:
		return uint(v)
	case uint16:
		return uint(v)
	case int16:
		return uint(v)
	case int32:
		return uint(v)
	case int:
		return uint(v)
	case uint:
		return v
	case string:
		i, _ := strconv.ParseUint(v, 10, 32)
		return uint(i)
	}

	return 0
}

// ToUint8 数字或字符串转 uint8
func ToUint8(value any) uint8 {
	switch v := value.(type) {
	case int8:
		return uint8(v)
	case uint8:
		return v
	case string:
		i, _ := strconv.ParseUint(v, 10, 8)
		return uint8(i)
	}

	return 0
}

// ToInt64 数字或字符串转 int64
func ToInt64(value any) int64 {
	switch v := value.(type) {
	case int:
		return int64(v)
	case uint8:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case uint16:
		return int64(v)
	case int32:
		return int64(v)
	case uint32:
		return int64(v)
	case int64:
		return v
	case string:
		i, _ := strconv.ParseInt(v, 10, 64)
		return i
	}

	return 0
}

// ToUtf8 指定字符集转 utf-8
func ToUtf8(origin []byte, encode string) ([]byte, error) {
	e, err := ianaindex.MIME.Encoding(encode)
	if err != nil {
		return nil, err
	}

	if e == nil {
		return nil, errors.New("unsupported encoding")
	}

	r := transform.NewReader(bytes.NewReader(origin), e.NewDecoder())
	s, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// Utf8To utf-8 转指定字符集
func Utf8To(utf8 []byte, encode string) ([]byte, error) {
	e, err := ianaindex.MIME.Encoding(encode)
	if err != nil {
		return nil, err
	}

	r := transform.NewReader(bytes.NewReader(utf8), e.NewEncoder())
	s, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// ToJson 将对象转换为 Json 字符串
func ToJson(object any) string {
	res, err := json.Marshal(object)
	if err != nil {
		return ""
	}

	return String(res)
}
