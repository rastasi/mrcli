package utils

type metadata struct {
	ProjectName string
	Version     string
}

func GetMetadata() metadata {
	return metadata{
		ProjectName: "mrcli",
		Version:     "1.0.0",
	}
}
