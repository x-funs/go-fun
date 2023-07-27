```
                      ____          
   ____ _____        / __/_  ______ 
  / __ `/ __ \______/ /_/ / / / __ \
 / /_/ / /_/ /_____/ __/ /_/ / / / /
 \__, /\____/     /_/  \__,_/_/ /_/ 
/____/                              

```

Go with Fun (Functions) 是一个短小能干的 Golang 工具函数库。它包含 Empty、Blank、Strtotime、Similarity、HttpGet 等常用函数。

简体中文 | [English](./README.md)

## 使用

```shell
go get -u github.com/x-funs/go-fun
```

## 示例

```go
package main

import (
	"fmt"

	"github.com/x-funs/go-fun"
)

func main() {
	// 判断 any 是否为空
	fmt.Println(fun.Empty(""))
	
	// 判断字符串是否为空白
	fmt.Println(fun.Blank("  "))
	
	// 返回 MD5 字符串
	fmt.Println(fun.Md5("go-fun"))
	
	// 自动解析时间格式为秒时间戳
	fmt.Println(fun.StrToTime("2015-04-06 16:03:03"))
	fmt.Println(fun.StrToTime("2015/04/06 16:03:03"))
	fmt.Println(fun.StrToTime("2015年04月06日 16时03分03秒"))

	// 切片过滤排重
	fmt.Println(fun.SliceUnique([]string{"a", "b", "c", "a", "b", "c"}))

	// 发送一个简单的 HTTP GET 请求, 返回字符串 HTML
	html, _ := fun.HttpGet("https://www.github.com")
	fmt.Println(fun.String(html))
}
```

## 文档

### 时间相关

#### 函数列表

- **<big>`Timestamp(millis ...any) int64`</big>** 返回秒级时间戳

- **<big>`Date(layouts ...any) string`</big>** 返回格式化后的日期时间字符串

- **<big>`StrToTime(args ...any) int64`</big>** 自动解析日期时间字符串为秒时间戳, 近似 PHP strtotime()

```go
package main

import (
	"fmt"
	"time"

	"github.com/x-funs/go-fun"
)

func main() {
	// 返回秒级时间戳
	fmt.Println(fun.Timestamp())
	// 1673225645

	// 返回毫秒级时间戳
	fmt.Println(fun.Timestamp(true))
	// 1673225645077

	// 无参, 格式化当前时间(默认格式2006-01-02 15:04:05 )
	fmt.Println(fun.Date())
	// 2006-01-02 15:04:05

	// 格式化指定时间(默认格式2006-01-02 15:04:05 )
	fmt.Println(fun.Date(1650732457))
	// 2022-04-24 00:47:37

	// 格式化指定时间
	fmt.Println(fun.Date(time.RFC3339, 1650732457))
	// 2022-04-24T00:47:37+08:00

	// 无参, 等同于 Timestamp()
	fmt.Println(fun.StrToTime())
	// 1673226381

	// 一天前的时间戳
	fmt.Println(fun.StrToTime("-1 day"))
	// 1673139981

	// 某一时间戳一天后的时间戳
	fmt.Println(fun.StrToTime("+1 day", 1673225645))
	// 1673312045
}
```

### 辅助相关

#### 函数列表

- **<big>`If(condition bool, trueVal, falseVal T) T`</big>** 三元运算符函数

- **<big>`Empty(value any) bool`</big>** 判断是否为空, 支持字符串、数值、数组、Slice、Map

- **<big>`EmptyAll(values ...any) bool`</big>** 判断是否全部为空

- **<big>`EmptyAny(values ...any) bool`</big>** 判断是否任意一个为空

- **<big>`MemoryBytes() map[string]int64`</big>** 返回当前主要的内存指标信息

- **<big>`Memory(format string) map[string]int64`</big>** 指定格式返回当前主要的内存指标信息

- **<big>`Bytes(s string) []byte`</big>** 更高效的字符串转字节数组, 参考来自 `Gin`

- **<big>`String(b []byte) string`</big>** 更高效的字节数组转字符串, 参考来自 `Gin`

- **<big>`Command(bin string, argv []string, baseDir string) ([]byte, error)`</big>** 执行系统命令

```go
package main

