package fun

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntsStrings(t *testing.T) {
	assert.Equal(t, []string{"1", "2", "3"}, IntsToStrings([]int{1, 2, 3}))
	assert.Equal(t, []string{}, IntsToStrings([]int{}))
	assert.Equal(t, []string{"1", "4900000000", "-7149"}, IntsToStrings([]int{1.0, 4900000000, -0x1bed}))
	assert.Equal(t, []int{}, StringsToInts([]string{}))
	assert.Equal(t, []int{23, 45}, StringsToInts([]string{"23", "45"}))
	assert.Equal(t, []int{12}, StringsToInts([]string{"a", "12"}))
	assert.Equal(t, []int{1, 4900000000, -7149}, StringsToInts([]string{"1", "4900000000", "-7149"}))
}

func BenchmarkIntsToStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntsToStrings([]int{1.0, 4900000000, -0x1bed})
	}
}

func BenchmarkStringsToInts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringsToInts([]string{"1", "4900000000", "-7149"})
	}
}

func TestSliceContains(t *testing.T) {
	assert.Equal(t, true, SliceContains([]string{"a", "b", "c"}, "a"))
	assert.Equal(t, true, SliceContains([]int{3, 5, 7}, 7))
	assert.Equal(t, false, SliceContains([]int{3, 5, 7}, 2))
	assert.Equal(t, false, SliceContains([]int{}, 2))
}

func BenchmarkContainers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceContains([]string{"red", "blue", "orange", "yellow", "green"}, "green")
	}
}

func TestSliceUnique(t *testing.T) {
	assert.Equal(t, 3, len(SliceUnique([]string{"a", "b", "c", "a", "b", "c"})))
	assert.Equal(t, []string{""}, SliceUnique([]string{"", "", ""}))
}

func BenchmarkSliceUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceUnique([]string{"red", "blue", "orange", "yellow", "blue", "green", "red", "green"})
	}
}

func TestSliceSplit(t *testing.T) {
	assert.Equal(t, [][]int{{2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}, SliceSplit([]int{2, 3, 4, 5, 6, 7, 8, 9}, 1))
	assert.Equal(t, [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}}, SliceSplit([]int{2, 3, 4, 5, 6, 7, 8, 9}, 2))
	assert.Equal(t, [][]int{{2, 3, 4}, {5, 6, 7}, {8, 9}}, SliceSplit([]int{2, 3, 4, 5, 6, 7, 8, 9}, 3))
	assert.Equal(t, [][]int{{2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}, SliceSplit([]int{2, 3, 4, 5, 6, 7, 8, 9}, 9))
	assert.Equal(t, [][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g"}}, SliceSplit([]string{"a", "b", "c", "d", "e", "f", "g"}, 3))
}

func BenchmarkSliceSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceSplit([]int{2, 3, 4, 5, 6, 7, 8, 9}, 1)
	}
}

func TestSliceRemove(t *testing.T) {
	assert.Equal(t, []int{}, SliceRemove([]int{}, 0))
	assert.Equal(t, []int{2, 4, 5}, SliceRemove([]int{2, 0, 4, 5}, 0))
	assert.Equal(t, []int{2, 0, 4, 5}, SliceRemove([]int{2, 0, 4, 5}, 1))
	assert.Equal(t, []string{"b", "c", "d"}, SliceRemove([]string{"a", "b", "c", "d"}, "a"))
	assert.Equal(t, []string{"a", "d"}, SliceRemove([]string{"a", "", "", "d"}, ""))
}

func BenchmarkSliceRemove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceRemove([]string{"a", "", "", "d"}, "")
	}
}

func TestSliceTrim(t *testing.T) {
	assert.Equal(t, []string{"b", "c", "d"}, SliceTrim([]string{"b", "c", "d"}))
	assert.Equal(t, []string{"b", "c", "d"}, SliceTrim([]string{" b ", "c", "    d "}))
	assert.Equal(t, []string{"a", "b"}, SliceTrim([]string{" a  ", "b", " ", "   	"}))
}

