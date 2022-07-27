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