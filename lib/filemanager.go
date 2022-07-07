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

func FileExistsDisplay(path string) {
	if FileExists(path) {
		LogSuccess(path, "OK\n", []any{})
	} else {
		LogFail(path, "Missing\n", []any{})
	}
}

func FileExistsDisplayBulk(paths []string) {
	for _, path := range paths {
		FileExistsDisplay(path)
	}
}

func CreateDir(dir string) {
	if !FileExists(dir) {
		permission := 0766
		LogSuccess(dir, "Create directory\n", []any{})
		os.Mkdir(dir, fs.FileMode(permission))
	} else {
		LogFail(dir, "Directory already exists\n", []any{})
	}
}

func CreateDirBulk(dirs []string) {
	for _, dir := range dirs {
		CreateDir(dir)
	}
}

func CreateFile(path string) {
	if !FileExists(path) {
		LogSuccess(path, "Create file\n", []any{})
		os.Create(path)
	} else {
		LogFail(path, "File already exists\n", []any{})
	}
}

func CreateFileBulk(paths []string) {
	for _, path := range paths {
		CreateFile(path)
	}
}