import (
	"fmt"

	"github.com/x-funs/go-fun"
)

func main() {
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
}
```

### 哈希相关

#### 函数列表

- **<big>`Md5(str string) string`</big>** 返回字符串 Md5 值

- **<big>`Md5Bit16(str string) string`</big>** 返回 16位 字符串 Md5 值

- **<big>`Sha1(str string) string`</big>** 返回字符串 Sha1 值

- **<big>`Sha256(str string) string`</big>** 返回字符串 Sha256 值

- **<big>`Sha384(str string) string`</big>** 返回字符串 Sha384 值

- **<big>`Sha512(str string) string`</big>** 返回字符串 Sha512 值

- **<big>`Base64Encode(str string) string`</big>** 返回字符串 Base64 值

- **<big>`Base64Decode(str string) string`</big>** 返回 Base64 值对应的字符串

- **<big>`Base64UrlEncode(str string) string`</big>** 返回字符串 Url Safe Base64 值

- **<big>`Base64UrlDecode(str string) string`</big>** 返回 Url Safe Base64 值对应的字符串

### 判断相关

#### 函数列表

- **<big>`IsNumber(str string) bool`</big>** 判断字符串是否全部为数字

- **<big>`IsUtf8(p []byte) bool`</big>** 判断是否为 UTF-8 编码

- **<big>`IsASCIILetter(str string) bool`</big>** 判断字符串是否全部为ASCII的字母

- **<big>`IsLetter(str string) bool`</big>** 判断字符串是否全部为字母

- **<big>`IsASCII(s string) bool`</big>** 判断字符串是否全部 ASCII

- **<big>`IsEmail(str string) bool`</big>** 验证 Email 是否合法

- **<big>`IsExist(path string) bool`</big>**  文件或目录是否存在

- **<big>`IsDir(path string) bool`</big>** 是否是目录

### Map 相关

#### 函数列表

- **<big>`MapKeys[K comparable, V any](m map[K]V) []K`</big>** 返回 map 所有键的切片

- **<big>`MapValues[K comparable, V any](m map[K]V) []V`</big>** 返回 map 所有值的切片

- **<big>`MapMerge[K comparable, V any](maps ...map[K]V) map[K]V`</big>** 合并多个 map, 如果有相同的键, 则后者会覆盖前者

### 数学相关

#### 函数列表

- **<big>`Max(a, b int) int`</big>** 取 int 最大值

- **<big>`Min(a, b int) int`</big>** 取 int 最小值

- **<big>`MaxInt64(a, b int64) int64`</big>** 取 int64 最大值

- **<big>`MinInt64(a, b int64) int64`</big>** 取 int64 最小值

- **<big>`MaxN[T GenNumber](args ...T) T`</big>** 取 N 个数字的最大值

- **<big>`MinN[T GenNumber](args ...T) T`</big>** 取 N 个数字的最小值

### 随机相关

#### 函数列表

- **<big>`Random() int`</big>** 返回随机数 `[0, MaxInt)`

- **<big>`RandomInt(min, max int) int`</big>** 返回随机数 `[min, max)`

- **<big>`RandomInt64(min, max int64) int64`</big>** 返回随机数 `[min, max)`

- **<big>`RandomString(length int) string`</big>** 返回指定长度的随机字符串, 包含字母和数字

- **<big>`RandomLetter(length int) string`</big>** 返回指定长度的随机字符串, 仅包含字母

- **<big>`RandomNumber(length int) string`</big>** 返回指定长度的随机字符串, 仅包含数字

- **<big>`RandomPool(pool string, length int) string`</big>** 从提供的字符串池中返回指定长度的随机字符串

### 正则相关

#### 函数列表

- **<big>`Matches(str, pattern string) bool`</big>** 判断字符串是否匹配指定的正则表达式

### 相似度相关

#### 函数列表

- **<big>`Similarity(a, b string) float64`</big>** 计算两个原始字符串的相似度

- **<big>`SimilarityText(a, b string) float64`</big>** 计算两个字符串移除特殊符号后的相似度

- **<big>`LongestCommonSubString(x, y string) int`</big>** 计算两个字符串最大公共子串长度

### 切片相关方法

#### 函数列表

- **<big>`SliceSplit[T comparable](slice []T, size int) [][]T`</big>** 对数值和字符串切片按照指定长度进行分割

- **<big>`SliceUnion[T comparable](slices ...[]T) []T`</big>** 顺序合并且去重

- **<big>`SliceColumn[T, V any](slice []T, key any) []V`</big>** 返回所有行的某一列

- **<big>`IntsToStrings(slice []int) []string`</big>** int 切片转换为字符串切片

- **<big>`StringsToInts(slice []string) []int`</big>** 字符串切片转换为 int 切片

- **<big>`SliceContains[T comparable](slice []T, v T) bool`</big>** 判断整型和字符串是否在切片中

- **<big>`SliceUnique[T comparable](slice []T) []T`</big>** 对数值和字符串切片进行去重(会改变元素的顺序)

- **<big>`SliceIndex[T comparable](slice []T, v T) int`</big>** 对数值和字符串切片按照指定值进行查找

- **<big>`SliceLastIndex[T comparable](slice []T, v T) int`</big>** 对数值和字符串切片按照指定值进行查找, 返回最后一个匹配的索引

- **<big>`SliceRemove[T comparable](slice []T, v T) []T`</big>** 移除数值和字符串切片中的指定值

- **<big>`SliceRemoveBlank(slice []string) []string`</big>** 移除字符串切片中的空值

- **<big>`SliceTrim(slice []string) []string`</big>** 对字符串切片进行 Trim, 并自动忽略空值

- **<big>`SliceConcat[T any](slice []T, values ...[]T) []T`</big>** 合并多个切片, 非去重, 非原始切片

- **<big>`SliceEqual[T comparable](slice1, slice2 []T) bool`</big>** 切片是否相等: 长度相同且所有元素的顺序和值相等

- **<big>`SliceEvery[T any](slice []T, predicate func(index int, item T) bool) bool`</big>** 切片中的所有元素都满足函数，则返回 true

- **<big>`SliceNone[T any](slice []T, predicate func(index int, item T) bool) bool`</big>** 切片中的所有元素都不满足函数，则返回 true

- **<big>`SliceSome[T any](slice []T, predicate func(index int, item T) bool) bool`</big>** 切片中有一个元素满足函数，就返回true

- **<big>`SliceFilter[T any](slice []T, predicate func(index int, item T) bool) []T`</big>** 筛选出切片中满足函数的所有元素

- **<big>`SliceForEach[T any](slice []T, iteratee func(index int, item T))`</big>** 切片中所有元素都执行函数

- **<big>`SliceMap[T any, U any](slice []T, iteratee func(index int, item T) U) []U`</big>** 切片中所有元素都执行函数, 有返回值

- **<big>`SliceReduce[T any](slice []T, iteratee func(index int, result, item T) T, initial T) T`</big>** 处理所有切片中元素得到结果

- **<big>`SliceReplace[T comparable](slice []T, old T, new T, n int) []T`</big>** 返回切片的副本，前n个元素替换为新的

- **<big>`SliceReplaceAll[T comparable](slice []T, old T, new T) []T`</big>** 返回切片的副本，所有匹配到的元素都替换为新的

- **<big>`SliceUnionBy[T any, V comparable](predicate func(item T) V, slices ...[]T) []T`</big>** 顺序合并且去重, 支持自定义函数

- **<big>`SliceIntersection[T comparable](slices ...[]T) []T`</big>** 切片交集且去重(顺序不能保证)

- **<big>`SliceSortBy(slice any, field string, sortType ...string) error`</big>** 根据字段排序(field的大小写应该和字段保持一致)


```go
package main

