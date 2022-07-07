package lib

func Header() {
	metadata := GetMetadata()
	Log("", "%s %s\n\n", metadata.ProjectName, metadata.Version)
}
