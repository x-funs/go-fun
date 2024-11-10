package fun

import (
	"fmt"
	"testing"
)

func BenchmarkCopyStruct(b *testing.B) {
	str := "123123"
	a := &aStruct{
		Name:    "test-a",
		Age:     12,
		State:   true,
		CStruct: cStruct{School: "ac", Grade: 1},
		Ps:      &str,
		Wrap:    "123",
	}
	bs := &bStruct{}

	for i := 0; i < b.N; i++ {
		StructCopy(a, bs)
	}
}

func TestCopyStruct(t *testing.T) {
	str := "s"
	a := &aStruct{
		Name:    "test-a",
		Age:     12,
		State:   true,
		CStruct: cStruct{School: "ac", Grade: 1},
	}
	a.Ps = &str
	a.Wrap = "123"
	fmt.Printf("%+v\n", a)

	b := &bStruct{}

	err := StructCopy(a, b)
	if err != nil {
		return
	}

	fmt.Printf("%+v\n", b)
	fmt.Printf("%+v\n", err)

	bb := &bStruct{
		Name: "test-bb",
	}
	bb.Wrap = "234"
	err = StructCopy(a, bb)
	if err != nil {
		return
	}

	fmt.Printf("%+v\n", bb)
	fmt.Printf("%+v\n", err)

}

func TestStructCompareSomeField(t *testing.T) {
	dst := aStruct{
		Name:  "test-a",
		Age:   12,
		State: true,
	}

	some := aStruct{
		Name: "test-a",
	}

	result, err := StructCompareSomeField(some, dst)

	t.Log(result)
	t.Log(err)
}
