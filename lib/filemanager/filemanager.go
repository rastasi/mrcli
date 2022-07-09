package filemanager

import (
	"fmt"
	"io/fs"
	"mrcli/lib/logger"
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
		logger.LogSuccess(path, "OK\n")
	} else {
		logger.LogFail(path, "Missing\n")
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
		logger.LogSuccess(dir, "Create directory\n")
		os.Mkdir(dir, fs.FileMode(permission))
	} else {
		logger.LogFail(dir, "Directory already exists\n")
	}
}

func CreateDirBulk(dirs []string, placeholders ...any) {
	for _, dir := range dirs {
		CreateDir(fmt.Sprintf(dir, placeholders...))
	}
}

func CreateFile(path string) {
	if !FileExists(path) {
		logger.LogSuccess(path, "Create file\n")
		os.Create(path)
	} else {
		logger.LogFail(path, "File already exists\n")
	}
}

func CreateFileBulk(paths []string, placeholders ...any) {
	for _, path := range paths {
		CreateFile(fmt.Sprintf(path, placeholders...))
	}
}
