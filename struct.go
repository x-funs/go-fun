package fun

import (
	"errors"
	"reflect"
)

// StructCopy 复制 struct 对象
func StructCopy(src, dst any) error {
	if src == nil || dst == nil {
		return errors.New("value is null")
	}

	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst)

	// 检查dst是否为指针
	if dstVal.Kind() != reflect.Ptr {
		return errors.New("dst must be a pointer to struct")
	}

	return structCopy(srcVal, dstVal)
}

// structCopy 复制 struct 对象
func structCopy(src, dst reflect.Value) error {
	tSrc := src.Type()
	tDst := dst.Type()

	// 处理src指针
	if tSrc.Kind() == reflect.Ptr {
		src = src.Elem()
		tSrc = tSrc.Elem()
	}

	// 处理dst指针
	if tDst.Kind() == reflect.Ptr {
		dst = dst.Elem()
		tDst = tDst.Elem()
	}

	// 检查是否为结构体
	if tSrc.Kind() != reflect.Struct || tDst.Kind() != reflect.Struct {
		return errors.New("value is not struct")
	}

	for i := 0; i < tSrc.NumField(); i++ {
		srcField := tSrc.Field(i)
		srcVal := src.Field(i)

		// 处理嵌入字段
		if srcField.Anonymous {
			// 检查嵌入字段是否为结构体
			if srcVal.Kind() == reflect.Struct {
				// 递归复制嵌入字段
				if err := structCopy(srcVal.Addr(), dst.Addr()); err != nil {
					return err
				}
			}
			continue
		}

		// 处理普通字段
		dstField := dst.FieldByName(srcField.Name)
		if !dstField.IsValid() {
			continue // 目标结构体中没有该字段，跳过
		}

		if !dstField.CanSet() {
			continue // 字段不可设置，跳过（如未导出字段）
		}

		// 检查字段类型是否匹配
		if srcVal.Type() != dstField.Type() {
			continue // 类型不匹配，跳过
		}

		dstField.Set(srcVal)
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
