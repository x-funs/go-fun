# Go with Fun (Functions)

## Install

```
go get -u github.com/x-funs/go-fun
```

## Example

```
import (
	"fmt"

	"github.com/x-funs/go-fun"
)

func main() {
	fmt.Println(fun.Md5("123456"))
}
```

## APIs

* Timestamp
* UnixTimestamp
* UnixMilliTimestamp
* MemoryBytes
* Date
* DateByDefault
* DateByPattern
* DateByPatternAndTime
* ToString
* ToInt
* ToLong
* ToInt64
* Md5
* Sha1
* Sha256
* Sha384
* Sha512
* Base64Encode
* Base64Decode
* Base64UrlEncode
* Base64UrlDecode
* BlankAll
* BlankAny
* Blank
* EmptyAll
* EmptyAny
* Empty
* Ip2Long
* Long2Ip
* StrToTime
* SplitTrim
* IsNumber
* IsLetter
* Contains
* ContainsCase
* ContainsAny
* Matches
* UnderToCamel
* CamelToUnder
* PadLeft
* PadRight
* PadBoth
* buildPadStr
* ToJson
* Reverse
* Random
* RandomInt
* RandomInt64
* RandomString
* RandomLetter
* RandomNumber
* RandomPool
* Remove
* RemovePrefix
* RemoveSuffix
* RemoveAny
* SubString
* HttpGet
* HttpPost