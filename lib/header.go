package lib

func Header() {
	metadata := GetMetadata()
	Log("", "%s %s\n\n", []any{metadata.ProjectName, metadata.Version})
}
