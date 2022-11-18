package fun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceContains(t *testing.T) {
	assert.Equal(t, true, SliceContains([]string{"a", "b", "c"}, "a"))
	assert.Equal(t, true, SliceContains([]int{3, 5, 7}, 7))
	assert.Equal(t, false, SliceContains([]int{3, 5, 7}, 2))
	assert.Equal(t, false, SliceContains([]int{}, 2))
}

func TestSliceUnique(t *testing.T) {
	assert.Equal(t, 3, len(SliceUnique([]string{"a", "b", "c", "a", "b", "c"})))
	assert.Equal(t, []string{""}, SliceUnique([]string{"", "", ""}))
}

func TestSliceSplit(t *testing.T) {
	assert.Equal(t, [][]int{{2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}, SliceSplit([]int{2, 3, 4, 5, 6, 7, 8, 9}, 1))
	assert.Equal(t, [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}}, SliceSplit([]int{2, 3, 4, 5, 6, 7, 8, 9}, 2))
	assert.Equal(t, [][]int{{2, 3, 4}, {5, 6, 7}, {8, 9}}, SliceSplit([]int{2, 3, 4, 5, 6, 7, 8, 9}, 3))
	assert.Equal(t, [][]int{{2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}, SliceSplit([]int{2, 3, 4, 5, 6, 7, 8, 9}, 9))
	assert.Equal(t, [][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g"}}, SliceSplit([]string{"a", "b", "c", "d", "e", "f", "g"}, 3))
}

func TestSliceRemove(t *testing.T) {
	assert.Equal(t, []int{}, SliceRemove([]int{}, 0))
	assert.Equal(t, []int{2, 4, 5}, SliceRemove([]int{2, 0, 4, 5}, 0))
	assert.Equal(t, []int{2, 0, 4, 5}, SliceRemove([]int{2, 0, 4, 5}, 1))
	assert.Equal(t, []string{"b", "c", "d"}, SliceRemove([]string{"a", "b", "c", "d"}, "a"))
	assert.Equal(t, []string{"a", "d"}, SliceRemove([]string{"a", "", "", "d"}, ""))
}

func TestSliceTrim(t *testing.T) {
	assert.Equal(t, []string{"b", "c", "d"}, SliceTrim([]string{"b", "c", "d"}))
	assert.Equal(t, []string{"b", "c", "d"}, SliceTrim([]string{" b ", "c", "    d "}))
	assert.Equal(t, []string{"a", "b"}, SliceTrim([]string{" a  ", "b", " ", "   	"}))
}

func TestSliceRemoveBlank(t *testing.T) {
	var s []string
	assert.Equal(t, []string{"a", "b"}, SliceRemoveBlank([]string{"a", "b"}))
	assert.Equal(t, []string{"b"}, SliceRemoveBlank([]string{"", "b"}))
	assert.Equal(t, []string{" a "}, SliceRemoveBlank([]string{" a ", " "}))
	assert.Equal(t, s, SliceRemoveBlank([]string{"  ", " "}))
}

func TestSliceIndex(t *testing.T) {
	assert.Equal(t, -1, SliceIndex([]string{"a", "b", "c", "d", "c"}, "e"))
	assert.Equal(t, 1, SliceIndex([]string{"a", "b", "c", "d", "c"}, "b"))
	assert.Equal(t, 2, SliceIndex([]string{"a", "b", "c", "d", "c"}, "c"))
	assert.Equal(t, 4, SliceLastIndex([]string{"a", "b", "c", "d", "c"}, "c"))
}

func TestIntsStrings(t *testing.T) {
	assert.Equal(t, []string{"1", "2", "3"}, IntsToStrings([]int{1, 2, 3}))
	assert.Equal(t, []string{}, IntsToStrings([]int{}))
	assert.Equal(t, []int{}, StringsToInts([]string{}))
	assert.Equal(t, []int{23, 45}, StringsToInts([]string{"23", "45"}))
	assert.Equal(t, []int{12}, StringsToInts([]string{"a", "12"}))
}
