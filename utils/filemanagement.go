package utils

import (
	"fmt"
	"io/fs"
	"os"
)

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func CreateDir(path string) {
	if !exists(path) {
		permission := 0766
		fmt.Println("Create directory: " + path)
		os.Mkdir(path, fs.FileMode(permission))
	} else {
		fmt.Println("Directory already exists: " + path)
	}
}
