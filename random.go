package fun

import (
	"math"
	"math/rand"
	"time"
)

var (
	randomNew = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// Random 返回随机数 `[0, MaxInt)`
func Random() int {
	return randomNew.Intn(math.MaxInt)
}

// RandomInt 返回随机数 `[min, max)`
func RandomInt(min, max int) int {
	if min > max {
		min, max = max, min
	}

	return randomNew.Intn(max-min) + min
}

// RandomInt64 返回随机数 `[min, max)`
func RandomInt64(min, max int64) int64 {
	if min > max {
		min, max = max, min
	}

	return randomNew.Int63n(max-min) + min
}

// RandomString 返回指定长度的随机字符串, 包含字母和数字
func RandomString(length int) string {
	return RandomPool(StringLetterAndNumber, length)
}

// RandomLetter 返回指定长度的随机字符串, 仅包含字母
func RandomLetter(length int) string {
	return RandomPool(StringLetter, length)
}

// RandomNumber 返回指定长度的随机字符串, 仅包含数字
func RandomNumber(length int) string {
	return RandomPool(StringNumber, length)
}

// RandomPool 从提供的字符串池中返回指定长度的随机字符串
func RandomPool(pool string, length int) string {
	if length <= 0 {
		return ""
	}
	var chars = Bytes(pool)
	var result []byte
	for i := 0; i < length; i++ {
		c := chars[RandomInt(0, len(chars))]
		result = append(result, c)
	}
	return String(result)
}
