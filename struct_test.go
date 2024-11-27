package fun

import (
	"fmt"
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

type someAStruct struct {
	Name  string
	State bool
}

type someBStruct struct {
	Name  string
	Age   int
	Float float64
	Wrap  string
}

type cStruct struct {
	School string
	Grade  int
}

type wrapStruct struct {
	Wrap string
}

type Primary struct {
	Id   uint
	Name string
}

type Member struct {
	Primary
	Username string
}

type MemberLite struct {
	Id       uint
	Username string
}

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

func TestCopyStructMember(t *testing.T) {
	member := &Member{}
	member.Id = 2
	member.Name = "name"
	member.Username = "username"

	memberLite := &MemberLite{}

	err := StructCopy(member, memberLite)
	fmt.Printf("member: %+v\n", member)
	fmt.Printf("memberLite: %+v\n", memberLite)
	fmt.Printf("err: %+v\n", err)
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
	fmt.Printf("a: %+v\n", a)

	b := &bStruct{}

	err := StructCopy(a, b)
	if err != nil {
		return
	}

	fmt.Printf("b: %+v\n", b)
	fmt.Printf("err: %+v\n", err)

	bb := &bStruct{
		Name: "test-bb",
	}
	bb.Wrap = "234"
	err = StructCopy(a, bb)
	if err != nil {
		return
	}

	fmt.Printf("bb: %+v\n", bb)
	fmt.Printf("err: %+v\n", err)
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
