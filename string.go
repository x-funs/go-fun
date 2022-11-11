package fun

import (
	"bytes"
	"math"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"
)

// BlankAll 判断 Trim 后的字符串集, 是否全部为空白
func BlankAll(strs ...string) bool {
	if len(strs) == 0 {
		return true
	}

	for _, v := range strs {
		if !Blank(v) {
			return false
		}
	}

	return true
}

// BlankAny 判断 Trim 后的字符串集, 是否任意一个包含空白
func BlankAny(strs ...string) bool {
	if len(strs) == 0 {
		return true
	}

	for _, v := range strs {
		if Blank(v) {
			return true
		}
	}

	return false
}

// Blank 判断 Trim 后的字符串, 是否为空白
func Blank(str string) bool {
	t := strings.TrimSpace(str)

	if t == "" {
		return true
	}

	return false
}

// HasPrefixCase 判断字符串是否以指定前缀开头, 忽略大小写
func HasPrefixCase(str, prefix string) bool {
	return strings.HasPrefix(strings.ToLower(str), strings.ToLower(prefix))
}

// HasSuffixCase 判断字符串是否以指定后缀结尾, 忽略大小写
func HasSuffixCase(str, prefix string) bool {
	return strings.HasSuffix(strings.ToLower(str), strings.ToLower(prefix))
}

// SplitTrim 分割字符串为字符串切片, 对分割后的值进行 Trim , 并自动忽略空值
func SplitTrim(str, sep string) []string {
	if len(str) == 0 || len(sep) == 0 {
		return []string{}
	}

	// 如果没找到 sep, strings.Split 返回包含 str 长度 1 的切片
	ss := strings.Split(str, sep)
	if len(ss) <= 1 {
		return []string{str}
	}

	slices := make([]string, 0, len(ss))
	for i := range ss {
		s := strings.TrimSpace(ss[i])
		if len(s) > 0 {
			slices = append(slices, s)
		}
	}

	return slices
}

// SplitTrimToInts 分割字符串为 int 切片, 对分割后的值进行 Trim , 并自动忽略空值
func SplitTrimToInts(str, sep string) []int {
	if len(str) == 0 || len(sep) == 0 {
		return []int{}
	}

	// 如果没找到 sep, strings.Split 返回包含 int(str) 长度 1 的切片
	ss := strings.Split(str, sep)
	if len(ss) <= 1 {
		s := ToInt(str)
		return []int{s}
	}

	slices := make([]int, 0, len(ss))
	for i := range ss {
		s := strings.TrimSpace(ss[i])
		if len(s) > 0 {
			if n, err := strconv.Atoi(s); err == nil {
				slices = append(slices, n)
			}
		}
	}

	return slices
}

// Contains 判断字符串是否包含指定的子串
func Contains(str, substr string) bool {
	return strings.Contains(str, substr)
}

// ContainsCase 判断字符串是否包含指定的子串, 不区分大小写
func ContainsCase(str, substr string) bool {
	return Contains(strings.ToLower(str), strings.ToLower(substr))
}

// ContainsAny 判断字符串是否包含任意一个指定的多个子串
func ContainsAny(str string, substr ...string) bool {
	if len(str) == 0 || len(substr) == 0 {
		return false
	}

	for _, s := range substr {
		if Contains(str, s) {
			return true
		}
	}

	return false
}

