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

## 函数说明

### 时间

#### `Timestamp(millis ...any) int64`

- 返回秒级时间戳

```
fmt.Println(fun.Timestamp())
// 1673225645

// 加参数 true 为毫秒级时间戳
fmt.Println(fun.Timestamp(true))
// 1673225645077
```

#### `Date(layouts ...any) string`

- 返回格式化后的日期时间字符串

```
fmt.Println(fun.Date())
// 2023-01-09 09:00:52

fmt.Println(fun.Date(1673225645))
// 2023-01-09 08:54:05

fmt.Println(fun.Date(fun.DateLayout))
// 2023-01-09

fmt.Println(fun.Date(fun.DateLayout, 1673225645))
// 2023-01-09
```

#### `StrToTime(args ...any) int64`

- 日期时间字符串转时间戳

```
fmt.Println(StrToTime())
// 1673226381

fmt.Println(StrToTime("-1 day"))
// 1673139981(一天前的时间戳)

fmt.Println(StrToTime("+1 day", 1673225645))
// 1673312045(某一时间戳一天后的时间戳)
```

### 辅助

#### `Empty(value any) bool`

- 判断是否为空, 支持字符串、数值、数组、Slice、Map

```
fmt.Println(fun.Empty(nil))
// true

fmt.Println(fun.Empty(0))
// true

fmt.Println(fun.Empty(""))
// true

fmt.Println(fun.Empty(false))
// true

fmt.Println(fun.Empty(" "))
// false

fmt.Println(fun.Empty(1))
// false

fmt.Println(fun.Empty(true))
// false

```

#### `EmptyAll(values ...any) bool`
- 判断是否全部为空

#### `EmptyAny(values ...any) bool`
- 判断是否任意一个为空

#### `MemoryBytes() map[string]int64`
- 返回当前主要的内存指标信息

#### `Memory(format string) map[string]int64`
- 指定格式返回当前主要的内存指标信息

#### `Bytes(s string) []byte`
- 更高效的字符串转字节数组

#### `String(b []byte) string`
- 更高效的字节数组转字符串

#### `Command(bin string, argv []string, baseDir string) ([]byte, error)`
- 执行系统命令

### 哈希

#### `Md5(str string) string`

- 返回字符串 Md5 值

```
fun.Md5("123456")
// e10adc3949ba59abbe56e057f20f883e
```

#### `Md5Bit16(str string) string`
- 返回 16位 字符串 Md5 值

#### `Sha1(str string) string`
- 返回字符串 Sha1 值

#### `Sha256(str string) string`
- 返回字符串 Sha256 值

#### `Sha384(str string) string`
- 返回字符串 Sha384 值

#### `Sha512(str string) string`
- 返回字符串 Sha512 值

#### `Base64Encode(str string) string`
- 返回字符串 Base64 值

#### `Base64Decode(str string) string`
- 返回 Base64 值对应的字符串

#### `Base64UrlEncode(str string) string`
- 返回字符串 Url Safe Base64 值

#### `Base64UrlDecode(str string) string`
- 返回 Url Safe Base64 值对应的字符串

### 判断

#### `IsNumber(str string) bool`
- 判断字符串是否全部为数字

#### `IsUtf8(p []byte) bool`
- 判断是否为 UTF-8 编码

#### `IsASCIILetter(str string) bool`
- 判断字符串是否全部为ASCII的字母

#### `IsLetter(str string) bool`
- 判断字符串是否全部为字母

#### `IsASCII(s string) bool`
- 判断字符串是否全部 ASCII

#### `IsEmail(str string) bool`
- 验证 Email 是否合法

#### `IsExist(path string) bool`
- 文件或目录是否存在

#### `IsDir(path string) bool`
- 是否是目录

### map 相关

#### `MapKeys[K comparable, V any](m map[K]V) []K`
- 返回map所有键的切片

#### `MapValues[K comparable, V any](m map[K]V) []V`
- 返回map所有值的切片

#### `MapMerge[K comparable, V any](maps ...map[K]V) map[K]V`
- 合并多个map, 如果有相同的键, 则后者会覆盖前者

### math

