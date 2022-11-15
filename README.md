# Go With Fun (Functions)

Go with functions is a small golang tools and utils library.

## Install

```shell
go get -u github.com/x-funs/go-fun
```

## Example

```go
package main

import (
	"fmt"

	"github.com/x-funs/go-fun"
)

func main() {
	// whether the string is empty after trim (blank or not)
	fmt.Println(fun.Blank("  "))
	
	// return string md5 hash
	fmt.Println(fun.Md5("go-fun"))
	
	// auto parse many datetime string to long timestamp
	fmt.Println(fun.StrToTime("2015-04-06 16:03:03"))

	// return slice unique
	fmt.Println(fun.SliceUnique([]string{"a", "b", "c", "a", "b", "c"}))

	// send a simple http and get html body
	html, _ := fun.HttpGet("https://www.163.com")
	fmt.Println(fun.String(html))
}
```

## Hash

func Base64Decode(str string) string
Base64Decode 返回 Base64 值对应的字符串

func Base64Encode(str string) string
Base64Encode 返回字符串 Base64 值

func Base64UrlDecode(str string) string
Base64UrlDecode 返回 Url Safe Base64 值对应的字符串

func Base64UrlEncode(str string) string
Base64UrlEncode 返回字符串 Url Safe Base64 值
*fun

func Blank(str string) bool
Blank 判断 Trim 后的字符串, 是否为空白

func BlankAll(strs ...string) bool
BlankAll 判断 Trim 后的字符串集, 是否全部为空白

func BlankAny(strs ...string) bool
BlankAny 判断 Trim 后的字符串集, 是否任意一个包含空白

func Bytes(s string) []byte
Bytes 更高效的字符串转字节数组

## string

func CamelToSnake(str string) string
CamelToSnake 驼峰转蛇形

## fun
func Command(bin string, argv []string, baseDir string) ([]byte, error)
Command 执行系统命令

## string
func Contains(str, substr string) bool
Contains 判断字符串是否包含指定的子串

func ContainsAny(str string, substr ...string) bool
ContainsAny 判断字符串是否包含任意一个指定的多个子串

func ContainsCase(str, substr string) bool
ContainsCase 判断字符串是否包含指定的子串, 不区分大小写
## datetime

func Date(layouts ...any) string
Date 返回格式化后的日期时间字符串。 支持 Date()、Date(unixstamp)、Date(layout)、Date(layout,
unixstamp)

func Empty(value any) bool
Empty 判断是否为空, 支持字符串、数值、数组、Slice、Map

func EmptyAll(values ...any) bool
EmptyAll 判断是否全部为空

func EmptyAny(values ...any) bool
EmptyAny 判断是否任意一个为空

func HasPrefixCase(str, prefix string) bool
HasPrefixCase 判断字符串是否以指定前缀开头, 忽略大小写

func HasSuffixCase(str, prefix string) bool
HasSuffixCase 判断字符串是否以指定后缀结尾, 忽略大小写

func HttpDelete(urlStr string, args ...any) ([]byte, error)
HttpDelete 参数为请求地址 (HttpReq, 超时时间) HttpDelete(url)、HttpDelete(url,
HttpReq)、HttpDelete(url, timeout)、HttpDelete(url, HttpReq, timeout) 返回 body,
错误信息

func HttpDeleteDo(urlStr string, r *HttpReq, timeout int) ([]byte, error)
HttpDeleteDo Http Delete 请求, 参数为请求地址, HttpReq, 超时时间(毫秒) 返回 body, 错误信息

func HttpDo(req *http.Request, r *HttpReq, timeout int) ([]byte, error)
HttpDo Http 请求, 参数为 http.Request, HttpReq, 超时时间(毫秒) 返回 body, 错误信息

func HttpGet(urlStr string, args ...any) ([]byte, error)
HttpGet 参数为请求地址 (HttpReq, 超时时间) HttpGet(url)、HttpGet(url,
HttpReq)、HttpGet(url, timeout)、HttpGet(url, HttpReq, timeout) 返回 body, 错误信息

func HttpGetDo(urlStr string, r *HttpReq, timeout int) ([]byte, error)
HttpGetDo Http Get 请求, 参数为请求地址, HttpReq, 超时时间(毫秒) 返回 body, 错误信息

func HttpPost(urlStr string, args ...any) ([]byte, error)
HttpPost 参数为请求地址 (body io.Reader, HttpReq, 超时时间) HttpPost(url)、HttpPost(url,
timeout)、HttpPost(url, body)、HttpPost(url, body, timeout)、HttpPost(url,
body, HttpReq)、HttpPostForm(url, body, HttpReq, timeout) 返回 body, 错误信息

func HttpPostDo(urlStr string, body io.Reader, r *HttpReq, timeout int) ([]byte, error)
HttpPostDo Http Post, 参数为请求地址, body io.Reader, HttpReq, 超时时间(毫秒) 返回 body,
错误信息

