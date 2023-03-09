package fun

import (
	"fmt"
	"reflect"
	"sort"
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

// SliceUnique 对数值和字符串切片进行去重(会改变元素的顺序)
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

// SliceConcat 合并多个切片, 非去重, 非原始切片
func SliceConcat[T any](slice []T, values ...[]T) []T {
	result := append([]T{}, slice...)

	for _, v := range values {
		result = append(result, v...)
	}

	return result
}

// SliceEqual 切片是否相等: 长度相同且所有元素的顺序和值相等
func SliceEqual[T comparable](slice1, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}

// SliceEvery 切片中的所有元素都满足函数，则返回 true
func SliceEvery[T any](slice []T, predicate func(index int, item T) bool) bool {
	var currentLength int

	for i, v := range slice {
		if predicate(i, v) {
			currentLength++
		}
	}

	return currentLength == len(slice)
}

// SliceNone 切片中的所有元素都不满足函数，则返回 true
func SliceNone[T any](slice []T, predicate func(index int, item T) bool) bool {
	var currentLength int

	for i, v := range slice {
		if !predicate(i, v) {
			currentLength++
		}
	}

	return currentLength == len(slice)
}

// SliceSome 切片中有一个元素满足函数，就返回true
func SliceSome[T any](slice []T, predicate func(index int, item T) bool) bool {
	for i, v := range slice {
		if predicate(i, v) {
			return true
		}
	}

	return false
}

// SliceFilter 筛选出切片中满足函数的所有元素
func SliceFilter[T any](slice []T, predicate func(index int, item T) bool) []T {
	result := make([]T, 0)

	for i, v := range slice {
		if predicate(i, v) {
			result = append(result, v)
		}
	}

	return result
}

// SliceForEach 切片中所有元素都执行函数
func SliceForEach[T any](slice []T, iteratee func(index int, item T)) {
	for i, v := range slice {
		iteratee(i, v)
	}
}

// SliceMap 切片中所有元素都执行函数, 有返回值
func SliceMap[T any, U any](slice []T, iteratee func(index int, item T) U) []U {
	result := make([]U, len(slice), cap(slice))

	for i, v := range slice {
		result[i] = iteratee(i, v)
	}

	return result
}

// SliceReduce 处理所有切片中元素得到结果
func SliceReduce[T any](slice []T, iteratee func(index int, result, item T) T, initial T) T {
	if len(slice) == 0 {
		return initial
	}

	result := iteratee(0, initial, slice[0])

	tmp := make([]T, 2)
	for i := 1; i < len(slice); i++ {
		tmp[0] = result
		tmp[1] = slice[i]
		result = iteratee(i, tmp[0], tmp[1])
	}

	return result
}

// SliceReplace 返回切片的副本，前n个元素替换为新的
func SliceReplace[T comparable](slice []T, old T, new T, n int) []T {
	result := make([]T, len(slice))
	copy(result, slice)

	for i := range result {
		if result[i] == old && n != 0 {
			result[i] = new
			n--
		}
	}

	return result
}

// SliceReplaceAll 返回切片的副本，所有匹配到的元素都替换为新的
func SliceReplaceAll[T comparable](slice []T, old T, new T) []T {
	return SliceReplace(slice, old, new, -1)
}

// SliceUnion 顺序合并且去重
func SliceUnion[T comparable](slices ...[]T) []T {
	var result []T
	contain := map[T]struct{}{}

	for _, slice := range slices {
		for _, item := range slice {
			if _, ok := contain[item]; !ok {
				contain[item] = struct{}{}
				result = append(result, item)
			}
		}
	}

	return result
}

// SliceUnionBy 顺序合并且去重, 支持自定义函数
func SliceUnionBy[T any, V comparable](predicate func(item T) V, slices ...[]T) []T {
	var result []T
	contain := map[V]struct{}{}

	for _, slice := range slices {
		for _, item := range slice {
			val := predicate(item)
			if _, ok := contain[val]; !ok {
				contain[val] = struct{}{}
				result = append(result, item)
			}
		}
	}

	return result
}

// SliceIntersection 切片交集且去重(顺序不能保证)
func SliceIntersection[T comparable](slices ...[]T) []T {
	if len(slices) == 0 {
		return []T{}
	}
	if len(slices) == 1 {
		return SliceUnique(slices[0])
	}

	reducer := func(sliceA, sliceB []T) []T {
		hashMap := make(map[T]int)
		for _, val := range sliceA {
			hashMap[val] = 1
		}

		out := make([]T, 0)
		for _, val := range sliceB {
			if v, ok := hashMap[val]; v == 1 && ok {
				out = append(out, val)
				hashMap[val]++
			}
		}
		return out
	}

	result := reducer(slices[0], slices[1])

	reduceSlice := make([][]T, 2)
	for i := 2; i < len(slices); i++ {
		reduceSlice[0] = result
		reduceSlice[1] = slices[i]
		result = reducer(reduceSlice[0], reduceSlice[1])
	}

	return result
}

// sliceValue 返回切片的反射类型
func sliceValue(slice any) reflect.Value {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Invalid slice type, value of type %T", slice))
	}
	return v
}

