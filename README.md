```
                      ____          
   ____ _____        / __/_  ______ 
  / __ `/ __ \______/ /_/ / / / __ \
 / /_/ / /_/ /_____/ __/ /_/ / / / /
 \__, /\____/     /_/  \__,_/_/ /_/ 
/____/                              

```

Go with Fun (Functions) is a small and useful Golang util function library. It Includes such as Empty、Blank、Strtotime、Similarity、HttpGet etc.

English | [简体中文](./README_zh.md)

## Installation

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
	// Whether any type value is empty
	fmt.Println(fun.Empty(""))
	
	// Whether string value is blank
	fmt.Println(fun.Blank("  "))
	
	// Return MD5 string from a string
	fmt.Println(fun.Md5("go-fun"))
	
	// Auto parse datetime layout to int64 timestamp
	fmt.Println(fun.StrToTime("2015-04-06 16:03:03"))
	fmt.Println(fun.StrToTime("2015/04/06 16:03:03"))
	fmt.Println(fun.StrToTime("2022-01-24T14:19:00Z"))
	fmt.Println(fun.StrToTime("2022-01-24T14:19:01+07:00"))
	
	// Slice deduplication filter
	fmt.Println(fun.SliceUnique([]string{"a", "b", "c", "a", "b", "c"}))

	// Send a Simple HTTP GET request, Return HTML string
	html, _ := fun.HttpGet("https://www.github.com")
	fmt.Println(fun.String(html))
}
```

## Documentation

### DateTime

#### Function List

- **<big>`Timestamp(millis ...any) int64`</big>** Return the current unix timestamp.

- **<big>`Date(layouts ...any) string`</big>** Return the formatted datetime string.

- **<big>`StrToTime(args ...any) int64`</big>** Auto parse datetime layout to int64 timestamp, just like PHP strtotime().

```go
package main

import (
	"fmt"

	"github.com/x-funs/go-fun"
)

func main() {
	// second timestamp
	fmt.Println(fun.Timestamp())
	// 1673225645
	
	// millisecond timestamp
	fmt.Println(fun.Timestamp(true))
	// 1673225645077

	// no arguments, format datetime now (default by '2006-01-02 15:04:05')
	fmt.Println(fun.Date())
	// 2006-01-02 15:04:05

	// format datetime by timestamp (default by '2006-01-02 15:04:05')
	fmt.Println(fun.Date(1650732457))
	// 2022-04-24 00:47:37

	// use layout format datetime by timestamp
	fmt.Println(fun.Date(time.RFC3339, 1650732457))
	// 2022-04-24T00:47:37+08:00
	
	// no arguments, same as Timestamp()
	fmt.Println(fun.StrToTime())
	// 1673226381

	// one day before now timestamp
	fmt.Println(fun.StrToTime("-1 day"))
	// 1673139981 (yesterday)

	fmt.Println(fun.StrToTime("+1 day", 1673225645))
	// 1673312045 (one day after a certain timestamp)
}
```

### Helpers

#### Function List

- **<big>`If(condition bool, trueVal, falseVal T) T`</big>** Verify condition is true, return trueVal or falseVal

- **<big>`Empty(value any) bool`</big>** Verify whether value it is empty, support string, integer, array, slice, map 验证

- **<big>`EmptyAll(values ...any) bool`</big>** Verify whether values all are empty

- **<big>`EmptyAny(values ...any) bool`</big>** Verify whether values any is empty

- **<big>`MemoryBytes() map[string]int64`</big>** Return the current main memory metrics.

- **<big>`Memory(format string) map[string]int64`</big>** Specified format return the current main memory metric.

- **<big>`Bytes(s string) []byte`</big>** Efficient string to byte array, reference from `Gin`

- **<big>`String(b []byte) string`</big>** Efficient byte array to string, reference from `Gin`

- **<big>`Command(bin string, argv []string, baseDir string) ([]byte, error)`</big>** Execute system commands

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

### Hash

#### Function List

- **<big>`Md5(str string) string`</big>** Return the Md5 string

- **<big>`Md5Bit16(str string) string`</big>** Return the 16-bit Md5 string

- **<big>`Sha1(str string) string`</big>** Return the Sha1 string

- **<big>`Sha256(str string) string`</big>** Return the Sha256 string

- **<big>`Sha384(str string) string`</big>** Return the Sha384 string

- **<big>`Sha512(str string) string`</big>** Return the Sha512 string

- **<big>`Base64Encode(str string) string`</big>** Return the Base64 string

- **<big>`Base64Decode(str string) string`</big>** Return the Base64 decode string

