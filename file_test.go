package fun

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMkdir(t *testing.T) {
	tempDir := os.TempDir()
	dir := filepath.Join(tempDir, "test")
	assert.Equal(t, nil, Mkdir(dir, 0755))
	assert.Equal(t, nil, MkdirAll(dir, 0755))
}

func TestWriteFileAppend(t *testing.T) {
	tempDir := os.TempDir()
	dir := filepath.Join(tempDir, "test")

	_ = Mkdir(dir, 0755)

	filename := filepath.Join(dir, "test.txt")

	data := []byte("All the data \nI wish to write to a file")

	t.Log(filename)

	for i := 0; i < 3; i++ {
		_ = WriteFileAppend(filename, data, 0777, false)
		_ = WriteFileAppend(filename, []byte("\n"), 0777, false)
	}
}
