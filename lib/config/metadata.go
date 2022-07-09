package config

type applicationMetainfo struct {
	BaseDirectoryStructure []string
	BaseFiles              []string
}

type projectMetainfo struct {
	BaseDirectoryStructure []string
	BaseFiles              []string
}

type libraryMetainfo struct {
	BaseDirectoryStructure []string
	BaseFiles              []string
}

type Metadatainfo struct {
	Name        string
	Version     string
	Project     projectMetainfo
	Application applicationMetainfo
	Library     libraryMetainfo
}

func GetMetadata() Metadatainfo {
	return Metadatainfo{
		Name:    "mrcli",
		Version: "1.0.0",
		Project: projectMetainfo{
			BaseDirectoryStructure: []string{
				"./%s",
				"./%s/apps",
				"./%s/libs",
				"./%s/dists",
			},
			BaseFiles: []string{
				"./%s/apps/.keep",
				"./%s/libs/.keep",
				"./%s/dists/.keep",
			},
		},
		Application: applicationMetainfo{
			BaseDirectoryStructure: []string{
				"./apps/%s",
				"./dists/%s",
			},
			BaseFiles: []string{
				"./apps/%s/.keep",
				"./dists/%s./keep",
			},
		},
		Library: libraryMetainfo{
			BaseDirectoryStructure: []string{
				"./libs/%s",
			},
			BaseFiles: []string{
				"./libs/%s/.keep",
			},
		},
	}
}