- **<big>`Base64UrlEncode(str string) string`</big>** Return the Url Safe Base64 string

- **<big>`Base64UrlDecode(str string) string`</big>** Return the Url Safe Base64 decode string

### Judgment

#### Function List

- **<big>`IsNumber(str string) bool`</big>** Determine whether all strings are numbers

- **<big>`IsUtf8(p []byte) bool`</big>** Determine whether it is a UTF-8 code

- **<big>`IsASCIILetter(str string) bool`</big>** Determine whether all strings are ASCII letters

- **<big>`IsLetter(str string) bool`</big>** Determine whether all strings are letters

- **<big>`IsASCII(s string) bool`</big>** Determine whether the string is all ASCII

- **<big>`IsEmail(str string) bool`</big>** Verify Email

- **<big>`IsExist(path string) bool`</big>**  Does the file or directory exist

- **<big>`IsDir(path string) bool`</big>** Is it a directory

### Map

#### Function List

- **<big>`MapKeys[K comparable, V any](m map[K]V) []K`</big>** Return slices of all keys of map

- **<big>`MapValues[K comparable, V any](m map[K]V) []V`</big>** Return a slice of all values of map

- **<big>`MapMerge[K comparable, V any](maps ...map[K]V) map[K]V`</big>** Merge multiple maps, if there are the same keys, the latter will overwrite the former

### Math

#### Function List

- **<big>`Max(a, b int) int`</big>** Take int maximum

- **<big>`Min(a, b int) int`</big>** Take int minimum

- **<big>`MaxInt64(a, b int64) int64`</big>** Take int64 maximum

- **<big>`MinInt64(a, b int64) int64`</big>** Take int64 minimum

- **<big>`MaxN[T GenNumber](args ...T) T`</big>** Take the maximum value of n numbers

- **<big>`MinN[T GenNumber](args ...T) T`</big>** Take the minimum value of n numbers

### Random

#### Function List

- **<big>`Random() int`</big>** Return a random number `[0, MaxInt)`

- **<big>`RandomInt(min, max int) int`</big>** Return a random number `[min, max)`

- **<big>`RandomInt64(min, max int64) int64`</big>** Return a random number `[min, max)`

- **<big>`RandomString(length int) string`</big>** Return a random string of the specified length, including letters and numbers.

- **<big>`RandomLetter(length int) string`</big>** Return a random string of the specified length, containing only letters.

- **<big>`RandomNumber(length int) string`</big>** Return a random string of the specified length, containing only numbers.

- **<big>`RandomPool(pool string, length int) string`</big>** Return a random string of the specified length from the supplied string pool

### Regex

#### Function List

- **<big>`Matches(str, pattern string) bool`</big>** Determines whether the string matches the specified regular expression.

### Similarity

#### Function List

- **<big>`Similarity(a, b string) float64`</big>** Calculates the similarity of two original strings

- **<big>`SimilarityText(a, b string) float64`</big>** Calculate the similarity of two strings after removing special symbols

- **<big>`LongestCommonSubString(x, y string) int`</big>** Calculates the maximum common substring length of two strings

### Slice

#### Function List

- **<big>`SliceSplit[T comparable](slice []T, size int) [][]T`</big>** Divide numeric and string slices according to the specified length

- **<big>`SliceUnion[T comparable](slices ...[]T) []T`</big>** Sequential merge and deweight

- **<big>`SliceColumn[T, V any](slice []T, key any) []V`</big>** Return a column of all rows

- **<big>`IntsToStrings(slice []int) []string`</big>** Int slice to string slice

- **<big>`StringsToInts(slice []string) []int`</big>** String slice to int slice

- **<big>`SliceContains[T comparable](slice []T, v T) bool`</big>** Determine whether integer and string are in slice

- **<big>`SliceUnique[T comparable](slice []T) []T`</big>** Devaluation of numeric and string slices (changes the order of elements)

- **<big>`SliceIndex[T comparable](slice []T, v T) int`</big>** Find numeric and string slices according to the specified value

- **<big>`SliceLastIndex[T comparable](slice []T, v T) int`</big>** The value and string slices are searched according to the specified value, and the last matching index is returned.

- **<big>`SliceRemove[T comparable](slice []T, v T) []T`</big>** Removes the specified value from numeric and string slices

- **<big>`SliceRemoveBlank(slice []string) []string`</big>** Remove null values from string slices

- **<big>`SliceTrim(slice []string) []string`</big>** Trim string slices and automatically ignore null values

- **<big>`SliceConcat[T any](slice []T, values ...[]T) []T`</big>** Merge multiple slices, non-degravimetric, non-original slices

