package fun

import "os"

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
