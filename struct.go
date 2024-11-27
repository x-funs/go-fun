package fun

import (
	"errors"
	"fmt"
	"reflect"
)

// StructCopy 复制 struct 对象
func StructCopy(src, dst any) error {
	if src == nil || dst == nil {
		return errors.New("value is null")
	}

	return structCopy(reflect.ValueOf(src), reflect.ValueOf(dst))
}

// structCopy 复制 struct 对象
func structCopy(src, dst reflect.Value) error {
	tSrc := src.Type()
	tDst := dst.Type()

	if tSrc.Kind() == reflect.Ptr {
		src = src.Elem()
		tSrc = tSrc.Elem()
	}

	if tDst.Kind() == reflect.Ptr {
		dst = dst.Elem()
		tDst = tDst.Elem()
	}

	// Only struct are supported
	if tSrc.Kind() != reflect.Struct || tDst.Kind() != reflect.Struct {
		return errors.New("value is not struct")
	}

	var dstField reflect.Value
	fmt.Println(tSrc.NumField())
	for i := 0; i < tSrc.NumField(); i++ {
		anonymous := tSrc.Field(i).Anonymous
		// 如果不是嵌入字段
		if !anonymous {
			srcFieldName := tSrc.Field(i).Name
			dstField = dst.FieldByName(srcFieldName)
			if dstField.IsValid() && dstField.CanSet() {
				dstField.Set(src.Field(i))
			}
			// 递归，如果出现错误就直接返回错误
		} else {
			if err := structCopy(src.Field(i).Addr(), dst); err != nil {
				return err
			}
		}
	}

	return nil
}

// StructCompareSomeField 比较结构体的部分字段，以 some 为基准，判断 some 中的字段与 dst 的同名字段值是否相同，通常用于编辑的场景
func StructCompareSomeField(some, dst any) (bool, error) {
	if some == nil || dst == nil {
		return false, errors.New("value is nil")
	}

	vSome := reflect.ValueOf(some)
	vDst := reflect.ValueOf(dst)

	tSome := vSome.Type()
	tDst := vDst.Type()

	if tSome.Kind() == reflect.Ptr {
		vSome = vSome.Elem()
		tSome = tSome.Elem()
	}

	if tDst.Kind() == reflect.Ptr {
		vDst = vDst.Elem()
		tDst = tDst.Elem()
	}

	// Only struct are supported
	if tSome.Kind() != reflect.Struct || tDst.Kind() != reflect.Struct {
		return false, errors.New("value is not struct")
	}

	// 遍历结构体的字段
	for i := 0; i < tSome.NumField(); i++ {
		fieldA := tSome.Field(i)
		fieldB := vDst.FieldByName(fieldA.Name)

		// 如果另一个结构体中存在相同名称的字段
		if fieldB.IsValid() {
			valueA := vSome.Field(i)
			valueB := fieldB
			// 比较字段值
			if !reflect.DeepEqual(valueA.Interface(), valueB.Interface()) {
				return false, nil
			}
		} else {
			return false, errors.New("dst struct field not match")
		}
	}

	return true, nil
}