func BenchmarkSliceTrim(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceTrim([]string{" a  ", "b", " ", "   	"})
	}
}

func TestSliceRemoveBlank(t *testing.T) {
	var s []string
	assert.Equal(t, []string{"a", "b"}, SliceRemoveBlank([]string{"a", "b"}))
	assert.Equal(t, []string{"b"}, SliceRemoveBlank([]string{"", "b"}))
	assert.Equal(t, []string{" a "}, SliceRemoveBlank([]string{" a ", " "}))
	assert.Equal(t, s, SliceRemoveBlank([]string{"  ", " "}))
}

func BenchmarkSliceRemoveBlank(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceRemoveBlank([]string{" a ", " ", "greed is good "})
	}
}

func TestSliceIndex(t *testing.T) {
	assert.Equal(t, -1, SliceIndex([]string{"a", "b", "c", "d", "c"}, "e"))
	assert.Equal(t, 1, SliceIndex([]string{"a", "b", "c", "d", "c"}, "b"))
	assert.Equal(t, 2, SliceIndex([]string{"a", "b", "c", "d", "c"}, "c"))
}

func BenchmarkSliceIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceIndex([]string{"a", "b", "c", "d", "c"}, "c")
	}
}

func TestSliceLastIndex(t *testing.T) {
	assert.Equal(t, 4, SliceLastIndex([]string{"a", "b", "c", "d", "c"}, "c"))
}

func BenchmarkSliceLastIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceLastIndex([]string{"a", "b", "c", "d", "c"}, "c")
	}
}

func TestSliceConcat(t *testing.T) {
	assert.Equal(t, []string{"apple", "banana", "orange", "red", "black", "blue"}, SliceConcat([]string{"apple", "banana", "orange"}, []string{"red", "black", "blue"}))
	assert.Equal(t, []int{147, 258, 369, 456, 123, 987}, SliceConcat([]int{147, 258, 369}, []int{456, 123, 987}))
}

func BenchmarkSliceConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceConcat([]string{"apple", "banana", "orange"}, []string{"red", "black", "blue"})
	}
}

func TestSliceEqual(t *testing.T) {
	assert.Equal(t, false, SliceEqual([]int{147, 258, 369}, []int{456, 123, 987}))
	assert.Equal(t, false, SliceEqual([]int{147, 258, 369}, []int{258, 147, 369}))
	assert.Equal(t, true, SliceEqual([]string{"147", "258", "369"}, []string{"147", "258", "369"}))
}

func BenchmarkSliceEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceEqual([]string{"apple", "banana", "orange", "red", "black", "blue"}, []string{"apple", "banana", "orange", "red", "black", "bleu"})
	}
}

func TestSliceEvery(t *testing.T) {
	// 是否都是正数
	assert.Equal(t, true, SliceEvery([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(index, item int) bool {
		return item > 0
	}))

	// 是否都是奇数
	assert.Equal(t, false, SliceEvery([]int{1, 3, 5, 6, 7, 9}, func(index, item int) bool {
		return item%2 == 1
	}))
}

func BenchmarkSliceEvery(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceEvery([]int{1, 3, 5, 6, 7, 9}, func(index, item int) bool {
			return item%2 == 1
		})
	}
}

func TestSliceNone(t *testing.T) {
	// 没有不及格的
	assert.Equal(t, false, SliceNone([]int{60, 65, 58, 78, 97, 99}, func(index, item int) bool {
		return item < 60
	}))

	// 没有包含美国的词语
	assert.Equal(t, true, SliceNone([]string{"中国", "美丽坚果", "小日子", "俄罗斯"}, func(index int, item string) bool {
		return strings.Contains(item, "美国")
	}))
}

func BenchmarkSliceNone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceNone([]int{60, 65, 58, 78, 97, 99}, func(index, item int) bool {
			return item < 60
		})
	}
}

