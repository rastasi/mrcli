package metadata

type metadatainfo struct {
	ProjectName string
	Version     string
}

func GetMetadata() metadatainfo {
	return metadatainfo{
		ProjectName: "mrcli",
		Version:     "1.0.0",
	}
}
