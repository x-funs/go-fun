package fun

import (
	"fmt"
	"testing"
)

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

	StructCopy(a, b)

	fmt.Printf("%+v\n", b)

	bb := &bStruct{
		Name: "test-bb",
	}
	bb.Wrap = "234"
	StructCopy(a, bb)

	fmt.Printf("%+v\n", bb)

}