func TestSliceSome(t *testing.T) {
	// 是否有分数不及格
	assert.Equal(t, false, SliceSome([]int{60, 65, 78, 78, 97, 99}, func(index, item int) bool {
		return item < 60
	}))

	// 是否有长度大于10的单词
	assert.Equal(t, true, SliceSome([]string{"doc", "document", "dictionary", "informational", "international"}, func(index int, item string) bool {
		return len(item) > 10
	}))
}

func BenchmarkSliceSome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceSome([]int{60, 65, 78, 78, 97, 99}, func(index, item int) bool {
			return item < 60
		})
	}
}

func TestSliceFilter(t *testing.T) {
	// 筛选出不及格的分数
	assert.Equal(t, []int{59, 5}, SliceFilter([]int{60, 59, 78, 5, 71, 83, 97, 99}, func(index, item int) bool {
		return item < 60
	}))
}

func BenchmarkSliceFilter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceFilter([]int{60, 59, 78, 83, 97, 99}, func(index, item int) bool {
			return item < 60
		})
	}
}

func TestSliceForEach(t *testing.T) {
	var temp []int
	// 低于60的分数统一增加20分, 最高100分
	SliceForEach([]int{60, 59, 78, 71, 83, 97, 99}, func(index, item int) {
		if item < 60 {
			temp = append(temp, item+20)
		} else {
			temp = append(temp, item)
		}
	})
	assert.Equal(t, []int{60, 79, 78, 71, 83, 97, 99}, temp)
}

func BenchmarkSliceForEach(b *testing.B) {
	var temp []int
	for i := 0; i < b.N; i++ {
		SliceForEach([]int{60, 59, 78, 71, 83, 97, 99}, func(index, item int) {
			if item < 60 {
				temp = append(temp, item+20)
			} else {
				temp = append(temp, item)
			}
		})
	}
}

func TestSliceMap(t *testing.T) {
	// 低于60的分数统一增加20分, 最高100分
	assert.Equal(t, []int{60, 79, 78, 71, 83, 97, 99}, SliceMap([]int{60, 59, 78, 71, 83, 97, 99}, func(index, item int) int {
		if item < 60 {
			return item + 20
		}
		return item
	}))
}

func BenchmarkSliceMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceMap([]int{60, 59, 78, 71, 83, 97, 99}, func(index, item int) int {
			if item < 60 {
				return item + 20
			}
			return item
		})
	}
}

func TestSliceReduce(t *testing.T) {
	// 所有元素的和
	assert.Equal(t, 45, SliceReduce([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(index, result, item int) int {
		return result + item
	}, 0))
}

func BenchmarkSliceReduce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceReduce([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(index, result, item int) int {
			return result + item
		}, 0)
	}
}

func TestSliceReplace(t *testing.T) {
	assert.Equal(t, []int{3, -1, -1, 0, 3}, SliceReplace([]int{0, -1, -1, 0, 3}, 0, 3, 1))
	assert.Equal(t, []int{2, -1, -1, 2, 3}, SliceReplaceAll([]int{0, -1, -1, 0, 3}, 0, 2))
}

func BenchmarkSliceReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceReplaceAll([]int{0, -1, -1, 0, 3}, 0, 2)
	}
}

func TestSliceUnion(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5}, SliceUnion([]int{1, 2, 3}, []int{1, 2, 4}, []int{1, 2, 5}))
	assert.Equal(t, []string{"123", "124", "125"}, SliceUnion([]string{"123", "124"}, []string{"124", "125"}, []string{"123", "125"}))
}

func BenchmarkSliceUnion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceUnion([]string{"123", "124"}, []string{"124", "125"}, []string{"123", "125"})
	}
}