#### `Max(a, b int) int`
- 取 int 最大值

#### `Min(a, b int) int`
- 取 int 最小值

#### `MaxInt64(a, b int64) int64`
- 取 int64 最大值

#### `MinInt64(a, b int64) int64`
- 取 int64 最小值

### 随机

#### `Random() int`
- 返回随机数 `[0, MaxInt)`

#### `RandomInt(min, max int) int`
- 返回随机数 `[min, max)`

#### `RandomInt64(min, max int64) int64`
- 返回随机数 `[min, max)`

#### `RandomString(length int) string`
- 返回指定长度的随机字符串, 包含字母和数字

#### `RandomLetter(length int) string`
- 返回指定长度的随机字符串, 仅包含字母

#### `RandomNumber(length int) string`
- 返回指定长度的随机字符串, 仅包含数字

#### `RandomPool(pool string, length int) string`
- 从提供的字符串池中返回指定长度的随机字符串

### 正则

#### `Matches(str, pattern string) bool`
- 判断字符串是否匹配指定的正则表达式

### 相似度

#### `Similarity(a, b string) float64`
- 计算两个原始字符串的相似度

#### `SimilarityText(a, b string) float64`
- 计算两个字符串移除特殊符号后的相似度

#### `LongestCommonSubString(x, y string) int`
- 计算两个字符串最大公共子串长度

### 切片

#### `IntsToStrings(slice []int) []string`
- int 切片转换为字符串切片

#### `StringsToInts(slice []string) []int`
- 字符串切片转换为 int 切片

#### `SliceContains[T comparable](slice []T, v T) bool`
- 判断整型和字符串是否在切片中

#### `SliceUnique[T comparable](slice []T) []T`
- 对数值和字符串切片进行去重(会改变元素的顺序)

#### `SliceSplit[T comparable](slice []T, size int) [][]T`
- 对数值和字符串切片按照指定长度进行分割

```
fmt.Println(fun.SliceSplit([]string{"a", "b", "c", "d", "e", "f", "g"}, 3))
// [[a b c] [d e f] [g]]
```

#### `SliceIndex[T comparable](slice []T, v T) int`
- 对数值和字符串切片按照指定值进行查找

#### `SliceLastIndex[T comparable](slice []T, v T) int`
- 对数值和字符串切片按照指定值进行查找, 返回最后一个匹配的索引

#### `SliceRemove[T comparable](slice []T, v T) []T`
- 移除数值和字符串切片中的指定值

#### `SliceRemoveBlank(slice []string) []string`
- 移除字符串切片中的空值

#### `SliceTrim(slice []string) []string`
- 对字符串切片进行 Trim, 并自动忽略空值

#### `SliceConcat[T any](slice []T, values ...[]T) []T`
- 合并多个切片, 非去重, 非原始切片

#### `SliceEqual[T comparable](slice1, slice2 []T) bool`
- 切片是否相等: 长度相同且所有元素的顺序和值相等

#### `SliceEvery[T any](slice []T, predicate func(index int, item T) bool) bool`
- 切片中的所有元素都满足函数，则返回 true

#### `SliceNone[T any](slice []T, predicate func(index int, item T) bool) bool`
- 切片中的所有元素都不满足函数，则返回 true

#### `SliceSome[T any](slice []T, predicate func(index int, item T) bool) bool`
- 切片中有一个元素满足函数，就返回true

#### `SliceFilter[T any](slice []T, predicate func(index int, item T) bool) []T`
- 筛选出切片中满足函数的所有元素

#### `SliceForEach[T any](slice []T, iteratee func(index int, item T))`
- 切片中所有元素都执行函数

#### `SliceMap[T any, U any](slice []T, iteratee func(index int, item T) U) []U`
- 切片中所有元素都执行函数, 有返回值

#### `SliceReduce[T any](slice []T, iteratee func(index int, result, item T) T, initial T) T`
- 处理所有切片中元素得到结果

#### `SliceReplace[T comparable](slice []T, old T, new T, n int) []T`
- 返回切片的副本，前n个元素替换为新的

#### `SliceReplaceAll[T comparable](slice []T, old T, new T) []T`
- 返回切片的副本，所有匹配到的元素都替换为新的