- **<big>`SliceEqual[T comparable](slice1, slice2 []T) bool`</big>** Are slices equal: the same length and the order and value of all elements are equal

- **<big>`SliceEvery[T any](slice []T, predicate func(index int, item T) bool) bool`</big>** All elements in the slice satisfy the function, return true

- **<big>`SliceNone[T any](slice []T, predicate func(index int, item T) bool) bool`</big>** Return true if all elements in the slice do not satisfy the function.

- **<big>`SliceSome[T any](slice []T, predicate func(index int, item T) bool) bool`</big>** If one element in the slice satisfies the function, it Return true.

- **<big>`SliceFilter[T any](slice []T, predicate func(index int, item T) bool) []T`</big>** Filter out all elements in the slice that satisfy the function

- **<big>`SliceForEach[T any](slice []T, iteratee func(index int, item T))`</big>** All elements in the slice execute functions

- **<big>`SliceMap[T any, U any](slice []T, iteratee func(index int, item T) U) []U`</big>** All elements in the slice execute functions and have return values.

- **<big>`SliceReduce[T any](slice []T, iteratee func(index int, result, item T) T, initial T) T`</big>** Process all elements in slices to get results

- **<big>`SliceReplace[T comparable](slice []T, old T, new T, n int) []T`</big>** Return a copy of the slice, with the first n elements replaced with the new

- **<big>`SliceReplaceAll[T comparable](slice []T, old T, new T) []T`</big>** Return a copy of the slice, and all matching elements are replaced with new ones.

- **<big>`SliceUnionBy[T any, V comparable](predicate func(item T) V, slices ...[]T) []T`</big>** Order merge and de-heavy, support custom functions

- **<big>`SliceIntersection[T comparable](slices ...[]T) []T`</big>** Slices intersect and deweight (order cannot be guaranteed)

