# Go with fun (functions)

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
	// 字符串 MD5
	fmt.Println(fun.Md5("go-fun"))

	// 切片去重
	fmt.Println(fun.UniqueSlice([]string{"a", "b", "c", "a", "b", "c"}))

	// Http Get 请求
	html, _ := fun.HttpGet("https://www.163.com")
	fmt.Println(html)
}
```