#### `SliceUnion[T comparable](slices ...[]T) []T`
- 顺序合并且去重

```
fmt.Println(fun.SliceUnion([]string{"123", "124"}, []string{"124", "125"}, []string{"123", "125"}))
// [123 124 125]
```

#### `SliceUnionBy[T any, V comparable](predicate func(item T) V, slices ...[]T) []T`
- 顺序合并且去重, 支持自定义函数

#### `SliceIntersection[T comparable](slices ...[]T) []T`
- 切片交集且去重(顺序不能保证)

#### `SliceSortBy(slice any, field string, sortType ...string) error`
- 根据字段排序(field的大小写应该和字段保持一致)

#### `SliceColumn[T, V any](slice []T, key any) []V`
- 返回所有行的某一列

```
fmt.Println(
    SliceColumn[map[string]string, string]([]map[string]string{
        {"name": "衣服", "code": "YF4133"},
        {"name": "面膜", "code": "MM8541"},
        {"name": "口红", "code": "KH0002"},
        {"name": "手机", "code": "SJ9642"},
    }, "code")
)
// [YF4133 MM8541 KH0002 SJ9642]
```

### 字符串

#### `BlankAll(strs ...string) bool`
- 判断 Trim 后的字符串集, 是否全部为空白

#### `BlankAny(strs ...string) bool`
- 判断 Trim 后的字符串集, 是否任意一个包含空白

#### `Blank(str string) bool`
- 判断 Trim 后的字符串, 是否为空白

#### `HasPrefixCase(str, prefix string) bool`
- 判断字符串是否以指定前缀开头, 忽略大小写

#### `HasSuffixCase(str, prefix string) bool`
- 判断字符串是否以指定后缀结尾, 忽略大小写

#### `SplitTrim(str, sep string) []string`
- 分割字符串为字符串切片, 对分割后的值进行 Trim , 并自动忽略空值

#### `SplitTrimToInts(str, sep string) []int`
- 分割字符串为 int 切片, 对分割后的值进行 Trim , 并自动忽略空值

#### `Contains(str, substr string) bool`
- 判断字符串是否包含指定的子串

#### `ContainsCase(str, substr string) bool`
- 判断字符串是否包含指定的子串, 不区分大小写

#### `ContainsAny(str string, substr ...string) bool`
- 判断字符串是否包含任意一个指定的多个子串

#### `SnakeToCamel(str string, bigCamel bool) string`
- 蛇形转驼峰

#### `CamelToSnake(str string) string`
- 驼峰转蛇形

#### `PadLeft(str string, padStr string, padLen int) string`
- 左侧填充字符串到指定长度

#### `PadRight(str string, padStr string, padLen int) string`
- 右侧填充字符串到指定长度

#### `PadBoth(str string, padStr string, padLen int) string`
- 两侧填充字符串到指定长度

#### `Wrap(str string, wrapStr string) string`
- 使用字符串包围原字符串

#### `Unwrap(str string, wrapStr string) string`
- 去除字符串包围, 非递归

#### `Reverse(str string) string`
- 反转字符串

#### `Remove(str, remove string) string`
- 移除字符串中指定的字符串

#### `RemovePrefix(str, prefix string) string`
- 左侧移除字符串中指定的字符串

#### `RemoveSuffix(str string, suffix string) string`
- 右侧移除字符串中指定的字符串

#### `RemoveAny(str string, removes ...string) string`
- 移除字符串中指定的字符串集

#### `RemoveSign(str string) string`
- 将字符串的所有数据依次写成一行, 去除无意义字符串(标点符号、符号), 性能原因, 不使用 strings.NewReplacer

#### `RemoveLines(str string) string`
- 移除换行符, 换行符包括 \n \r\n, 性能原因, 不使用 strings.NewReplacer

#### `SubString(str string, pos, length int) string`
- 字符串截取

#### `NormaliseSpace(str string) string`
- 规范化此字符串中的空白, 多个空格合并为一个空格, 所有空白字符例如换行符、制表符, 都转换为一个简单的空格。

