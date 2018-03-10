package util

import (
	"os"
	"path/filepath"
	"runtime"
)

func DataFile(filename string) (*os.File, error) {
	_, source_file, _, _ := runtime.Caller(1)
	path := filepath.Join(filepath.Dir(source_file), filename)
	return os.Open(path)
}
