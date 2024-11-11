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
	dstA := &aStruct{
		Name:  "test-a",
		Age:   12,
		State: true,
	}
	someA := &someAStruct{
		Name:  "test-a",
		State: true,
	}

	resultA, err := StructCompareSomeField(someA, dstA)
	t.Log(resultA)
	t.Log(err)

	dstB := &bStruct{
		Name: "test-b",
		Age:  18,
	}
	dstB.Wrap = "abc"
	someB := &someBStruct{
		Name:  "test-b",
		Age:   18,
		Float: 0,
		Wrap:  "abc",
	}
	resultB, err := StructCompareSomeField(someB, dstB)
	t.Log(resultB)
	t.Log(err)
}
