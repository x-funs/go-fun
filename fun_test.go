package fun

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemory(t *testing.T) {
	fmt.Println(MemoryBytes())
	fmt.Println(Memory(SizeKB))
	fmt.Println(Memory(SizeMB))
}

func TestEmpty(t *testing.T) {
	assert.Equal(t, true, Empty(nil))
	assert.Equal(t, true, Empty(0))
	assert.Equal(t, true, Empty(""))
	assert.Equal(t, true, Empty(false))

	assert.Equal(t, false, Empty(" "))
	assert.Equal(t, false, Empty(1))
	assert.Equal(t, false, Empty(true))

	var a1 [0]int
	a2 := [2]int{3, 5}
	assert.Equal(t, true, Empty(a1))
	assert.Equal(t, false, Empty(a2))

	var s1 []string
	s2 := []int{2, 3, 5}
	assert.Equal(t, true, Empty(s1))
	assert.Equal(t, false, Empty(s2))

	m1 := map[string]int{}
	m2 := map[string]string{"k1": "v1", "k2": "v2"}
	assert.Equal(t, true, Empty(m1))
	assert.Equal(t, false, Empty(m2))

	assert.Equal(t, true, EmptyAll("", nil))
	assert.Equal(t, true, EmptyAny("a", "", "b"))
}

func TestCommand(t *testing.T) {
	result, _ := Command("ls", []string{"-l"}, "/")
	fmt.Println(String(result))
}

func TestIf(t *testing.T) {
	assert.Equal(t, "a", If(true, "a", "b"))
	assert.Equal(t, 2, If(false, 1, 2))

	datas := []string{"a", "b", "c"}
	assert.Equal(t, []string{"d"}, If(Empty(datas), []string{"c"}, []string{"d"}))
}
