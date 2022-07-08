package config

type Metadatainfo struct {
	Name    string
	Version string
}

func GetMetadata() Metadatainfo {
	return Metadatainfo{
		Name:    "mrcli",
		Version: "1.0.0",
	}
}
