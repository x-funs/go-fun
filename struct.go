package fun

import "reflect"

// CopyStruct 复制 struct 对象
func CopyStruct(src, dst any) {
	if src == nil || dst == nil {
		return
	}
	copyStruct(reflect.ValueOf(src), reflect.ValueOf(dst))
}

// copyStruct 复制 struct 对象
func copyStruct(src, dst reflect.Value) {
	st := src.Type()
	dt := dst.Type()
	if st.Kind() == reflect.Ptr {
		src = src.Elem()
		st = st.Elem()
	}
	if dt.Kind() == reflect.Ptr {
		dst = dst.Elem()
		dt = dt.Elem()
	}

	// Only struct are supported
	if st.Kind() != reflect.Struct || dt.Kind() != reflect.Struct {
		return
	}
	var field reflect.Value
	for i := 0; i < st.NumField(); i++ {
		if !st.Field(i).Anonymous {
			field = dst.FieldByName(st.Field(i).Name)
			if field.IsValid() && field.CanSet() {
				field.Set(src.Field(i))
			}
		} else {
			copyStruct(src.Field(i).Addr(), dst)
		}
	}
}