func TestSliceUnionBy(t *testing.T) {
	type Node struct {
		Code  string
		Value int
	}
	by := func(node Node) string {
		return node.Code
	}

	expect := []Node{{Code: "123", Value: 2}, {Code: "124", Value: 3}, {Code: "125", Value: 7}}
	assert.Equal(t, expect, SliceUnionBy(by,
		[]Node{{Code: "123", Value: 2}, {Code: "124", Value: 3}},
		[]Node{{Code: "124", Value: 3}, {Code: "125", Value: 7}},
		[]Node{{Code: "123", Value: 2}, {Code: "124", Value: 3}, {Code: "125", Value: 7}},
	))
}

func BenchmarkSliceUnionBy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceUnionBy(func(item string) string {
			return item
		}, []string{"123", "124"}, []string{"124", "125"}, []string{"123", "125"})
	}
}

func TestSliceIntersection(t *testing.T) {
	// 1~9中即为奇数又为平方数的数字
	assert.Equal(t, []int{1, 9}, SliceIntersection([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{1, 3, 5, 7, 9, 11}, []int{1, 4, 9, 16, 25}))

	// 几个人的共同爱好
	assert.Equal(t, []string{"旅游", "美食"}, SliceIntersection(
		[]string{"美食", "棋类", "数独", "密室逃脱", "书法", "绘画", "旅游"},
		[]string{"哲学", "旅游", "历史", "心理学", "吉他", "美食"},
		[]string{"旅游", "美食", "看书", "看电影", "听音乐", "学技术"},
		[]string{"乒乓球", "篮球", "足球", "台球", "羽毛球", "旅游", "美食"},
	))
}

func BenchmarkSliceIntersection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceIntersection([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{1, 3, 5, 7, 9, 11}, []int{1, 4, 9, 16, 25})
	}
}

func TestSortBy(t *testing.T) {
	// 商品的销量
	type goods struct {
		Name string
		Num  int
	}

	goodsList := []goods{
		{Name: "手机", Num: 20000},
		{Name: "衣服", Num: 30000},
		{Name: "面膜", Num: 10000},
		{Name: "口红", Num: 15000},
	}
	sortedGoodsList := []goods{
		{Name: "衣服", Num: 30000},
		{Name: "手机", Num: 20000},
		{Name: "口红", Num: 15000},
		{Name: "面膜", Num: 10000},
	}
	_ = SliceSortBy(goodsList, "Num", "desc")
	assert.Equal(t, sortedGoodsList, goodsList)
}

func BenchmarkSortBy(b *testing.B) {
	type goods struct {
		Name string
		Num  int
	}
	goodsList := []goods{
		{Name: "手机", Num: 20000},
		{Name: "衣服", Num: 30000},
		{Name: "面膜", Num: 10000},
		{Name: "口红", Num: 15000},
	}
	for i := 0; i < b.N; i++ {
		_ = SliceSortBy(goodsList, "Num", "desc")
	}
}

func TestSliceColumn(t *testing.T) {
	// 数组/切片使用
	assert.Equal(t, []int{2, 9, 16, 23},
		SliceColumn[[]int, int]([][]int{
			{1, 2, 3, 4, 5, 6, 7},
			{8, 9, 10, 11, 12, 13, 14},
			{15, 16, 17, 18, 19, 20, 21},
			{22, 23, 24, 25, 26, 27, 28},
		}, 1),
	)

	// map 使用
	assert.Equal(t, []string{"YF4133", "MM8541", "KH0002", "SJ9642"},
		SliceColumn[map[string]string, string]([]map[string]string{
			{"name": "衣服", "code": "YF4133"},
			{"name": "面膜", "code": "MM8541"},
			{"name": "口红", "code": "KH0002"},
			{"name": "手机", "code": "SJ9642"},
		}, "code"),
	)

	// struct 使用
	type user struct {
		Id   int
		Name string
	}
	assert.Equal(t, []string{"admin", "user", "guest", "test"},
		SliceColumn[user, string]([]user{
			{Id: 1, Name: "admin"},
			{Id: 23, Name: "user"},
			{Id: 42, Name: "guest"},
			{Id: 169, Name: "test"},
		}, "Name"),
	)
}