// SliceSortBy 根据字段排序(field的大小写应该和字段保持一致)
func SliceSortBy(slice any, field string, sortType ...string) error {
	sv := sliceValue(slice)
	t := sv.Type().Elem()

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return fmt.Errorf("data type %T not support, shuld be struct or pointer to struct", slice)
	}

	// Find the field.
	sf, ok := t.FieldByName(field)
	if !ok {
		return fmt.Errorf("field name %s not found", field)
	}

	// Create a less function based on the field's kind.
	var compare func(a, b reflect.Value) bool
	switch sf.Type.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if len(sortType) > 0 && sortType[0] == "desc" {
			compare = func(a, b reflect.Value) bool { return a.Int() > b.Int() }
		} else {
			compare = func(a, b reflect.Value) bool { return a.Int() < b.Int() }
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if len(sortType) > 0 && sortType[0] == "desc" {
			compare = func(a, b reflect.Value) bool { return a.Uint() > b.Uint() }
		} else {
			compare = func(a, b reflect.Value) bool { return a.Uint() < b.Uint() }
		}
	case reflect.Float32, reflect.Float64:
		if len(sortType) > 0 && sortType[0] == "desc" {
			compare = func(a, b reflect.Value) bool { return a.Float() > b.Float() }
		} else {
			compare = func(a, b reflect.Value) bool { return a.Float() < b.Float() }
		}
	case reflect.String:
		if len(sortType) > 0 && sortType[0] == "desc" {
			compare = func(a, b reflect.Value) bool { return a.String() > b.String() }
		} else {
			compare = func(a, b reflect.Value) bool { return a.String() < b.String() }
		}
	case reflect.Bool:
		if len(sortType) > 0 && sortType[0] == "desc" {
			compare = func(a, b reflect.Value) bool { return a.Bool() && !b.Bool() }
		} else {
			compare = func(a, b reflect.Value) bool { return !a.Bool() && b.Bool() }
		}
	default:
		return fmt.Errorf("field type %s not supported", sf.Type)
	}

	sort.Slice(slice, func(i, j int) bool {
		a := sv.Index(i)
		b := sv.Index(j)
		if t.Kind() == reflect.Ptr {
			a = a.Elem()
			b = b.Elem()
		}
		a = a.FieldByIndex(sf.Index)
		b = b.FieldByIndex(sf.Index)
		return compare(a, b)
	})

	return nil
}

// SliceColumn 返回所有行的某一列
func SliceColumn[T, V any](slice []T, key any) []V {
	values := make([]V, len(slice))
	switch reflect.TypeOf(slice).Elem().Kind() {
	case reflect.Slice, reflect.Array:
		for i, v := range slice {
			values[i] = reflect.ValueOf(v).Index(int(reflect.ValueOf(key).Int())).Interface().(V)
		}
	case reflect.Map:
		for i, v := range slice {
			values[i] = reflect.ValueOf(v).MapIndex(reflect.ValueOf(key)).Interface().(V)
		}
	case reflect.Struct:
		for i, v := range slice {
			values[i] = reflect.ValueOf(v).FieldByName(reflect.ValueOf(key).String()).Interface().(V)
		}
	}

	return values
}
