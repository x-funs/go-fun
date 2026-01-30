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
	"strings"
	"time"

	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
)

// Ip2Long 字符串 IP 转整型
func Ip2Long(ipStr string) uint32 {
	result, _ := Ip2LongE(ipStr)
	return result
}

// Ip2LongE 字符串 IP 转整型，带错误返回
func Ip2LongE(ipStr string) (uint32, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return 0, errors.New("invalid IP address")
	}
	ip = ip.To4()
	if ip == nil {
		return 0, errors.New("invalid IPv4 address")
	}

	return binary.BigEndian.Uint32(ip), nil
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
	result, _ := ToIntE(value)
	return result
}

// ToIntE 数字或字符串转 int 类型，带错误返回
func ToIntE(value any) (int, error) {
	switch v := value.(type) {
	case int8:
		return int(v), nil
	case uint8:
		return int(v), nil
	case uint16:
		return int(v), nil
	case int16:
		return int(v), nil
	case int32:
		return int(v), nil
	case int:
		return v, nil
	case string:
		i, err := strconv.Atoi(v)
		return i, err
	}

	return 0, fmt.Errorf("unsupported type for ToInt: %T", value)
}

// ToLong ToInt64 别名, 数字或字符串转 int64
func ToLong(value any) int64 {
	return ToInt64(value)
}

// ToLongE ToInt64E 别名, 数字或字符串转 int64，带错误返回
func ToLongE(value any) (int64, error) {
	return ToInt64E(value)
}

// ToBool 字符串转 bool 类型
func ToBool(str string) bool {
	result, _ := ToBoolE(str)
	return result
}

// ToBoolE 字符串转 bool 类型，带错误返回
func ToBoolE(str string) (bool, error) {
	return strconv.ParseBool(str)
}

// ToUint 数字或字符串转 uint
func ToUint(value any) uint {
	result, _ := ToUintE(value)
	return result
}

// ToUintE 数字或字符串转 uint，带错误返回
func ToUintE(value any) (uint, error) {
	switch v := value.(type) {
	case int8:
		return uint(v), nil
	case uint8:
		return uint(v), nil
	case uint16:
		return uint(v), nil
	case int16:
		return uint(v), nil
	case int32:
		return uint(v), nil
	case int:
		return uint(v), nil
	case uint:
		return v, nil
	case string:
		i, err := strconv.ParseUint(v, 10, 32)
		return uint(i), err
	}

	return 0, fmt.Errorf("unsupported type for ToUint: %T", value)
}

// ToUint8 数字或字符串转 uint8
func ToUint8(value any) uint8 {
	result, _ := ToUint8E(value)
	return result
}

// ToUint8E 数字或字符串转 uint8，带错误返回
func ToUint8E(value any) (uint8, error) {
	switch v := value.(type) {
	case int8:
		return uint8(v), nil
	case uint8:
		return v, nil
	case string:
		i, err := strconv.ParseUint(v, 10, 8)
		return uint8(i), err
	}

	return 0, fmt.Errorf("unsupported type for ToUint8: %T", value)
}

// ToInt64 数字或字符串转 int64
func ToInt64(value any) int64 {
	result, _ := ToInt64E(value)
	return result
}

// ToInt64E 数字或字符串转 int64，带错误返回
func ToInt64E(value any) (int64, error) {
	switch v := value.(type) {
	case int:
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case int64:
		return v, nil
	case uint64:
		if v > uint64(9223372036854775807) {
			return 0, fmt.Errorf("uint64 value %d exceeds int64 maximum", v)
		}
		return int64(v), nil
	case string:
		i, err := strconv.ParseInt(v, 10, 64)
		return i, err
	}

	return 0, fmt.Errorf("unsupported type for ToInt64: %T", value)
}

// ToUint64 范型，数字或字符串转 uint64
func ToUint64(value any) uint64 {
	result, _ := ToUint64E(value)
	return result
}

// ToUint64E 范型，数字或字符串转 uint64，带错误返回
func ToUint64E(value any) (uint64, error) {
	switch v := (value).(type) {
	case int:
		return uint64(v), nil
	case uint8:
		return uint64(v), nil
	case int8:
		return uint64(v), nil
	case int16:
		return uint64(v), nil
	case uint16:
		return uint64(v), nil
	case int32:
		return uint64(v), nil
	case uint32:
		return uint64(v), nil
	case int64:
		return uint64(v), nil
	case uint64:
		return v, nil
	case string:
		i, err := strconv.ParseUint(v, 10, 64)
		return i, err
	}

	return 0, fmt.Errorf("unsupported type for ToUint64: %T", value)
}

// ToFloat32 数字或字符串转 float32
func ToFloat32(value any) float32 {
	result, _ := ToFloat32E(value)
	return result
}

// ToFloat32E 数字或字符串转 float32，带错误返回
func ToFloat32E(value any) (float32, error) {
	switch v := value.(type) {
	case int:
		return float32(v), nil
	case int8:
		return float32(v), nil
	case int16:
		return float32(v), nil
	case int32:
		return float32(v), nil
	case int64:
		return float32(v), nil
	case uint:
		return float32(v), nil
	case uint8:
		return float32(v), nil
	case uint16:
		return float32(v), nil
	case uint32:
		return float32(v), nil
	case float32:
		return v, nil
	case float64:
		return float32(v), nil
	case string:
		i, err := strconv.ParseFloat(v, 32)
		return float32(i), err
	}

	return 0, fmt.Errorf("unsupported type for ToFloat32: %T", value)
}

// ToFloat64 数字或字符串转 float64
func ToFloat64(value any) float64 {
	result, _ := ToFloat64E(value)
	return result
}

// ToFloat64E 数字或字符串转 float64，带错误返回
func ToFloat64E(value any) (float64, error) {
	switch v := value.(type) {
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case string:
		i, err := strconv.ParseFloat(v, 64)
		return i, err
	}

	return 0, fmt.Errorf("unsupported type for ToFloat64: %T", value)
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

// ToJsonIndent 将对象转换为 Json 字符串, 带缩进
func ToJsonIndent(object any) string {
	res, err := json.MarshalIndent(object, "", "  ")
	if err != nil {
		return ""
	}

	return String(res)
}

// ToDuration 将数字、字符串转换为 time.Duration，默认是 ns, 如果是字符串，支持 ns,ms,us,s,m,h
func ToDuration(value any) time.Duration {
	switch v := value.(type) {
	case time.Duration:
		return v
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8:
		return time.Duration(ToInt64(v))
	case float32, float64:
		return time.Duration(ToFloat64(v))
	case string:
		if strings.ContainsAny(v, "nsuµmh") {
			d, _ := time.ParseDuration(v)
			return d
		} else {
			d, _ := time.ParseDuration(v + "ns")
			return d
		}
	}

	return 0
}

// ToDurationMs 将数字、字符串转换为 time.Duration，默认是 ms, 如果是字符串，支持 ns,ms,us,s,m,h
func ToDurationMs(value any) time.Duration {
	switch v := value.(type) {
	case time.Duration:
		return v
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8:
		return ToDuration(value) * time.Millisecond
	case float32, float64:
		return ToDuration(value) * time.Millisecond
	case string:
		if strings.ContainsAny(v, "nsuµmh") {
			d, _ := time.ParseDuration(v)
			return d
		} else {
			d, _ := time.ParseDuration(v + "ms")
			return d
		}
	}

	return 0
}
