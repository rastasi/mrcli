package config

type Metadatainfo struct {
	ProjectName string
	Version     string
}

func GetMetadata() Metadatainfo {
	return Metadatainfo{
		ProjectName: "mrcli",
		Version:     "1.0.0",
	}
}
