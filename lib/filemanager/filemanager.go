package filemanager

import (
	"fmt"
	"io/fs"
	"mrcli/lib/logger"
	"os"
	"strings"
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
		logger.LogSuccess(path, "OK")
	} else {
		logger.LogFail(path, "Missing")
	}
}

func FileExistsDisplayBulk(paths []string, placeholders ...any) {
	for _, path := range paths {
		FileExistsDisplay(strings.ReplaceAll(fmt.Sprintf(path, placeholders...), "//", "/"))
	}
}

func CreateDir(dir string, placeholders ...any) {
	dir = fmt.Sprintf(dir, placeholders...)
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
		CreateDir(dir, placeholders...)
	}
}

func CreateFile(path string, placeholders ...any) {
	path = fmt.Sprintf(path, placeholders...)
	if !FileExists(path) {
		logger.LogSuccess(path, "Create file\n")
		os.Create(path)
	} else {
		logger.LogFail(path, "File already exists\n")
	}
}

func CreateFileBulk(paths []string, placeholders ...any) {
	for _, path := range paths {
		CreateFile(path, placeholders...)
	}
}
