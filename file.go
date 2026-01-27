package fun

import (
	"os"
	"path/filepath"
)

// Mkdir 创建一个目录，如果目录已存在则忽略
func Mkdir(dir string, perm os.FileMode) error {
	err := os.Mkdir(dir, perm)
	if err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}

// MkdirAll 创建任何必要的父目录
func MkdirAll(dir string, perm os.FileMode) error {
	return os.MkdirAll(dir, perm)
}

// WriteFile 写入文件，如果目录不存在回自动创建目录
func WriteFile(name string, data []byte, flag int, perm os.FileMode, sync bool) error {
	// 自动创建目录
	if err := os.MkdirAll(filepath.Dir(name), 0755); err != nil {
		return err
	}

	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return err
	}

	_, err = f.Write(data)

	if sync {
		if syncErr := f.Sync(); syncErr != nil && err == nil {
			err = syncErr
		}
	}

	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}

	return err
}

// WriteFileAppend 追加写入文件，如果目录不存在回自动创建目录
func WriteFileAppend(name string, data []byte, perm os.FileMode, sync bool) error {
	return WriteFile(name, data, os.O_APPEND|os.O_CREATE|os.O_WRONLY, perm, sync)
}

// WriteFileDefault 使用默认参数写入文件，如果目录不存在回自动创建目录
func WriteFileDefault(name string, data []byte) error {
	return WriteFile(name, data, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644, false)
}

// WriteFileDefaultSync 使用默认参数写入文件并同步，如果目录不存在回自动创建目录
func WriteFileDefaultSync(name string, data []byte) error {
	return WriteFile(name, data, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644, true)
}