func HttpPostForm(urlStr string, args ...any) ([]byte, error)
HttpPostForm 参数为请求地址 (Form 数据 map[string]string, HttpReq, 超时时间)
HttpPostForm(url)、HttpPostForm(url, timeout)、HttpPostForm(url,
posts)、HttpPostForm(url, posts, timeout)、HttpPostForm(url, posts,
HttpReq)、HttpPostForm(url, posts, HttpReq, timeout) 返回 body, 错误信息

func HttpPostFormDo(urlStr string, posts map[string]string, r *HttpReq, timeout int) ([]byte, error)
HttpPostFormDo Http Post Form, 参数为请求地址, Form 数据 map[string]string, HttpReq,
超时时间(毫秒) 返回 body, 错误信息

func HttpPostJson(urlStr string, args ...any) ([]byte, error)
HttpPostJson 参数为请求地址 (Json 数据 string, HttpReq, 超时时间)
HttpPostJson(url)、HttpPostJson(url, timeout)、HttpPostJson(url,
json)、HttpPost(url, json, timeout)、HttpPost(url, json,
HttpReq)、HttpPost(url, json, HttpReq, timeout) 返回 body, 错误信息

func HttpPostJsonDo(urlStr string, json string, r *HttpReq, timeout int) ([]byte, error)
HttpPostJsonDo Http Post Json 请求, 参数为请求地址, Json 数据 string, HttpReq, 超时时间(毫秒)
返回 body, 错误信息

func HttpPut(urlStr string, args ...any) ([]byte, error)
HttpPut 参数为请求地址 (body io.Reader, HttpReq, 超时时间) HttpPut(url)、HttpPut(url,
timeout)、HttpPut(url, body)、HttpPut(url, body, timeout)、HttpPut(url, body,
HttpReq)、HttpPut(url, body, HttpReq, timeout) 返回 body, 错误信息

func HttpPutDo(urlStr string, body io.Reader, r *HttpReq, timeout int) ([]byte, error)
HttpPutDo Http Put, 参数为请求地址, body io.Reader, HttpReq, 超时时间(毫秒) 返回 body, 错误信息

func HttpPutForm(urlStr string, args ...any) ([]byte, error)
HttpPutForm 参数为请求地址 (Form 数据 map[string]string, HttpReq, 超时时间)
HttpPutForm(url)、HttpPutForm(url, timeout)、HttpPutForm(url,
posts)、HttpPutForm(url, posts, timeout)、HttpPutForm(url, posts,
HttpReq)、HttpPutForm(url, posts, HttpReq, timeout) 返回 body, 错误信息

func HttpPutFormDo(urlStr string, posts map[string]string, r *HttpReq, timeout int) ([]byte, error)
HttpPutFormDo Http Put Form, 参数为请求地址, Form 数据 map[string]string, HttpReq,
超时时间(毫秒) 返回 body, 错误信息

func HttpPutJson(urlStr string, args ...any) ([]byte, error)
HttpPutJson 参数为请求地址 (Json 数据 string, HttpReq, 超时时间)
HttpPutJson(url)、HttpPutJson(url, timeout)、HttpPutJson(url,
json)、HttpPutJson(url, json, timeout)、HttpPutJson(url, json,
httpReq)、HttpPutJson(url, json, httpReq, timeout) 返回 body, 错误信息

func HttpPutJsonDo(urlStr string, json string, r *HttpReq, timeout int) ([]byte, error)
HttpPutJsonDo Http Put Json 请求, 参数为请求地址, Json 数据 string, HttpReq, 超时时间(毫秒)
返回 body, 错误信息

func IntsToStrings(slice []int) []string
IntsToStrings int 切片转换为字符串切片
## is
func Ip2Long(ipStr string) uint32
Ip2Long 字符串 IP 转整型

func IsASCII(s string) bool
IsASCII 判断字符串是否全部 ASCII

func IsASCIILetter(str string) bool
IsASCIILetter 判断字符串是否全部为ASCII的字母

func IsDir(path string) bool
IsDir 是否是目录

func IsEmail(str string) bool
IsEmail 验证 Email 是否合法

func IsExist(path string) bool
IsExist 文件或目录是否存在

func IsLetter(str string) bool
IsLetter 判断字符串是否全部为字母

func IsNumber(str string) bool
IsNumber 判断字符串是否全部为数字

func IsUtf8(p []byte) bool
IsUtf8 判断是否为 UTF-8 编码

func Long2Ip(long uint32) string
Long2Ip 整型转字符串 IP
## similarity

func LongestCommonSubString(x, y string) int
LongestCommonSubString 计算两个字符串最大公共子串长度

