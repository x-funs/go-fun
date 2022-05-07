package fun

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/x-funs/go-fun/strtotime"
)

var (
	randomNew = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// Timestamp 返回当前时间的 Unix 时间戳，默认返回秒级，支持 Timestamp(true) 返回毫秒级
func Timestamp(millis ...any) int64 {
	l := len(millis)
	switch l {
	case 0:
		return UnixTimestamp()
	case 1:
		switch v := millis[0].(type) {
		case bool:
			if v {
				return UnixMilliTimestamp()
			}
		}
	}

	return UnixTimestamp()
}

// UnixTimestamp 返回当前时间的 Unix 秒级时间戳
func UnixTimestamp() int64 {
	return time.Now().Unix()
}

// UnixMilliTimestamp 返回当前时间的 Unix 毫秒级时间戳
func UnixMilliTimestamp() int64 {
	return time.Now().UnixMilli()
}

// MemoryBytes 返回当前主要的内存指标信息
func MemoryBytes() map[string]int64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	maps := make(map[string]int64)
	maps["sys"] = int64(m.Sys)
	maps["heapSys"] = int64(m.HeapSys)
	maps["heapInuse"] = int64(m.HeapInuse)
	maps["heapAlloc"] = int64(m.HeapAlloc)

	return maps
}

// Date 返回格式化后的日期时间字符串，支持 Date()、Date(int)、Date(string)、Date(string, int)
func Date(layouts ...any) string {
	l := len(layouts)

	switch l {
	case 0:
		return DateByDefault()
	case 1:
		switch v := layouts[0].(type) {
		case string:
			return DateByPattern(ToString(v))
		case int, int8, int16, int32, int64:
			return DateByPatternAndTime("", ToInt64(v))
		}
	case 2:
		switch layouts[0].(type) {
		case string:
			switch v := layouts[1].(type) {
			case int, int8, int16, int32, int64:
				return DateByPatternAndTime(ToString(layouts[0]), ToInt64(v))
			}
		}
	}

	return ""
}

// DateByDefault 返回格式化后的日期时间字符串
func DateByDefault() string {
	now := time.Now()
	return now.Format(DatetimePattern)
}

// DateByPattern 返回指定格式化后的日期时间字符串
func DateByPattern(layout string) string {
	now := time.Now()

	if Blank(layout) {
		return now.Format(DatetimePattern)
	} else {
		return now.Format(layout)
	}
}

// DateByPatternAndTime 返回指定时间戳格式化后的日期时间字符串
func DateByPatternAndTime(layout string, timeStamp int64) string {
	if timeStamp < 0 {
		timeStamp = 0
	}
	uTime := time.Unix(timeStamp, 0)

	if Blank(layout) {
		return uTime.Format(DatetimePattern)
	} else {
		return uTime.Format(layout)
	}
}

// ToString 将任意一个类型转换为字符串
func ToString(value any) string {
	return fmt.Sprintf("%v", value)
}

// ToInt 数字和字符串转 Int
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

// ToLong 数字和字符串转 Int64
func ToLong(value any) int64 {
	return ToInt64(value)
}

// ToInt64 数字和字符串转 Int64
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

// Md5 返回字符串 Md5 值
func Md5(str string) string {
	hexStr := md5.Sum([]byte(str))
	return hex.EncodeToString(hexStr[:])
}

// Sha1 返回字符串 Sha1 值
func Sha1(str string) string {
	hexStr := sha1.Sum([]byte(str))
	return hex.EncodeToString(hexStr[:])
}

// Sha256 返回字符串 Sha256 值
func Sha256(str string) string {
	hexStr := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hexStr[:])
}

// Sha384 返回字符串 Sha384 值
func Sha384(str string) string {
	hexStr := sha512.Sum384([]byte(str))
	return hex.EncodeToString(hexStr[:])
}

// Sha512 返回字符串 Sha512 值
func Sha512(str string) string {
	hexStr := sha512.Sum512([]byte(str))
	return hex.EncodeToString(hexStr[:])
}

// Base64Encode 返回字符串 Base64 值
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Base64Decode 返回 Base64 对应的字符串
func Base64Decode(str string) string {
	decode, _ := base64.StdEncoding.DecodeString(str)
	return string(decode)
}

// Base64UrlEncode 返回字符串 Url Safe Base64 值
func Base64UrlEncode(str string) string {
	return base64.URLEncoding.EncodeToString([]byte(str))
}

// Base64UrlDecode 返回 Url Safe Base64 对应的字符串
func Base64UrlDecode(str string) string {
	decode, _ := base64.URLEncoding.DecodeString(str)
	return string(decode)
}

// BlankAll 判断 Trim 后的字符串集，是否全部为空白
func BlankAll(strs ...string) bool {
	if len(strs) == 0 {
		return true
	}

	for _, v := range strs {
		if !Blank(v) {
			return false
		}
	}

	return true
}