- **<big>`SliceSortBy(slice any, field string, sortType ...string) error`</big>** Sort by field (field case should be consistent with field)

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
			{"name": "admin", "code": "YF4133"},
			{"name": "user", "code": "MM8541"},
			{"name": "test", "code": "KH0002"},
			{"name": "demo", "code": "SJ9642"},
		}, "code"),
	)
	// [YF4133 MM8541 KH0002 SJ9642]
}
```

### String

#### Function List

- **<big>`StrBefore(s, char string) string`</big>** Intercept the substring before the position of the character when it first appears.

- **<big>`StrBeforeLast(s, char string) string`</big>** Intercept the substring before the last appearance of the character

- **<big>`StrAfter(s, char string) string`</big>** Interception of substrings after the position of the character when it first appears

- **<big>`StrAfterLast(s, char string) string`</big>** Interception of substrings after the last appearance of the character

- **<big>`Blank(str string) bool`</big>** Determine whether the string after Trim is blank.

- **<big>`BlankAll(strs ...string) bool`</big>** Determine whether the string set after Trim is all blank.

- **<big>`BlankAny(strs ...string) bool`</big>** Determine whether any string set after Trim contains a blank.

- **<big>`HasPrefixCase(str, prefix string) bool`</big>** Determines whether the string starts with the specified prefix, ignoring case

- **<big>`HasSuffixCase(str, prefix string) bool`</big>** Determine whether the string ends with the specified suffix, ignoring the case.

- **<big>`SplitTrim(str, sep string) []string`</big>** The split string is a string slice, the split value is Trim, and the null value is automatically ignored.

- **<big>`SplitTrimToInts(str, sep string) []int`</big>** The split string is an int slice, the split value is Trim, and the null value is automatically ignored.

- **<big>`Contains(str, substr string) bool`</big>** Determines whether the string contains the specified substring

- **<big>`ContainsCase(str, substr string) bool`</big>** Determine whether the string contains the specified substring, case-insensitive

- **<big>`ContainsAny(str string, substr ...string) bool`</big>** Determine whether the string contains any of the specified substrings.

- **<big>`SnakeToCamel(str string, bigCamel bool) string`</big>** Serpentine hump

- **<big>`CamelToSnake(str string) string`</big>** Hump turns to snake

- **<big>`PadLeft(str string, padStr string, padLen int) string`</big>** Fill the string on the left to the specified length

- **<big>`PadRight(str string, padStr string, padLen int) string`</big>** The right side fills the string to the specified length.

- **<big>`PadBoth(str string, padStr string, padLen int) string`</big>** Both sides fill the string to the specified length

- **<big>`Wrap(str string, wrapStr string) string`</big>** Enclosed the original string with a string

- **<big>`Unwrap(str string, wrapStr string) string`</big>** Remove string bounding, non-recursive

- **<big>`Reverse(str string) string`</big>** Reverse string

- **<big>`Remove(str, remove string) string`</big>** Removes the specified string in the string

- **<big>`RemovePrefix(str, prefix string) string`</big>** Removes the string specified in the string on the left

- **<big>`RemoveSuffix(str string, suffix string) string`</big>** The right side removes the specified string in the string.

- **<big>`RemoveAny(str string, removes ...string) string`</big>** Removes the string set specified in the string

- **<big>`RemoveSign(str string) string`</big>** Write all the data of the string into one line in turn, and remove meaningless strings (punctuation marks, symbols)

- **<big>`RemoveLines(str string) string`</big>** Remove line breaks, which include \n \r\n.

- **<big>`SubString(str string, pos, length int) string`</big>** String interception

- **<big>`NormaliseSpace(str string) string`</big>** Normalized the white space in this string, multiple spaces are merged into one space, and all white space characters such as line breaks and tabs are converted to a simple space.

- **<big>`NormaliseLine(str string) string`</big>** Standardize line breaks in this string, and merge multiple line breaks into one line break.

- **<big>`Template(tpl string, data any) (string, error)`</big>** Template rendering

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

### Struct

#### Function List

- **<big>`StructCopy(src, dst any)`</big>** Copy struct object

### To

#### Function List

- **<big>`Ip2Long(ipStr string) uint32`</big>** String IP to integer

- **<big>`Long2Ip(long uint32) string`</big>** Integer to string IP

- **<big>`ToString(value any) string`</big>** Converts any type to a string

- **<big>`ToInt(value any) int`</big>** Number or string to int type

- **<big>`ToLong(value any) int64`</big>** ToInt64 alias, number or string to int64

- **<big>`ToBool(str string) bool`</big>** string to bool type

- **<big>`ToUint(value any) uint`</big>** Number or string to uint

- **<big>`ToUint8(value any) uint8`</big>** Number or string to uint8

- **<big>`ToInt64(value any) int64`</big>** Number or string to int64

- **<big>`ToFloat32(value any) float32`</big>** Number or string to float32

- **<big>`ToFloat64(value any) float64`</big>** Number or string to float64

- **<big>`ToUtf8(origin []byte, encode string) ([]byte, error)`</big>** Specify character set conversion utf-8

- **<big>`Utf8To(utf8 []byte, encode string) ([]byte, error)`</big>** utf-8 to specify character set

- **<big>`ToJson(object any) string`</big>** Converts an object to a Json string

- **<big>`ToJsonIndent(object any) string`</big>** Converts an object to a Indent Json string

- **<big>`ToDuration(value any) time.Duration`</big>** Converts number or string to time.Duration, default is Nanosecond, string support "ns,ms,us,s,m,h"

- **<big>`ToDurationMs(value any) time.Duration`</big>** Converts number or string to time.Duration, default is Millisecond, string support "ns,ms,us,s,m,h"


### File

#### Function List

- **<big>`Mkdir(dir string, perm os.FileMode) error`</big>** Create a directory, ignoring if the directory already exists

- **<big>`FileExists(path string) bool`</big>** Check whether the directory or file exists, return bool

- **<big>`WriteFile(name string, data []byte, flag int, perm os.FileMode, sync bool) error`</big>** write file shortcut

- **<big>`WriteFileAppend(name string, data []byte, perm os.FileMode, sync bool) error`</big>** write file shortcut with append mode

### Http

> HttpXXResp the suffix, the return value is *Response
>
> HttpXXDo the suffix, Need to pass parameters *Request

#### Function List

- **<big>`HttpGet(urlStr string, args ...any) ([]byte, error)`</big>** The HttpGet parameter is the request address (HttpReq, timeout)

- **<big>`HttpPost(urlStr string, args ...any) ([]byte, error)`</big>** The HttpPost parameter is the request address (body io.Reader, HttpReq, timeout)

- **<big>`HttpPostForm(urlStr string, args ...any) ([]byte, error)`</big>** The HttpPostForm parameter is the request address (FormData map[string]string, HttpReq, timeout)

- **<big>`HttpPostJson(urlStr string, args ...any) ([]byte, error)`</big>** The HttpPostJson parameter is the request address (JsonData string, HttpReq, timeout)

- **<big>`UrlParse(rawURL string) (*url.URL, error)`</big>** Parses the string URL to the URL object. There will be no mistakes without scheme.

- **<big>`UserAgentRandom() string`</big>** generates a random DESKTOP browser user-agent on every requests .

- **<big>`UserAgentRandomMobile() string`</big>** generates a random MOBILE browser user-agent on every requests.
