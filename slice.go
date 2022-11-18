package fun

import (
	"strconv"
	"strings"
)

// IntsToStrings int 切片转换为字符串切片
func IntsToStrings(slice []int) []string {
	if len(slice) == 0 {
		return []string{}
	}
	var str []string
	for _, v := range slice {
		str = append(str, strconv.Itoa(v))
	}
	return str
}

// StringsToInts 字符串切片转换为 int 切片
func StringsToInts(slice []string) []int {
	if len(slice) == 0 {
		return []int{}
	}
	var ints []int
	for _, v := range slice {
		if i, err := strconv.Atoi(v); err == nil {
			ints = append(ints, i)
		}
	}
	return ints
}

// SliceContains 判断整型和字符串是否在切片中
func SliceContains[T comparable](slice []T, v T) bool {
	if len(slice) == 0 {
		return false
	}

	for _, s := range slice {
		if s == v {
			return true
		}
	}
	return false
}

// SliceUnique 对数值和字符串切片进行去重
func SliceUnique[T comparable](slice []T) []T {
	if len(slice) == 0 {
		return slice
	}
	m := make(map[T]bool)
	for i := range slice {
		m[slice[i]] = true
	}

	slice = slice[:0]
	for k := range m {
		slice = append(slice, k)
	}

	return slice
}

// SliceSplit 对数值和字符串切片按照指定长度进行分割
func SliceSplit[T comparable](slice []T, size int) [][]T {
	var res [][]T

	if len(slice) == 0 || size <= 0 {
		return res
	}

	length := len(slice)
	if size == 1 || size >= length {
		for _, v := range slice {
			var tmp []T
			tmp = append(tmp, v)
			res = append(res, tmp)
		}
		return res
	}

	// divide slice equally
	divideNum := length/size + 1
	for i := 0; i < divideNum; i++ {
		if i == divideNum-1 {
			if len(slice[i*size:]) > 0 {
				res = append(res, slice[i*size:])
			}
		} else {
			res = append(res, slice[i*size:(i+1)*size])
		}
	}

	return res
}

// SliceIndex 对数值和字符串切片按照指定值进行查找
func SliceIndex[T comparable](slice []T, v T) int {
	for i, s := range slice {
		if s == v {
			return i
		}
	}
	return -1
}

// SliceLastIndex 对数值和字符串切片按照指定值进行查找, 返回最后一个匹配的索引
func SliceLastIndex[T comparable](slice []T, v T) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// SliceRemove 移除数值和字符串切片中的指定值
func SliceRemove[T comparable](slice []T, v T) []T {
	if len(slice) == 0 {
		return slice
	}

	var res []T
	for _, s := range slice {
		if s != v {
			res = append(res, s)
		}
	}
	return res
}

// SliceRemoveBlank 移除字符串切片中的空值
func SliceRemoveBlank(slice []string) []string {
	if len(slice) == 0 {
		return slice
	}

	var res []string
	for _, s := range slice {
		str := strings.TrimSpace(s)
		if len(str) > 0 {
			res = append(res, s)
		}
	}
	return res
}

// SliceTrim 对字符串切片进行 Trim, 并自动忽略空值
func SliceTrim(slice []string) []string {
	if len(slice) == 0 {
		return slice
	}

	var res []string
	for _, s := range slice {
		str := strings.TrimSpace(s)
		if len(str) > 0 {
			res = append(res, str)
		}
	}
	return res
}