// BlankAny 判断 Trim 后的字符串集，是否任意一个包含空白
func BlankAny(strs ...string) bool {
	if len(strs) == 0 {
		return true
	}

	for _, v := range strs {
		if Blank(v) {
			return true
		}
	}

	return false
}

// Blank 判断 Trim 后的字符串，是否为空白
func Blank(str string) bool {
	t := strings.TrimSpace(str)

	if t == "" {
		return true
	}

	return false
}

// EmptyAll 判断是否全部为空
func EmptyAll(values ...any) bool {
	if len(values) == 0 {
		return true
	}

	for _, v := range values {
		if !Empty(v) {
			return false
		}
	}

	return true
}

// EmptyAny 判断是否任意一个为空
func EmptyAny(values ...any) bool {
	if len(values) == 0 {
		return true
	}

	for _, v := range values {
		if Empty(v) {
			return true
		}
	}

	return false
}

// Empty 判断是否为空，支持字符串、数值、数组、切片、Map
func Empty(value any) bool {
	if value == nil {
		return true
	}

	switch value.(type) {
	case string:
		return value == ""
	case int, int8, int16, int32, int64:
		return value == 0
	case uint, uint8, uint16, uint32, uint64:
		return value == 0
	case float32, float64:
		return value == 0
	case bool:
		return value == false
	default:
		r := reflect.ValueOf(value)
		switch r.Kind() {
		case reflect.Slice, reflect.Map:
			return r.Len() == 0 || r.IsNil()
		case reflect.Array:
			return r.Len() == 0
		case reflect.Ptr, reflect.Interface:
			return r.IsNil()
		}
	}

	return false
}

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

// StrToTime 日期时间字符串转时间戳，支持 StrToTime()、StrToTime(string)、StrToTime(string, int64)
func StrToTime(args ...any) int64 {
	l := len(args)

	switch l {
	case 0:
		return Timestamp()
	case 1:
		exp := ToString(args[0])
		if !Blank(exp) {
			v, err := strtotime.Parse(ToString(exp), Timestamp())
			if err == nil {
				return v
			}
		}
	case 2:
		exp := ToString(args[0])
		if !Blank(exp) {
			timeStamp := ToInt64(args[1])
			if timeStamp > 0 {
				v, err := strtotime.Parse(exp, timeStamp)
				if err == nil {
					return v
				}
			}
		}
	}

	return 0
}

// SplitTrim 分割字符串为切片，对分割后的值进行 Trim ，并自动忽略空值
func SplitTrim(str, sep string) []string {
	if len(str) == 0 {
		return []string{}
	}

	if len(sep) == 0 {
		sep = " "
	}

	ss := strings.Split(str, sep)
	if len(ss) == 0 {
		return []string{}
	}

	slices := make([]string, 0, len(ss))
	for i := range ss {
		s := strings.TrimSpace(ss[i])
		if len(s) > 0 {
			slices = append(slices, s)
		}
	}

	return slices
}

// IsNumber 判断字符串是否全部为数字
func IsNumber(str string) bool {
	if len(str) == 0 {
		return false
	}

	for _, r := range str {
		if !unicode.IsNumber(r) {
			return false
		}
	}

	return true
}

// IsLetter 判断字符串是否全部为字母
func IsLetter(str string) bool {
	if len(str) == 0 {
		return false
	}

	for _, r := range str {
		if !unicode.IsLetter(r) {
			return false
		}
	}

	return true
}

// Contains 判断字符串是否包含子串
func Contains(str, substr string) bool {
	return strings.Contains(str, substr)
}

// ContainsCase 判断字符串是否包含子串，不区分大小写
func ContainsCase(str, substr string) bool {
	return Contains(strings.ToLower(str), strings.ToLower(substr))
}

func ContainsAny(str string, substr ...string) bool {
	if len(str) == 0 || len(substr) == 0 {
		return false
	}

	for _, s := range substr {
		if Contains(str, s) {
			return true
		}
	}

	return false
}

// Matches 判断字符串是否匹配正则表达式
func Matches(str, pattern string) bool {
	match, _ := regexp.MatchString(pattern, str)
	return match
}

