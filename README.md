# Go with Fun (Functions)

## Install

```shell
go get -u github.com/x-funs/go-fun
```

## Example

```go
import (
	"fmt"

	"github.com/x-funs/go-fun"
)

func main() {
	fmt.Println(fun.Md5("123456"))
}
```

## Functions

```
func Base64Decode(str string) string
    Base64Decode 返回 Base64 对应的字符串

func Base64Encode(str string) string
    Base64Encode 返回字符串 Base64 值

func Base64UrlDecode(str string) string
    Base64UrlDecode 返回 Url Safe Base64 对应的字符串

func Base64UrlEncode(str string) string
    Base64UrlEncode 返回字符串 Url Safe Base64 值

func Blank(str string) bool
    Blank 判断 Trim 后的字符串，是否为空白

func BlankAll(strs ...string) bool
    BlankAll 判断 Trim 后的字符串集，是否全部为空白

func BlankAny(strs ...string) bool
    BlankAny 判断 Trim 后的字符串集，是否任意一个包含空白

func CamelToUnder(str string) string
    CamelToUnder 大驼峰转下划线

func Contains(str, substr string) bool
    Contains 判断字符串是否包含子串

func ContainsAny(str string, substr ...string) bool
func ContainsCase(str, substr string) bool
    ContainsCase 判断字符串是否包含子串，不区分大小写

func Date(layouts ...any) string
    Date 返回格式化后的日期时间字符串，支持 Date()、Date(int)、Date(string)、Date(string, int)

func DateByDefault() string
    DateByDefault 返回格式化后的日期时间字符串

func DateByPattern(layout string) string
    DateByPattern 返回指定格式化后的日期时间字符串

func DateByPatternAndTime(layout string, timeStamp int64) string
    DateByPatternAndTime 返回指定时间戳格式化后的日期时间字符串

func Empty(value any) bool
    Empty 判断是否为空，支持字符串、数值、数组、切片、Map

func EmptyAll(values ...any) bool
    EmptyAll 判断是否全部为空

func EmptyAny(values ...any) bool
    EmptyAny 判断是否任意一个为空

func HttpGet(urlStr string, args ...any) (string, error)
    HttpGet 参数为请求地址（超时时间，请求头map[string]string），返回值为请求内容，错误信息

func HttpGetBody(urlStr string, timeout int, headers map[string]string) (string, error)
    HttpGetBody Http Get，参数为请求地址，超时时间，请求头 map[string]string，返回值为请求内容，错误信息

func HttpPost(urlStr string, args ...any) (string, error)
    HttpPost
    参数为请求地址（数据map[string]string，超时时间，请求头map[string]string），返回值为请求内容，错误信息

func HttpPostBody(urlStr string, posts map[string]string, timeout int, headers map[string]string) (string, error)
    HttpPostBody Http Post Form，参数为请求地址，Form 数据map[string]string，超时时间，请求头
    map[string]string，返回值为请求内容，错误信息

func HttpPostJsonBody(urlStr string, json string, timeout int, headers map[string]string) (string, error)
    HttpPostJsonBody Http Post Json，参数为请求地址，Json数据string，超时时间，请求头
    map[string]string，返回值为请求内容，错误信息

func Ip2Long(ipStr string) uint32
    Ip2Long 字符串 IP 转整型

func IsLetter(str string) bool
    IsLetter 判断字符串是否全部为字母

func IsNumber(str string) bool
    IsNumber 判断字符串是否全部为数字

func Long2Ip(long uint32) string
    Long2Ip 整型转字符串 IP

func Matches(str, pattern string) bool
    Matches 判断字符串是否匹配正则表达式

func Md5(str string) string
    Md5 返回字符串 Md5 值

func MemoryBytes() map[string]int64
    MemoryBytes 返回当前主要的内存指标信息

func PadBoth(str string, padStr string, padLen int) string
    PadBoth 两侧填充字符串到指定长度

func PadLeft(str string, padStr string, padLen int) string
    PadLeft 左侧填充字符串到指定长度

func PadRight(str string, padStr string, padLen int) string
    PadRight 右侧填充字符串到指定长度

func Random() int
    Random 返回随机数 [0, MaxInt)

func RandomInt(min, max int) int
    RandomInt 返回随机数 [min, max)

func RandomInt64(min, max int64) int64
    RandomInt64 返回随机数 [min, max)

func RandomLetter(length int) string
    RandomLetter 返回指定长度的随机字符串，包含大小写字母

func RandomNumber(length int) string
    RandomNumber 返回指定长度的随机字符串，包含数字

func RandomPool(pool string, length int) string
    RandomPool 从字符串池中返回指定长度的随机字符串

func RandomString(length int) string
    RandomString 返回指定长度的随机字符串，包含大小写字母和数字

func Remove(str, remove string) string
    Remove 移除指定字符串中给定字符串

func RemoveAny(str string, removes ...string) string
    RemoveAny 移除指定字符串中给定字符串集

func RemovePrefix(str, prefix string) string
    RemovePrefix 移除指定字符串左侧给定字符串

func RemoveSuffix(str string, suffix string) string
    RemoveSuffix 移除指定字符串右侧给定字符串

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

func SplitTrim(str, sep string) []string
    SplitTrim 分割字符串为切片，对分割后的值进行 Trim ，并自动忽略空值

func StrToTime(args ...any) int64
    StrToTime 日期时间字符串转时间戳，支持 StrToTime()、StrToTime(string)、StrToTime(string,
    int64)

func SubString(str string, pos, length int) string
    SubString 字符串截取

func Timestamp(millis ...any) int64
    Timestamp 返回当前时间的 Unix 时间戳，默认返回秒级，支持 Timestamp(true) 返回毫秒级

func ToInt(value any) int
    ToInt 数字和字符串转 Int

func ToInt64(value any) int64
    ToInt64 数字和字符串转 Int64

func ToJson(object any) string
    ToJson 将对象转换为json字符串

func ToLong(value any) int64
    ToLong 数字和字符串转 Int64

func ToString(value any) string
    ToString 将任意一个类型转换为字符串

func UnderToCamel(str string) string
    UnderToCamel 下划线转大驼峰

func UnixMilliTimestamp() int64
    UnixMilliTimestamp 返回当前时间的 Unix 毫秒级时间戳

func UnixTimestamp() int64
    UnixTimestamp 返回当前时间的 Unix 秒级时间戳
```