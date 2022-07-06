package utils

import "fmt"

func Header() {
	metadata := GetMetadata()
	fmt.Printf("%s %s\n\n", metadata.ProjectName, metadata.Version)
}
