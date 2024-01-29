package fun

import (
	"errors"
	"net"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

// IsNumber 判断字符串是否全部为数字
func IsNumber(str string) bool {
	if len(str) == 0 {
		return false
	}

	for _, r := range str {
		if !unicode.IsNumber(r) {
			return false
		}
	}

	return true
}

// IsUtf8 判断是否为 UTF-8 编码
func IsUtf8(p []byte) bool {
	return utf8.Valid(p)
}

// IsASCIILetter 判断字符串是否全部为ASCII的字母
func IsASCIILetter(str string) bool {
	if len(str) == 0 {
		return false
	}

	runeList := []rune(str)
	for _, r := range runeList {
		if !((65 <= r && r <= 90) || (97 <= r && r <= 122)) {
			return false
		}
	}

	return true
}

// IsLetter 判断字符串是否全部为字母
func IsLetter(str string) bool {
	if len(str) == 0 {
		return false
	}

	for _, r := range str {
		if !unicode.IsLetter(r) {
			return false
		}
	}

	return true
}

// IsASCII 判断字符串是否全部 ASCII
func IsASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

// IsEmail 验证 Email 是否合法
func IsEmail(str string) bool {
	if !Blank(str) {
		return RegexEmailPattern.MatchString(str)
	}

	return false
}

// IsExist 文件或目录是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return false
}

// IsDir 是否是目录
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}

	return s.IsDir()
}

// IsIp 是否是有效的 IP
func IsIp(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	return ip != nil
}

// IsIpV4 是否是有效的 IpV4
func IsIpV4(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return false
	}
	return strings.Contains(ipStr, ".")
}

// IsIpV6 是否是有效的 IpV6
func IsIpV6(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return false
	}
	return strings.Contains(ipStr, ":")
}