func MapKeys[K comparable, V any](m map[K]V) []K
MapKeys 返回map的键切片

func MapMerge[K comparable, V any](maps ...map[K]V) map[K]V
MapMerge 合并多个map, 如果有相同的键, 则后者会覆盖前者

func MapValues[K comparable, V any](m map[K]V) []V
MapValues 返回map的值切片

func Matches(str, pattern string) bool
Matches 判断字符串是否匹配指定的正则表达式
## math
func Max(a, b int) int
Max 取 int 最大值

func MaxInt64(a, b int64) int64
MaxInt64 取 int64 最大值

func Md5(str string) string
Md5 返回字符串 Md5 值

func Md5Bit16(str string) string
Md5Bit16 返回 16位 字符串 Md5 值

func Memory(format string) map[string]int64
Memory 指定格式返回当前主要的内存指标信息, (ReadMemStats 会 stopTheWorld, 谨慎非频繁使用)

func MemoryBytes() map[string]int64
MemoryBytes 返回当前主要的内存指标信息 (ReadMemStats 会 stopTheWorld, 谨慎非频繁使用)

func Min(a, b int) int
Min 取 int 最小值

func MinInt64(a, b int64) int64
MinInt64 取 int64 最小值
## string
func NormaliseLine(str string) string
NormaliseLine 规范化此字符串中的换行, 多个换行合并为一个换行

func NormaliseSpace(str string) string
NormaliseSpace 规范化此字符串中的空白, 多个空格合并为一个空格, 所有空白字符例如换行符、制表符, 都转换为一个简单的空格。

func PadBoth(str string, padStr string, padLen int) string
PadBoth 两侧填充字符串到指定长度

func PadLeft(str string, padStr string, padLen int) string
PadLeft 左侧填充字符串到指定长度

func PadRight(str string, padStr string, padLen int) string
PadRight 右侧填充字符串到指定长度
## random
func Random() int
Random 返回随机数 `[0, MaxInt)`

func RandomInt(min, max int) int
RandomInt 返回随机数 `[min, max)`

func RandomInt64(min, max int64) int64
RandomInt64 返回随机数 `[min, max)`

func RandomLetter(length int) string
RandomLetter 返回指定长度的随机字符串, 仅包含字母

func RandomNumber(length int) string
RandomNumber 返回指定长度的随机字符串, 仅包含数字

func RandomPool(pool string, length int) string
RandomPool 从提供的字符串池中返回指定长度的随机字符串

func RandomString(length int) string
RandomString 返回指定长度的随机字符串, 包含字母和数字

func Remove(str, remove string) string
Remove 移除字符串中指定的字符串

func RemoveAny(str string, removes ...string) string
RemoveAny 移除字符串中指定的字符串集

func RemoveLines(str string) string
RemoveLines 移除换行符, 换行符包括 \n \r\n, 性能原因, 不使用 strings.NewReplacer

func RemovePrefix(str, prefix string) string
RemovePrefix 左侧移除字符串中指定的字符串

func RemoveSign(str string) string
RemoveSign 将字符串的所有数据依次写成一行, 去除无意义字符串(标点符号、符号), 性能原因, 不使用 strings.NewReplacer

func RemoveSuffix(str string, suffix string) string
RemoveSuffix 右侧移除字符串中指定的字符串

func Reverse(str string) string
Reverse 反转字符串

func Sha1(str string) string
Sha1 返回字符串 Sha1 值

func Sha256(str string) string
Sha256 返回字符串 Sha256 值

func Sha384(str string) string
Sha384 返回字符串 Sha384 值

func Sha512(str string) string
Sha512 返回字符串 Sha512 值

func Similarity(a, b string) float64
Similarity 计算两个原始字符串的相似度

func SimilarityText(a, b string) float64
SimilarityText 计算两个字符串移除特殊符号后的相似度
## slice

func SliceContains[T GenInteger | string](slice []T, v T) bool
SliceContains 判断整型和字符串是否在切片中

func SliceIndex[T GenInteger | string](slice []T, v T) int
SliceIndex 对数值和字符串切片按照指定值进行查找

func SliceLastIndex[T GenInteger | string](slice []T, v T) int
SliceLastIndex 对数值和字符串切片按照指定值进行查找, 返回最后一个匹配的索引

func SliceRemove[T GenInteger | string](slice []T, v T) []T
SliceRemove 移除数值和字符串切片中的指定值

func SliceRemoveBlank(slice []string) []string
SliceRemoveBlank 移除字符串切片中的空值

func SliceSplit[T GenInteger | string](slice []T, size int) [][]T
SliceSplit 对数值和字符串切片按照指定长度进行分割

func SliceTrim(slice []string) []string
SliceTrim 对字符串切片进行 Trim, 并自动忽略空值

func SliceUnique[T GenInteger | string](slice []T) []T
SliceUnique 对数值和字符串切片进行去重