import (
	"fmt"

	"github.com/x-funs/go-fun"
)

func main() {
	fmt.Println(fun.SliceSplit([]string{"a", "b", "c", "d", "e", "f", "g"}, 3))
	// [[a b c] [d e f] [g]]

	fmt.Println(fun.SliceUnion([]string{"123", "124"}, []string{"124", "125"}, []string{"123", "125"}))
	// [123 124 125]

	fmt.Println(
		fun.SliceColumn[map[string]string, string]([]map[string]string{
			{"name": "管理员", "code": "YF4133"},
			{"name": "用户", "code": "MM8541"},
			{"name": "测试", "code": "KH0002"},
			{"name": "演示", "code": "SJ9642"},
		}, "code"),
	)
	// [YF4133 MM8541 KH0002 SJ9642]
}
```

### 字符串相关

#### 函数列表

- **<big>`StrBefore(s, char string) string`</big>** 截取在字符首次出现时的位置之前的子字符串

- **<big>`StrBeforeLast(s, char string) string`</big>** 截取在字符最后出现时的位置之前的子字符串

- **<big>`StrAfter(s, char string) string`</big>** 截取在字符首次出现时的位置之后的子字符串

- **<big>`StrAfterLast(s, char string) string`</big>** 截取在字符最后出现时的位置之后的子字符串

- **<big>`Blank(str string) bool`</big>** 判断 Trim 后的字符串, 是否为空白

- **<big>`BlankAll(strs ...string) bool`</big>** 判断 Trim 后的字符串集, 是否全部为空白

- **<big>`BlankAny(strs ...string) bool`</big>** 判断 Trim 后的字符串集, 是否任意一个包含空白

- **<big>`HasPrefixCase(str, prefix string) bool`</big>** 判断字符串是否以指定前缀开头, 忽略大小写

- **<big>`HasSuffixCase(str, prefix string) bool`</big>** 判断字符串是否以指定后缀结尾, 忽略大小写

- **<big>`SplitTrim(str, sep string) []string`</big>** 分割字符串为字符串切片, 对分割后的值进行 Trim , 并自动忽略空值

- **<big>`SplitTrimToInts(str, sep string) []int`</big>** 分割字符串为 int 切片, 对分割后的值进行 Trim , 并自动忽略空值

- **<big>`Contains(str, substr string) bool`</big>** 判断字符串是否包含指定的子串

- **<big>`ContainsCase(str, substr string) bool`</big>** 判断字符串是否包含指定的子串, 不区分大小写

- **<big>`ContainsAny(str string, substr ...string) bool`</big>** 判断字符串是否包含任意一个指定的多个子串

- **<big>`SnakeToCamel(str string, bigCamel bool) string`</big>** 蛇形转驼峰

- **<big>`CamelToSnake(str string) string`</big>** 驼峰转蛇形

- **<big>`PadLeft(str string, padStr string, padLen int) string`</big>** 左侧填充字符串到指定长度

- **<big>`PadRight(str string, padStr string, padLen int) string`</big>** 右侧填充字符串到指定长度

- **<big>`PadBoth(str string, padStr string, padLen int) string`</big>** 两侧填充字符串到指定长度

- **<big>`Wrap(str string, wrapStr string) string`</big>** 使用字符串包围原字符串

- **<big>`Unwrap(str string, wrapStr string) string`</big>** 去除字符串包围, 非递归

- **<big>`Reverse(str string) string`</big>** 反转字符串

- **<big>`Remove(str, remove string) string`</big>** 移除字符串中指定的字符串

- **<big>`RemovePrefix(str, prefix string) string`</big>** 左侧移除字符串中指定的字符串

- **<big>`RemoveSuffix(str string, suffix string) string`</big>** 右侧移除字符串中指定的字符串

- **<big>`RemoveAny(str string, removes ...string) string`</big>** 移除字符串中指定的字符串集

- **<big>`RemoveSign(str string) string`</big>** 将字符串的所有数据依次写成一行, 去除无意义字符串(标点符号、符号)

- **<big>`RemoveLines(str string) string`</big>** 移除换行符, 换行符包括 \n \r\n

- **<big>`SubString(str string, pos, length int) string`</big>** 字符串截取

- **<big>`NormaliseSpace(str string) string`</big>** 规范化此字符串中的空白, 多个空格合并为一个空格, 所有空白字符例如换行符、制表符, 都转换为一个简单的空格

- **<big>`NormaliseLine(str string) string`</big>** 规范化此字符串中的换行, 多个换行合并为一个换行

- **<big>`Template(tpl string, data any) (string, error)`</big>** 模板渲染

```go
package main

