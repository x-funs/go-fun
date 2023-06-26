package fun

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"unsafe"
)

// GenInteger 整型范型集合
type GenInteger interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// GenFloat 浮点型范型集合
type GenFloat interface {
	float32 | float64
}

// GenNumber 数值范型集合
type GenNumber interface {
	GenInteger | GenFloat
}

// MemoryBytes 返回当前主要的内存指标信息 (ReadMemStats 会 stopTheWorld, 谨慎非频繁使用)
func MemoryBytes() map[string]int64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	maps := make(map[string]int64)
	maps["sys"] = int64(m.Sys)
	maps["heapSys"] = int64(m.HeapSys)
	maps["heapInuse"] = int64(m.HeapInuse)
	maps["heapAlloc"] = int64(m.HeapAlloc)

	return maps
}

// Memory 指定格式返回当前主要的内存指标信息, (ReadMemStats 会 stopTheWorld, 谨慎非频繁使用)
func Memory(format string) map[string]int64 {
	m := MemoryBytes()
	for k, v := range m {
		if v > 0 {
			switch format {
			case SizeB:
				m[k] = v
			case SizeKB:
				m[k] = v / BytesPerKB
			case SizeMB:
				m[k] = v / BytesPerMB
			case SizeGB:
				m[k] = v / BytesPerGB
			case SizeTB:
				m[k] = v / BytesPerTB
			case SizePB:
				m[k] = v / BytesPerPB
			case SizeEB:
				m[k] = v / BytesPerEB
			default:
				m[k] = v
			}
		}
	}

	return m
}

// EmptyAll 判断是否全部为空
func EmptyAll(values ...any) bool {
	if len(values) == 0 {
		return true
	}

	for _, v := range values {
		if !Empty(v) {
			return false
		}
	}

	return true
}

// EmptyAny 判断是否任意一个为空
func EmptyAny(values ...any) bool {
	if len(values) == 0 {
		return true
	}

	for _, v := range values {
		if Empty(v) {
			return true
		}
	}

	return false
}

// Empty 判断是否为空, 支持字符串、数值、数组、Slice、Map
func Empty(value any) bool {
	if value == nil {
		return true
	}

	switch value.(type) {
	case string:
		return value == ""
	case int, int8, int16, int32, int64:
		return value == 0
	case uint, uint8, uint16, uint32, uint64:
		return value == 0
	case float32, float64:
		return value == 0
	case bool:
		return value == false
	default:
		r := reflect.ValueOf(value)
		switch r.Kind() {
		case reflect.Slice, reflect.Map:
			return r.Len() == 0 || r.IsNil()
		case reflect.Array:
			return r.Len() == 0
		case reflect.Ptr, reflect.Interface:
			return r.IsNil()
		}
	}

	return false
}

// Bytes 更高效的字符串转字节数组
func Bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// String 更高效的字节数组转字符串
func String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Command 执行系统命令
func Command(bin string, argv []string, baseDir string) ([]byte, error) {
	cmd := exec.Command(bin, argv...)

	if baseDir != "" {
		cmd.Dir = baseDir
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		return stdout.Bytes(), fmt.Errorf("%s: %s", err, stderr.Bytes())
	}

	return stdout.Bytes(), nil
}

// If 三元运算符
func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}

	return falseVal
}

// Mkdir 创建一个目录，如果目录已存在则忽略
func Mkdir(dir string, perm os.FileMode) error {
	if !FileExists(dir) {
		return os.Mkdir(dir, perm)
	}

	return nil
}

// FileExists 检测目录或者文件是否存在，返回 bool
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// WriteFile 写入文件
func WriteFile(name string, data []byte, flag int, perm os.FileMode, sync bool) error {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return err
	}

	_, err = f.Write(data)

	if sync {
		_ = f.Sync()
	}

	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}

	return err
}

// WriteFileAppend 追加写入文件
func WriteFileAppend(name string, data []byte, perm os.FileMode, sync bool) error {
	return WriteFile(name, data, os.O_APPEND|os.O_CREATE|os.O_WRONLY, perm, sync)
}
