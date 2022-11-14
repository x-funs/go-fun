package fun

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type aStruct struct {
	Name    string
	Age     int
	State   bool
	Ps      *string
	CStruct cStruct
	Wrap    string
}

type bStruct struct {
	Name    string
	Age     int
	State   bool
	Float   float64
	Ps      *string
	PStr    *string
	CStruct cStruct
	wrapStruct
}

type cStruct struct {
	School string
	Grade  int
}

type wrapStruct struct {
	Wrap string
}

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