// UnderToCamel 下划线转大驼峰
func UnderToCamel(str string) string {
	if len(str) == 0 {
		return ""
	}

	if !Contains(str, "_") {
		return str
	}

	str = strings.ToLower(str)

	var sb strings.Builder
	sb.Grow(len(str))

	underscore := false
	for i, r := range str {
		if i == 0 {
			sb.WriteRune(unicode.ToUpper(r))
		} else if r == '_' {
			if i < len(str) {
				underscore = true
			}
		} else if underscore {
			sb.WriteRune(unicode.ToUpper(r))
			underscore = false
		} else {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}

// CamelToUnder 大驼峰转下划线
func CamelToUnder(str string) string {
	if len(str) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.Grow(len(str))

	for i, r := range str {
		if i == 0 {
			sb.WriteRune(unicode.ToLower(r))
		} else if unicode.IsUpper(r) {
			sb.WriteRune('_')
			sb.WriteRune(unicode.ToLower(r))
		} else {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}

// PadLeft 左侧填充字符串到指定长度
func PadLeft(str string, padStr string, padLen int) string {
	if len(str) >= padLen || padStr == "" {
		return str
	}
	return buildPadStr(str, padStr, padLen, true, false)
}

// PadRight 右侧填充字符串到指定长度
func PadRight(str string, padStr string, padLen int) string {
	if len(str) >= padLen || padStr == "" {
		return str
	}
	return buildPadStr(str, padStr, padLen, false, true)
}

// PadBoth 两侧填充字符串到指定长度
func PadBoth(str string, padStr string, padLen int) string {
	if len(str) >= padLen || padStr == "" {
		return str
	}
	return buildPadStr(str, padStr, padLen, true, true)
}

// buildPadStr
func buildPadStr(str string, padStr string, padLen int, padLeft bool, padRight bool) string {
	if padLen < utf8.RuneCountInString(str) {
		return str
	}

	padLen -= utf8.RuneCountInString(str)

	targetLen := padLen

	targetLenLeft := targetLen
	targetLenRight := targetLen
	if padLeft && padRight {
		targetLenLeft = padLen / 2
		targetLenRight = padLen - targetLenLeft
	}

	strToRepeatLen := utf8.RuneCountInString(padStr)

	repeatTimes := int(math.Ceil(float64(targetLen) / float64(strToRepeatLen)))
	repeatedString := strings.Repeat(padStr, repeatTimes)

	leftSide := ""
	if padLeft {
		leftSide = repeatedString[0:targetLenLeft]
	}

	rightSide := ""
	if padRight {
		rightSide = repeatedString[0:targetLenRight]
	}

	return leftSide + str + rightSide
}

// ToJson 将对象转换为json字符串
func ToJson(object any) string {
	res, err := json.Marshal(object)
	if err != nil {
		res = []byte("")
	}
	return string(res)
}

// Reverse 反转字符串
func Reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Random 返回随机数 [0, MaxInt)
func Random() int {
	return randomNew.Intn(math.MaxInt)
}

// RandomInt 返回随机数 [min, max)
func RandomInt(min, max int) int {
	if min > max {
		min, max = max, min
	}

	return randomNew.Intn(max-min) + min
}

// RandomInt64 返回随机数 [min, max)
func RandomInt64(min, max int64) int64 {
	if min > max {
		min, max = max, min
	}

	return randomNew.Int63n(max-min) + min
}

// RandomString 返回指定长度的随机字符串，包含大小写字母和数字
func RandomString(length int) string {
	return RandomPool(StringLetterAndNumber, length)
}

// RandomLetter 返回指定长度的随机字符串，包含大小写字母
func RandomLetter(length int) string {
	return RandomPool(StringLetter, length)
}

// RandomNumber 返回指定长度的随机字符串，包含数字
func RandomNumber(length int) string {
	return RandomPool(StringNumber, length)
}

// RandomPool 从字符串池中返回指定长度的随机字符串
func RandomPool(pool string, length int) string {
	if length <= 0 {
		return ""
	}
	var chars = []byte(pool)
	var result []byte
	for i := 0; i < length; i++ {
		c := chars[RandomInt(0, len(chars))]
		result = append(result, c)
	}
	return string(result)
}

// Remove 移除指定字符串中给定字符串
func Remove(str, remove string) string {
	if str == "" || remove == "" {
		return remove
	}
	return strings.Replace(str, remove, "", -1)
}

// RemovePrefix 移除指定字符串左侧给定字符串
func RemovePrefix(str, prefix string) string {
	if str == "" || prefix == "" {
		return str
	}
	return strings.TrimPrefix(str, prefix)
}

// RemoveSuffix 移除指定字符串右侧给定字符串
func RemoveSuffix(str string, suffix string) string {
	if str == "" || suffix == "" {
		return str
	}
	return strings.TrimSuffix(str, suffix)
}

// RemoveAny 移除指定字符串中给定字符串集
func RemoveAny(str string, removes ...string) string {
	if str == "" || len(removes) == 0 {
		return str
	}
	for _, rr := range removes {
		str = Remove(str, rr)
	}
	return str
}

// SubString 字符串截取
func SubString(str string, pos, length int) string {
	runes := []rune(str)
	max := len(runes)

	if pos < 0 || length <= 0 {
		return str
	}

	if pos > max {
		return ""
	}

	l := pos + length
	if l > max {
		l = max
	}
	return string(runes[pos:l])
}

func InSlice(str string, slice []string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