// SnakeToCamel 蛇形转驼峰
func SnakeToCamel(str string, bigCamel bool) string {
	if len(str) == 0 {
		return ""
	}

	if !Contains(str, UNDERSCORE) {
		return str
	}

	str = strings.ToLower(str)

	var sb strings.Builder
	sb.Grow(len(str))

	underscore := false
	for i, r := range str {
		if i == 0 {
			if bigCamel {
				sb.WriteRune(unicode.ToUpper(r))
			} else {
				sb.WriteRune(unicode.ToLower(r))
			}
		} else if r == '_' {
			if i < len(str) {
				underscore = true
			}
		} else if underscore {
			sb.WriteRune(unicode.ToUpper(r))
			underscore = false
		} else {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}

// CamelToSnake 驼峰转蛇形
func CamelToSnake(str string) string {
	if len(str) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.Grow(len(str))

	for i, r := range str {
		if i == 0 {
			sb.WriteRune(unicode.ToLower(r))
		} else if unicode.IsUpper(r) {
			sb.WriteRune('_')
			sb.WriteRune(unicode.ToLower(r))
		} else {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}

// PadLeft 左侧填充字符串到指定长度
func PadLeft(str string, padStr string, padLen int) string {
	if len(str) >= padLen || padStr == "" {
		return str
	}
	return buildPadStr(str, padStr, padLen, true, false)
}

// PadRight 右侧填充字符串到指定长度
func PadRight(str string, padStr string, padLen int) string {
	if len(str) >= padLen || padStr == "" {
		return str
	}
	return buildPadStr(str, padStr, padLen, false, true)
}

// PadBoth 两侧填充字符串到指定长度
func PadBoth(str string, padStr string, padLen int) string {
	if len(str) >= padLen || padStr == "" {
		return str
	}
	return buildPadStr(str, padStr, padLen, true, true)
}

// Wrap 使用字符串包围原字符串
func Wrap(str string, wrapStr string) string {
	if len(str) == 0 || wrapStr == "" {
		return str
	}
	var sb strings.Builder
	sb.WriteString(wrapStr)
	sb.WriteString(str)
	sb.WriteString(wrapStr)

	return sb.String()
}

// Unwrap 去除字符串包围, 非递归
func Unwrap(str string, wrapStr string) string {
	if str == "" || wrapStr == "" {
		return str
	}

	firstIndex := strings.Index(str, wrapStr)
	lastIndex := strings.LastIndex(str, wrapStr)

	if firstIndex == 0 && lastIndex > 0 && lastIndex <= len(str)-1 {
		if len(wrapStr) <= lastIndex {
			str = str[len(wrapStr):lastIndex]
		}
	}

	return str
}

// buildPadStr
func buildPadStr(str string, padStr string, padLen int, padLeft bool, padRight bool) string {
	if padLen < utf8.RuneCountInString(str) {
		return str
	}

	padLen -= utf8.RuneCountInString(str)

	targetLen := padLen

	targetLenLeft := targetLen
	targetLenRight := targetLen
	if padLeft && padRight {
		targetLenLeft = padLen / 2
		targetLenRight = padLen - targetLenLeft
	}

	strToRepeatLen := utf8.RuneCountInString(padStr)

	repeatTimes := int(math.Ceil(float64(targetLen) / float64(strToRepeatLen)))
	repeatedString := strings.Repeat(padStr, repeatTimes)

	leftSide := ""
	if padLeft {
		leftSide = repeatedString[0:targetLenLeft]
	}

	rightSide := ""
	if padRight {
		rightSide = repeatedString[0:targetLenRight]
	}

	return leftSide + str + rightSide
}

// Reverse 反转字符串
func Reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Remove 移除字符串中指定的字符串
func Remove(str, remove string) string {
	if str == "" || remove == "" {
		return remove
	}
	return strings.Replace(str, remove, "", -1)
}

// RemovePrefix 左侧移除字符串中指定的字符串
func RemovePrefix(str, prefix string) string {
	if str == "" || prefix == "" {
		return str
	}
	return strings.TrimPrefix(str, prefix)
}

// RemoveSuffix 右侧移除字符串中指定的字符串
func RemoveSuffix(str string, suffix string) string {
	if str == "" || suffix == "" {
		return str
	}
	return strings.TrimSuffix(str, suffix)
}

// RemoveAny 移除字符串中指定的字符串集
func RemoveAny(str string, removes ...string) string {
	if str == "" || len(removes) == 0 {
		return str
	}
	for _, rr := range removes {
		str = Remove(str, rr)
	}
	return str
}

// RemoveSign 将字符串的所有数据依次写成一行, 去除无意义字符串(标点符号、符号), 性能原因, 不使用 strings.NewReplacer
func RemoveSign(str string) string {
	if strings.Contains(str, LF) {
		str = strings.ReplaceAll(str, LF, "")
	}

	if strings.Contains(str, CRLF) {
		str = strings.ReplaceAll(str, CRLF, "")
	}

	if strings.Contains(str, TAB) {
		str = strings.ReplaceAll(str, TAB, "")
	}

	if strings.Contains(str, SPACE) {
		str = strings.ReplaceAll(str, SPACE, "")
	}

	m := regexp.MustCompile(`[\pP\pS]`)
	return m.ReplaceAllString(str, "")
}

// RemoveLines 移除换行符, 换行符包括 \n \r\n, 性能原因, 不使用 strings.NewReplacer
func RemoveLines(str string) string {
	if strings.Contains(str, LF) {
		str = strings.ReplaceAll(str, LF, "")
	}

	if strings.Contains(str, CRLF) {
		str = strings.ReplaceAll(str, CRLF, "")
	}

	return str
}

// SubString 字符串截取
func SubString(str string, pos, length int) string {
	runes := []rune(str)
	max := len(runes)

	if pos < 0 || length <= 0 {
		return str
	}

	if pos > max {
		return ""
	}

	l := pos + length
	if l > max {
		l = max
	}
	return string(runes[pos:l])
}

// NormaliseSpace 规范化此字符串中的空白, 多个空格合并为一个空格, 所有空白字符例如换行符、制表符, 都转换为一个简单的空格。
func NormaliseSpace(str string) string {
	str = strings.Join(strings.Fields(str), " ")

	return str
}

// NormaliseLine 规范化此字符串中的换行, 多个换行合并为一个换行
func NormaliseLine(str string) string {
	lines := SplitTrim(str, LF)
	if len(lines) > 0 {
		str = strings.Join(lines, LF)
	}

	return str
}

// Template 模板渲染
func Template(tpl string, data any) (string, error) {
	t := template.Must(template.New("").Parse(tpl))

	buf := new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		return "", err
	}

	return String(buf.Bytes()), nil
}
