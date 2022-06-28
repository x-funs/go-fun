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
	"errors"
	"fmt"
	"math"
	"math/rand"
	"net"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
	"unsafe"

	"github.com/x-funs/go-fun/strtotime"
)

var (
	randomNew = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// GenInteger 整型范型集合
type GenInteger interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// GenFloat 浮点型范型集合
type GenFloat interface {
	float32 | float64
}

// GenNumber 数值范型集合
type GenNumber interface {
	GenInteger | GenFloat
}

// Timestamp 返回当前时间的 Unix 时间戳。
// 默认返回秒级，支持 Timestamp(true) 返回毫秒级
func Timestamp(millis ...any) int64 {
	l := len(millis)
	switch l {
	case 0:
		return unixTimestamp()
	case 1:
		switch v := millis[0].(type) {
		case bool:
			if v {
				return unixMilliTimestamp()
			}
		}
	}

	return unixTimestamp()
}

// unixTimestamp 返回当前时间的 Unix 秒级时间戳
func unixTimestamp() int64 {
	return time.Now().Unix()
}

// unixMilliTimestamp 返回当前时间的 Unix 毫秒级时间戳
func unixMilliTimestamp() int64 {
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

// Date 返回格式化后的日期时间字符串。
// 支持 Date()、Date(unixstamp)、Date(layout)、Date(layout, unixstamp)
func Date(layouts ...any) string {
	l := len(layouts)

	switch l {
	case 0:
		return dateByDefault()
	case 1:
		switch v := layouts[0].(type) {
		case string:
			return dateByPattern(ToString(v))
		case int, int8, int16, int32, int64:
			return dateByPatternAndTime("", ToInt64(v))
		}
	case 2:
		switch layouts[0].(type) {
		case string:
			switch v := layouts[1].(type) {
			case int, int8, int16, int32, int64:
				return dateByPatternAndTime(ToString(layouts[0]), ToInt64(v))
			}
		}
	}

	return ""
}

// dateByDefault 返回默认 layout 格式化后的日期时间字符串
func dateByDefault() string {
	now := time.Now()
	return now.Format(DatetimePattern)
}

// dateByPattern 返回指定 layout 格式化后的日期时间字符串
func dateByPattern(layout string) string {
	now := time.Now()

	if Blank(layout) {
		return now.Format(DatetimePattern)
	} else {
		return now.Format(layout)
	}
}

// dateByPatternAndTime 返回指定时间戳、指定 layout 格式化后的日期时间字符串
func dateByPatternAndTime(layout string, timeStamp int64) string {
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

// ToLong ToInt64 别名，数字或字符串转 int64
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

// Md5 返回字符串 Md5 值
func Md5(str string) string {
	hexStr := md5.Sum(StringToBytes(str))
	return hex.EncodeToString(hexStr[:])
}

// Sha1 返回字符串 Sha1 值
func Sha1(str string) string {
	hexStr := sha1.Sum(StringToBytes(str))
	return hex.EncodeToString(hexStr[:])
}

// Sha256 返回字符串 Sha256 值
func Sha256(str string) string {
	hexStr := sha256.Sum256(StringToBytes(str))
	return hex.EncodeToString(hexStr[:])
}

// Sha384 返回字符串 Sha384 值
func Sha384(str string) string {
	hexStr := sha512.Sum384(StringToBytes(str))
	return hex.EncodeToString(hexStr[:])
}

// Sha512 返回字符串 Sha512 值
func Sha512(str string) string {
	hexStr := sha512.Sum512(StringToBytes(str))
	return hex.EncodeToString(hexStr[:])
}

// Base64Encode 返回字符串 Base64 值
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString(StringToBytes(str))
}

// Base64Decode 返回 Base64 值对应的字符串
func Base64Decode(str string) string {
	decode, _ := base64.StdEncoding.DecodeString(str)
	return BytesToString(decode)
}

// Base64UrlEncode 返回字符串 Url Safe Base64 值
func Base64UrlEncode(str string) string {
	return base64.URLEncoding.EncodeToString(StringToBytes(str))
}

// Base64UrlDecode 返回 Url Safe Base64 值对应的字符串
func Base64UrlDecode(str string) string {
	decode, _ := base64.URLEncoding.DecodeString(str)
	return BytesToString(decode)
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

// Empty 判断是否为空，支持字符串、数值、数组、Slice、Map
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

// StrToTime 日期时间字符串转时间戳
// 支持 StrToTime()、StrToTime(string)、StrToTime(string, int64)
func StrToTime(args ...any) int64 {
	l := len(args)

	switch l {
	case 0:
		return Timestamp()
	case 1:
		exp := ToString(args[0])
		if !Blank(exp) {
			v, err := strtotime.Parse(exp, Timestamp())
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

// SplitTrim 分割字符串为字符串切片，对分割后的值进行 Trim ，并自动忽略空值
func SplitTrim(str, sep string) []string {
	if len(str) == 0 || len(sep) == 0 {
		return []string{}
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

// SplitTrimToInts 分割字符串为 int 切片，对分割后的值进行 Trim ，并自动忽略空值
func SplitTrimToInts(str, sep string) []int {
	if len(str) == 0 || len(sep) == 0 {
		return []int{}
	}

	ss := strings.Split(str, sep)
	if len(ss) == 0 {
		return []int{}
	}

	slices := make([]int, 0, len(ss))
	for i := range ss {
		s := strings.TrimSpace(ss[i])
		if len(s) > 0 {
			if n, err := strconv.Atoi(s); err == nil {
				slices = append(slices, n)
			}
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

// IsASCII 判断字符串是否全部 ASCII
func IsASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

// Contains 判断字符串是否包含指定的子串
func Contains(str, substr string) bool {
	return strings.Contains(str, substr)
}

// ContainsCase 判断字符串是否包含指定的子串，不区分大小写
func ContainsCase(str, substr string) bool {
	return Contains(strings.ToLower(str), strings.ToLower(substr))
}

// ContainsAny 判断字符串是否包含任意一个指定的多个子串
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

// Matches 判断字符串是否匹配指定的正则表达式
func Matches(str, pattern string) bool {
	match, _ := regexp.MatchString(pattern, str)
	return match
}

// SnakeToCamel 蛇形转驼峰
func SnakeToCamel(str string, bigCamel bool) string {
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
			if bigCamel {
				sb.WriteRune(unicode.ToUpper(r))
			} else {
				sb.WriteRune(unicode.ToLower(r))
			}
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

// CamelToSnake 驼峰转蛇形
func CamelToSnake(str string) string {
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

// Wrap 使用字符串包围原字符串
func Wrap(str string, wrapStr string) string {
	if len(str) == 0 || wrapStr == "" {
		return str
	}
	var sb strings.Builder
	sb.WriteString(wrapStr)
	sb.WriteString(str)
	sb.WriteString(wrapStr)

	return sb.String()
}

// Unwrap 去除字符串包围，非递归
func Unwrap(str string, wrapStr string) string {
	if str == "" || wrapStr == "" {
		return str
	}

	firstIndex := strings.Index(str, wrapStr)
	lastIndex := strings.LastIndex(str, wrapStr)

	if firstIndex == 0 && lastIndex > 0 && lastIndex <= len(str)-1 {
		if len(wrapStr) <= lastIndex {
			str = str[len(wrapStr):lastIndex]
		}
	}

	return str
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

// ToJson 将对象转换为 Json 字符串
func ToJson(object any) string {
	res, err := json.Marshal(object)
	if err != nil {
		return ""
	}
	return BytesToString(res)
}

// Reverse 反转字符串
func Reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Random 返回随机数 `[0, MaxInt)`
func Random() int {
	return randomNew.Intn(math.MaxInt)
}

// RandomInt 返回随机数 `[min, max)`
func RandomInt(min, max int) int {
	if min > max {
		min, max = max, min
	}

	return randomNew.Intn(max-min) + min
}

// RandomInt64 返回随机数 `[min, max)`
func RandomInt64(min, max int64) int64 {
	if min > max {
		min, max = max, min
	}

	return randomNew.Int63n(max-min) + min
}

// RandomString 返回指定长度的随机字符串，包含字母和数字
func RandomString(length int) string {
	return RandomPool(StringLetterAndNumber, length)
}

// RandomLetter 返回指定长度的随机字符串，仅包含字母
func RandomLetter(length int) string {
	return RandomPool(StringLetter, length)
}

// RandomNumber 返回指定长度的随机字符串，仅包含数字
func RandomNumber(length int) string {
	return RandomPool(StringNumber, length)
}

// RandomPool 从提供的字符串池中返回指定长度的随机字符串
func RandomPool(pool string, length int) string {
	if length <= 0 {
		return ""
	}
	var chars = StringToBytes(pool)
	var result []byte
	for i := 0; i < length; i++ {
		c := chars[RandomInt(0, len(chars))]
		result = append(result, c)
	}
	return BytesToString(result)
}

// Remove 移除字符串中指定的字符串
func Remove(str, remove string) string {
	if str == "" || remove == "" {
		return remove
	}
	return strings.Replace(str, remove, "", -1)
}

// RemovePrefix 左侧移除字符串中指定的字符串
func RemovePrefix(str, prefix string) string {
	if str == "" || prefix == "" {
		return str
	}
	return strings.TrimPrefix(str, prefix)
}

// RemoveSuffix 右侧移除字符串中指定的字符串
func RemoveSuffix(str string, suffix string) string {
	if str == "" || suffix == "" {
		return str
	}
	return strings.TrimSuffix(str, suffix)
}

// RemoveAny 移除字符串中指定的字符串集
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

// SliceContains 判断数值和字符串是否在切片中
func SliceContains[T GenNumber | string](v T, slice []T) bool {
	if len(slice) == 0 {
		return false
	}

	for _, s := range slice {
		if s == v {
			return true
		}
	}
	return false
}

// SliceUnique 对数值和字符串切片进行去重
func SliceUnique[T GenNumber | string](slice []T) []T {
	if len(slice) == 0 {
		return slice
	}
	m := make(map[T]bool)
	for i := range slice {
		m[slice[i]] = true
	}

	slice = slice[:0]
	for k := range m {
		slice = append(slice, k)
	}

	return slice
}

// SliceSplit 对数值和字符串切片按照指定长度进行分割
func SliceSplit[T GenNumber | string](slice []T, size int) [][]T {
	var res [][]T

	if len(slice) == 0 || size <= 0 {
		return res
	}

	length := len(slice)
	if size == 1 || size >= length {
		for _, v := range slice {
			var tmp []T
			tmp = append(tmp, v)
			res = append(res, tmp)
		}
		return res
	}

	// divide slice equally
	divideNum := length/size + 1
	for i := 0; i < divideNum; i++ {
		if i == divideNum-1 {
			if len(slice[i*size:]) > 0 {
				res = append(res, slice[i*size:])
			}
		} else {
			res = append(res, slice[i*size:(i+1)*size])
		}
	}

	return res
}

// SliceIndex 对数值和字符串切片按照指定值进行查找
func SliceIndex[T GenNumber | string](slice []T, v T) int {
	for i, s := range slice {
		if s == v {
			return i
		}
	}
	return -1
}

// SliceLastIndex 对数值和字符串切片按照指定值进行查找，返回最后一个匹配的索引
func SliceLastIndex[T GenNumber | string](slice []T, v T) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// MapKeys 返回map的键切片
func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

// MapValues 返回map的值切片
func MapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))

	for _, v := range m {
		values = append(values, v)
	}

	return values
}

// MapMerge 合并两个map，如果有相同的键，则后者会覆盖前者
func MapMerge[K comparable, V any](maps ...map[K]V) map[K]V {
	res := make(map[K]V, 0)

	for _, m := range maps {
		for k, v := range m {
			res[k] = v
		}
	}

	return res
}

// IsExist 文件或目录是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

// IsDir 是否是目录
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IntsToStrings int 切片转换为字符串切片
func IntsToStrings(slice []int) []string {
	if len(slice) == 0 {
		return []string{}
	}
	var str []string
	for _, v := range slice {
		str = append(str, strconv.Itoa(v))
	}
	return str
}

// StringsToInts 字符串切片转换为 int 切片
func StringsToInts(slice []string) []int {
	if len(slice) == 0 {
		return []int{}
	}
	var ints []int
	for _, v := range slice {
		if i, err := strconv.Atoi(v); err == nil {
			ints = append(ints, i)
		}
	}
	return ints
}

// StringToBytes 更高效的字符串转字节数组
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// BytesToString 更高效的字节数组转字符串
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Max 取 int 最大值
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min 取 int 最小值
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxInt64 取 int64 最大值
func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// MinInt64 取 int64 最小值
func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
