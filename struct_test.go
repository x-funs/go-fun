package fun

import (
	"testing"
)

type User struct {
	ID   int
	Name string
	Age  int
}

type UserWithAddress struct {
	User
	Address string
}

type UserCopy struct {
	ID   int
	Name string
	Age  int
	City string
}

func TestStructCopy(t *testing.T) {
	// 测试普通结构体复制
	src := User{ID: 1, Name: "Alice", Age: 30}
	var dst User

	err := StructCopy(&src, &dst)
	if err != nil {
		t.Fatal(err)
	}

	if dst.ID != src.ID || dst.Name != src.Name || dst.Age != src.Age {
		t.Error("结构体复制失败")
	}

	// 测试嵌入结构体复制
	srcWithAddr := UserWithAddress{
		User:    User{ID: 2, Name: "Bob", Age: 25},
		Address: "Beijing",
	}
	var dstWithAddr UserWithAddress

	err = StructCopy(&srcWithAddr, &dstWithAddr)
	if err != nil {
		t.Fatal(err)
	}

	if dstWithAddr.ID != srcWithAddr.ID || dstWithAddr.Name != srcWithAddr.Name || dstWithAddr.Address != srcWithAddr.Address {
		t.Error("嵌入结构体复制失败")
	}

	// 测试部分字段复制
	src2 := User{ID: 3, Name: "Charlie", Age: 35}
	var dst2 UserCopy

	err = StructCopy(&src2, &dst2)
	if err != nil {
		t.Fatal(err)
	}

	if dst2.ID != src2.ID || dst2.Name != src2.Name || dst2.Age != src2.Age {
		t.Error("部分字段复制失败")
	}

	// 测试非指针dst
	err = StructCopy(&src, dst)
	if err == nil {
		t.Error("非指针dst应该返回错误")
	}

	// 测试类型不匹配
	src3 := User{ID: 4, Name: "David", Age: 40}
	var dst3 struct {
		ID   string
		Name string
	}

	err = StructCopy(&src3, &dst3)
	if err != nil {
		t.Fatal(err)
	}

	// ID类型不匹配，应该被跳过
	if dst3.ID != "" {
		t.Error("类型不匹配的字段应该被跳过")
	}

	// 测试未导出字段
	src4 := struct {
		ID   int
		name string // 未导出字段
	}{ID: 5, name: "Eve"}
	var dst4 struct {
		ID   int
		name string
	}

	err = StructCopy(&src4, &dst4)
	if err != nil {
		t.Fatal(err)
	}

	if dst4.ID != src4.ID || dst4.name != "" {
		t.Error("未导出字段应该被跳过")
	}

	// 测试nil指针
	err = StructCopy(nil, &dst)
	if err == nil {
		t.Error("nil src应该返回错误")
	}

	err = StructCopy(&src, nil)
	if err == nil {
		t.Error("nil dst应该返回错误")
	}

	// 测试递归嵌入结构体
	src5 := struct {
		User
		Level int
	}{User: User{ID: 6, Name: "Frank", Age: 45}, Level: 1}
	var dst5 struct {
		User
		Level int
	}

	err = StructCopy(&src5, &dst5)
	if err != nil {
		t.Fatal(err)
	}

	if dst5.ID != src5.ID || dst5.Name != src5.Name || dst5.Level != src5.Level {
		t.Error("递归嵌入结构体复制失败")
	}

	// 测试指针字段
	src6 := struct {
		ID   *int
		Name string
	}{ID: new(int), Name: "Grace"}
	*src6.ID = 7
	var dst6 struct {
		ID   *int
		Name string
	}

	err = StructCopy(&src6, &dst6)
	if err != nil {
		t.Fatal(err)
	}

	if dst6.Name != src6.Name || *dst6.ID != *src6.ID {
		t.Error("指针字段复制失败")
	}

	// 测试切片字段
	src7 := struct {
		ID   int
		Tags []string
	}{ID: 8, Tags: []string{"go", "test"}}
	var dst7 struct {
		ID   int
		Tags []string
	}

	err = StructCopy(&src7, &dst7)
	if err != nil {
		t.Fatal(err)
	}

	if dst7.ID != src7.ID || len(dst7.Tags) != len(src7.Tags) || dst7.Tags[0] != src7.Tags[0] {
		t.Error("切片字段复制失败")
	}

	// 测试map字段复制
	src8 := struct {
		ID   int
		Data map[string]int
	}{ID: 9, Data: map[string]int{"a": 1, "b": 2}}
	var dst8 struct {
		ID   int
		Data map[string]int
	}

	err = StructCopy(&src8, &dst8)
	if err != nil {
		t.Fatal(err)
	}

	if dst8.ID != src8.ID || dst8.Data["a"] != src8.Data["a"] || dst8.Data["b"] != src8.Data["b"] {
		t.Error("map字段复制失败")
	}
}

func TestStructCompareSomeField(t *testing.T) {
	// 测试普通结构体比较
	src := User{ID: 1, Name: "Alice", Age: 30}
	dst := User{ID: 1, Name: "Alice", Age: 30}

	match, err := StructCompareSomeField(&src, &dst)
	if err != nil {
		t.Fatal(err)
	}

	if !match {
		t.Error("相同结构体应该匹配")
	}

	// 测试不同结构体比较
	dst2 := User{ID: 1, Name: "Alice", Age: 31}
	match, err = StructCompareSomeField(&src, &dst2)
	if err != nil {
		t.Fatal(err)
	}

	if match {
		t.Error("不同结构体应该不匹配")
	}

	// 测试嵌入结构体比较
	srcWithAddr := UserWithAddress{
		User:    User{ID: 2, Name: "Bob", Age: 25},
		Address: "Beijing",
	}
	dstWithAddr := UserWithAddress{
		User:    User{ID: 2, Name: "Bob", Age: 25},
		Address: "Beijing",
	}

	match, err = StructCompareSomeField(&srcWithAddr, &dstWithAddr)
	if err != nil {
		t.Fatal(err)
	}

	if !match {
		t.Error("相同嵌入结构体应该匹配")
	}

	// 测试部分字段比较
	src2 := User{ID: 3, Name: "Charlie", Age: 35}
	dst3 := UserCopy{ID: 3, Name: "Charlie", Age: 35, City: "Shanghai"}

	match, err = StructCompareSomeField(&src2, &dst3)
	if err != nil {
		t.Fatal(err)
	}

	if !match {
		t.Error("部分字段匹配应该返回true")
	}

	// 测试字段不匹配
	src3 := User{ID: 4, Name: "David", Age: 40}
	dst4 := UserCopy{ID: 4, Name: "David", Age: 41, City: "Beijing"}

	match, err = StructCompareSomeField(&src3, &dst4)
	if err != nil {
		t.Fatal(err)
	}

	if match {
		t.Error("字段不匹配应该返回false")
	}

	// 测试nil指针
	match, err = StructCompareSomeField(nil, &dst)
	if err == nil {
		t.Error("nil src应该返回错误")
	}

	match, err = StructCompareSomeField(&src, nil)
	if err == nil {
		t.Error("nil dst应该返回错误")
	}

	// 测试指针指向nil
	var nilSrc *User
	match, err = StructCompareSomeField(nilSrc, &dst)
	if err == nil {
		t.Error("nil src指针应该返回错误")
	}

	// 测试类型不匹配
	src4 := User{ID: 5, Name: "Eve", Age: 28}
	dst5 := struct {
		ID   string
		Name string
	}{ID: "5", Name: "Eve"}

	match, err = StructCompareSomeField(&src4, &dst5)
	if err == nil {
		t.Error("类型不匹配应该返回错误")
	}
}
