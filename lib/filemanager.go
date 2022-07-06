package lib

import (
	"io/fs"
	"os"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func CreateDir(name string) {
	if !FileExists(name) {
		permission := 0766
		LogSuccess(name, "Create directory\n", []any{})
		os.Mkdir(name, fs.FileMode(permission))
	} else {
		LogFail(name, "Directory already exists\n", []any{})
	}
}