func SnakeToCamel(str string, bigCamel bool) string
SnakeToCamel 蛇形转驼峰

func SplitTrim(str, sep string) []string
SplitTrim 分割字符串为字符串切片, 对分割后的值进行 Trim , 并自动忽略空值

func SplitTrimToInts(str, sep string) []int
SplitTrimToInts 分割字符串为 int 切片, 对分割后的值进行 Trim , 并自动忽略空值

func StrToTime(args ...any) int64
StrToTime 日期时间字符串转时间戳 支持 StrToTime()、StrToTime(string)、StrToTime(string,
int64)

func String(b []byte) string
String 更高效的字节数组转字符串

func StringsToInts(slice []string) []int
StringsToInts 字符串切片转换为 int 切片

func StructCopy(src, dst any)
StructCopy 复制 struct 对象

func SubString(str string, pos, length int) string
SubString 字符串截取

func Template(tpl string, data any) (string, error)
Template 模板渲染

func Timestamp(millis ...any) int64
Timestamp 返回当前时间的 Unix 时间戳。 默认返回秒级, 支持 Timestamp(true) 返回毫秒级
## to
func ToBool(str string) bool
ToBool 字符串转 bool 类型

func ToInt(value any) int
ToInt 数字或字符串转 int 类型

func ToInt64(value any) int64
ToInt64 数字或字符串转 int64

func ToJson(object any) string
ToJson 将对象转换为 Json 字符串

func ToLong(value any) int64
ToLong ToInt64 别名, 数字或字符串转 int64

func ToString(value any) string
ToString 将任意一个类型转换为字符串

func ToUint(value any) uint
ToUint 数字或字符串转 uint

func ToUint8(value any) uint8
ToUint8 数字或字符串转 uint8

func ToUtf8(origin []byte, encode string) ([]byte, error)
ToUtf8 指定字符集转 utf-8

func Unwrap(str string, wrapStr string) string
Unwrap 去除字符串包围, 非递归
## http

func UrlParse(rawURL string) (*url.URL, error)
UrlParse url.Parse 在没有 scheme 时不会出错
## to
func Utf8To(utf8 []byte, encode string) ([]byte, error)
Utf8To utf-8 转指定字符集

func Wrap(str string, wrapStr string) string
Wrap 使用字符串包围原字符串


## http

func HttpDeleteResp(urlStr string, r *HttpReq, timeout int) (*HttpResp, error)
HttpDeleteResp Http Delete 请求, 参数为请求地址, HttpReq, 超时时间(毫秒) 返回 HttpResp, 错误信息

func HttpDoResp(req *http.Request, r *HttpReq, timeout int) (*HttpResp, error)
HttpDoResp Http 请求, 参数为 http.Request, HttpReq, 超时时间(毫秒) 返回 HttpResp, 错误信息

func HttpGetResp(urlStr string, r *HttpReq, timeout int) (*HttpResp, error)
HttpGetResp Http Get 请求, 参数为请求地址, HttpReq, 超时时间(毫秒) 返回 HttpResp, 错误信息

func HttpPostFormResp(urlStr string, posts map[string]string, r *HttpReq, timeout int) (*HttpResp, error)
HttpPostFormResp Http Post Form, 参数为请求地址, Form 数据 map[string]string,
HttpReq, 超时时间(毫秒) 返回 HttpResp, 错误信息

func HttpPostJsonResp(urlStr string, json string, r *HttpReq, timeout int) (*HttpResp, error)
HttpPostJsonResp Http Post Json 请求, 参数为请求地址, Json 数据 string, HttpReq,
超时时间(毫秒) 返回 HttpResp, 错误信息

func HttpPostResp(urlStr string, body io.Reader, r *HttpReq, timeout int) (*HttpResp, error)
HttpPostResp Http Post, 参数为请求地址, body io.Reader, HttpReq, 超时时间(毫秒) 返回
HttpResp, 错误信息

func HttpPutFormResp(urlStr string, posts map[string]string, r *HttpReq, timeout int) (*HttpResp, error)
HttpPutFormResp Http Put Form, 参数为请求地址, Form 数据 map[string]string, HttpReq,
超时时间(毫秒) 返回 HttpResp, 错误信息

func HttpPutJsonResp(urlStr string, json string, r *HttpReq, timeout int) (*HttpResp, error)
HttpPutJsonResp Http Put Json 请求, 参数为请求地址, Json 数据 string, HttpReq, 超时时间(毫秒)
返回 HttpResp, 错误信息

func HttpPutResp(urlStr string, body io.Reader, r *HttpReq, timeout int) (*HttpResp, error)
HttpPutResp Http Put, 参数为请求地址, body io.Reader, HttpReq, 超时时间(毫秒) 返回
HttpResp, 错误信息