#### `NormaliseLine(str string) string`
- 规范化此字符串中的换行, 多个换行合并为一个换行

#### `Template(tpl string, data any) (string, error)`
- 模板渲染

#### `StrBefore(s, char string) string`

- 截取在字符首次出现时的位置之前的子字符串

```
fun.StrBefore("http://admin:123123@127.0.0.1:27017", ":")
// http
```

#### `StrBeforeLast(s, char string) string`

- 截取在字符最后出现时的位置之前的子字符串

```
fun.StrAfter("https://github.com", "://")
// github.com
```

#### `StrAfter(s, char string) string`

- 截取在字符首次出现时的位置之后的子字符串

```
fun.StrBeforeLast("video.mp4.bak", ".")
// video.mp4
```

#### `StrAfterLast(s, char string) string`

- 截取在字符最后出现时的位置之后的子字符串

```
fun.StrAfterLast("video.mp4.bak", ".")
// bak
```

### 结构体

#### `StructCopy(src, dst any)`
- 复制 struct 对象

### to 相关

#### `Ip2Long(ipStr string) uint32`
- 字符串 IP 转整型

#### `Long2Ip(long uint32) string`
- 整型转字符串 IP

#### `ToString(value any) string`
- 将任意一个类型转换为字符串

#### `ToInt(value any) int`
- 数字或字符串转 int 类型

#### `ToLong(value any) int64`
- ToInt64 别名, 数字或字符串转 int64

#### `ToBool(str string) bool`
- 字符串转 bool 类型

#### `ToUint(value any) uint`
- 数字或字符串转 uint

#### `ToUint8(value any) uint8`
- 数字或字符串转 uint8

#### `ToInt64(value any) int64`
- 数字或字符串转 int64

#### `ToUtf8(origin []byte, encode string) ([]byte, error)`
- 指定字符集转 utf-8

#### `Utf8To(utf8 []byte, encode string) ([]byte, error)`
- utf-8 转指定字符集

#### `ToJson(object any) string`
- 将对象转换为 Json 字符串

### http

> Get, Post, Put, Delete 四种方法
> 
> 带有 Resp 后缀的, 返回值是 *Response
> 
> 带有 Do 后缀的, 需要传参数 *Request
> 
> post form, post json 等此类命名的, 请求数据的类型不一样

#### `HttpGet(urlStr string, args ...any) ([]byte, error)`
- HttpGet 参数为请求地址 (HttpReq, 超时时间)
- HttpGet(url)、HttpGet(url, HttpReq)、HttpGet(url, timeout)、HttpGet(url, HttpReq, timeout)

#### `HttpGetResp(urlStr string, r *HttpReq, timeout int) (*HttpResp, error)`
- HttpGetResp Http Get 请求, 参数为请求地址, HttpReq, 超时时间(毫秒)

#### `HttpGetDo(urlStr string, r *HttpReq, timeout int) ([]byte, error)`
- HttpGetDo Http Get 请求, 参数为请求地址, HttpReq, 超时时间(毫秒)

#### `HttpDelete(urlStr string, args ...any) ([]byte, error)`
- HttpDelete 参数为请求地址 (HttpReq, 超时时间)
- HttpDelete(url)、HttpDelete(url, HttpReq)、HttpDelete(url, timeout)、HttpDelete(url, HttpReq, timeout)

#### `HttpDeleteDo(urlStr string, r *HttpReq, timeout int) ([]byte, error)`
- HttpDeleteDo Http Delete 请求, 参数为请求地址, HttpReq, 超时时间(毫秒)

#### `HttpDeleteResp(urlStr string, r *HttpReq, timeout int) (*HttpResp, error)`
- HttpDeleteResp Http Delete 请求, 参数为请求地址, HttpReq, 超时时间(毫秒)

#### `HttpPost(urlStr string, args ...any) ([]byte, error)`
- HttpPost 参数为请求地址 (body io.Reader, HttpReq, 超时时间)
- HttpPost(url)、HttpPost(url, timeout)、HttpPost(url, body)、HttpPost(url, body, timeout)、HttpPost(url, body, HttpReq)、HttpPostForm(url, body, HttpReq, timeout)

