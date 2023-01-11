# Go With Fun (Functions)

Go with Fun (Functions) is a small and useful Golang util function library.

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
	fmt.Println(fun.StrToTime("2015年04月06日 16时03分03秒"))
	
	// Slice deduplication fileter
	fmt.Println(fun.SliceUnique([]string{"a", "b", "c", "a", "b", "c"}))

	// Send a Simple HTTP GET request, Return HTML string
	html, _ := fun.HttpGet("https://www.github.com")
	fmt.Println(fun.String(html))
}
```

## Documentation

### DateTime

#### Function List

- **<big>`Timestamp(millis ...any) int64`</big>** Return the current unix timestamp

- **<big>`Date(layouts ...any) string`</big>** Return the formatted datetime string

- **<big>`StrToTime(args ...any) int64`</big>** Auto parse datetime layout to int64 timestamp, just like PHP strtotime()

```go
package main

import (
	"fmt"

	"github.com/x-funs/go-fun"
)

func main() {
	fmt.Println(fun.Timestamp())
	// 1673225645
	
	fmt.Println(fun.Timestamp(true))
	// 1673225645077
	
	fmt.Println(fun.StrToTime())
	// 1673226381

	fmt.Println(fun.StrToTime("-1 day"))
	// 1673139981 (yesterday)

	fmt.Println(fun.StrToTime("+1 day", 1673225645))
	// 1673312045 (one day after a certain timestamp)
}
```