# Go With Fun (Functions)

Go with Fun (Functions) is a small and useful Golang util function library.

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

#### **<big>`Timestamp(millis ...any) int64`</big>** Return the current unix timestamp

```
fmt.Println(fun.Timestamp())
// 1673225645

// 返回毫秒级时间戳
fmt.Println(fun.Timestamp(true))
// 1673225645077
```

#### **<big>`Date(layouts ...any) string`</big>** Return the formatted datetime string

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