#### `HttpPostForm(urlStr string, args ...any) ([]byte, error)`
- HttpPostForm 参数为请求地址 (Form 数据 map[string]string, HttpReq, 超时时间)
- HttpPostForm(url)、HttpPostForm(url, timeout)、HttpPostForm(url, posts)、HttpPostForm(url, posts, timeout)、HttpPostForm(url, posts, HttpReq)、HttpPostForm(url, posts, HttpReq, timeout)

#### `HttpPostJson(urlStr string, args ...any) ([]byte, error)`
- HttpPostJson 参数为请求地址 (Json 数据 string, HttpReq, 超时时间)
- HttpPostJson(url)、HttpPostJson(url, timeout)、HttpPostJson(url, json)、HttpPost(url, json, timeout)、HttpPost(url, json, HttpReq)、HttpPost(url, json, HttpReq, timeout)

#### `HttpPostDo(urlStr string, body io.Reader, r *HttpReq, timeout int) ([]byte, error)`

#### `HttpPostFormDo(urlStr string, posts map[string]string, r *HttpReq, timeout int) ([]byte, error)`

#### `HttpPostJsonDo(urlStr string, json string, r *HttpReq, timeout int) ([]byte, error)`

#### `HttpPostResp(urlStr string, body io.Reader, r *HttpReq, timeout int) (*HttpResp, error)`

#### `HttpPostFormResp(urlStr string, posts map[string]string, r *HttpReq, timeout int) (*HttpResp, error)`

#### `HttpPostJsonResp(urlStr string, json string, r *HttpReq, timeout int) (*HttpResp, error)`

#### `HttpPut(urlStr string, args ...any) ([]byte, error)`
- HttpPut 参数为请求地址 (body io.Reader, HttpReq, 超时时间)
- HttpPut(url)、HttpPut(url, timeout)、HttpPut(url, body)、HttpPut(url, body, timeout)、HttpPut(url, body, HttpReq)、HttpPut(url, body, HttpReq, timeout)

#### `HttpPutForm(urlStr string, args ...any) ([]byte, error)`
- HttpPutForm 参数为请求地址 (Form 数据 map[string]string, HttpReq, 超时时间)
- HttpPutForm(url)、HttpPutForm(url, timeout)、HttpPutForm(url, posts)、HttpPutForm(url, posts, timeout)、HttpPutForm(url, posts, HttpReq)、HttpPutForm(url, posts, HttpReq, timeout)

#### `HttpPutJson(urlStr string, args ...any) ([]byte, error)`
- HttpPutJson 参数为请求地址 (Json 数据 string, HttpReq, 超时时间)
- HttpPutJson(url)、HttpPutJson(url, timeout)、HttpPutJson(url, json)、HttpPutJson(url, json, timeout)、HttpPutJson(url, json, httpReq)、HttpPutJson(url, json, httpReq, timeout)

#### `HttpPutDo(urlStr string, body io.Reader, r *HttpReq, timeout int) ([]byte, error)`

#### `HttpPutFormDo(urlStr string, posts map[string]string, r *HttpReq, timeout int) ([]byte, error)`

#### `HttpPutJsonDo(urlStr string, json string, r *HttpReq, timeout int) ([]byte, error)`

#### `HttpPutResp(urlStr string, body io.Reader, r *HttpReq, timeout int) (*HttpResp, error)`

#### `HttpPutFormResp(urlStr string, posts map[string]string, r *HttpReq, timeout int) (*HttpResp, error)`

#### `HttpPutJsonResp(urlStr string, json string, r *HttpReq, timeout int) (*HttpResp, error)`

#### `HttpDo(req *http.Request, r *HttpReq, timeout int) ([]byte, error)`
- HttpDo Http 请求, 参数为 http.Request, HttpReq, 超时时间(毫秒)

#### `HttpDoResp(req *http.Request, r *HttpReq, timeout int) (*HttpResp, error)`
- HttpDoResp Http 请求, 参数为 http.Request, HttpReq, 超时时间(毫秒)

#### `UrlParse(rawURL string) (*url.URL, error)`
- 解析字符串 URL 到 URL 对象。在没有 scheme 时不会出错