import (
	"fmt"

	"github.com/x-funs/go-fun"
)

func main() {
	fmt.Println(fun.StrBefore("http://admin:123123@127.0.0.1:27017", ":"))
	// http

	fmt.Println(fun.StrAfter("https://github.com", "://"))
	// github.com

	fmt.Println(fun.StrBeforeLast("video.mp4.bak", "."))
	// video.mp4

	fmt.Println(fun.StrAfterLast("video.mp4.bak", "."))
	// bak
}
```

### 结构体相关

#### 函数列表

- **<big>`StructCopy(src, dst any)`</big>** 复制 struct 对象

### 转换相关

#### 函数列表

- **<big>`Ip2Long(ipStr string) uint32`</big>** 字符串 IP 转整型

- **<big>`Long2Ip(long uint32) string`</big>** 整型转字符串 IP

- **<big>`ToString(value any) string`</big>** 将任意一个类型转换为字符串

- **<big>`ToInt(value any) int`</big>** 数字或字符串转 int 类型

- **<big>`ToLong(value any) int64`</big>** ToInt64 别名, 数字或字符串转 int64

- **<big>`ToBool(str string) bool`</big>** 字符串转 bool 类型

- **<big>`ToUint(value any) uint`</big>** 数字或字符串转 uint

- **<big>`ToUint8(value any) uint8`</big>** 数字或字符串转 uint8

- **<big>`ToInt64(value any) int64`</big>** 数字或字符串转 int64

- **<big>`ToFloat32(value any) float32`</big>** 数字或字符串转 float32

- **<big>`ToFloat64(value any) float64`</big>** 数字或字符串转 float64

- **<big>`ToUtf8(origin []byte, encode string) ([]byte, error)`</big>** 指定字符集转 utf-8

- **<big>`Utf8To(utf8 []byte, encode string) ([]byte, error)`</big>** utf-8 转指定字符集

- **<big>`ToJson(object any) string`</big>** 将对象转换为 Json 字符串

- **<big>`ToJsonIndent(object any) string`</big>** 将对象转换为具有锁进的 Json 字符串

- **<big>`ToDuration(value any) time.Duration`</big>** 数字或字符串转 time.Duration，默认是纳秒，字符串支持 "ns,ms,us,s,m,h"

- **<big>`ToDurationMs(value any) time.Duration`</big>** 数字或字符串转 time.Duration，默认是毫秒，字符串支持 "ns,ms,us,s,m,h"

### 文件相关

#### 函数列表

- **<big>`Mkdir(dir string, perm os.FileMode) error`</big>** 创建一个目录，如果目录已存在则忽略

- **<big>`FileExists(path string) bool`</big>** 检测目录或者文件是否存在，返回 bool

- **<big>`WriteFile(name string, data []byte, flag int, perm os.FileMode, sync bool) error`</big>** WriteFile 写入文件

- **<big>`WriteFileAppend(name string, data []byte, perm os.FileMode, sync bool) error`</big>** 追加写入文件

### http 相关

> HttpXXResp 后缀的, 返回值是 *Response
> HttpXXDo 后缀的, 需要传参数 *Request

#### 函数列表

- **<big>`HttpGet(urlStr string, args ...any) ([]byte, error)`</big>** HttpGet 参数为请求地址 (HttpReq, 超时时间)

- **<big>`HttpPost(urlStr string, args ...any) ([]byte, error)`</big>** HttpPost 参数为请求地址 (body io.Reader, HttpReq, 超时时间)

- **<big>`HttpPostForm(urlStr string, args ...any) ([]byte, error)`</big>** HttpPostForm 参数为请求地址 (Form 数据 map[string]string, HttpReq, 超时时间)

- **<big>`HttpPostJson(urlStr string, args ...any) ([]byte, error)`</big>** HttpPostJson 参数为请求地址 (Json 数据 string, HttpReq, 超时时间)

- **<big>`UrlParse(rawURL string) (*url.URL, error)`</big>** 解析字符串 URL 到 URL 对象。在没有 scheme 时不会出错




