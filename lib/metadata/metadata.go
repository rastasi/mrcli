package metadata

const ENTITY_TYPE_FILE = "FILE"
const ENTITY_TYPE_DIR = "DIRECTORY"

type StructureEntity struct {
	EntityType string
	Pattern    string
	Template   string
}

type applicationMetainfo struct {
	Structure []StructureEntity
}

type projectMetainfo struct {
	Structure []StructureEntity
}

type libraryMetainfo struct {
	Structure []StructureEntity
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
			Structure: []StructureEntity{
				{
					EntityType: ENTITY_TYPE_DIR,
					Pattern:    "./%s",
				},
				{
					EntityType: ENTITY_TYPE_DIR,
					Pattern:    "./%s/apps",
				},
				{
					EntityType: ENTITY_TYPE_DIR,
					Pattern:    "./%s/libs",
				},
				{
					EntityType: ENTITY_TYPE_DIR,
					Pattern:    "./%s/dists",
				},
				{
					EntityType: ENTITY_TYPE_FILE,
					Pattern:    "./%s/apps/.keep",
				},
				{
					EntityType: ENTITY_TYPE_FILE,
					Pattern:    "./%s/libs/.keep",
				},
				{
					EntityType: ENTITY_TYPE_FILE,
					Pattern:    "./%s/dists/.keep",
				},
			},
		},
		Application: applicationMetainfo{
			Structure: []StructureEntity{
				{
					EntityType: ENTITY_TYPE_DIR,
					Pattern:    "./apps/%s",
				},
				{
					EntityType: ENTITY_TYPE_DIR,
					Pattern:    "./dists/%s",
				},
				{
					EntityType: ENTITY_TYPE_FILE,
					Pattern:    "./apps/%s/main.go",
				},
				{
					EntityType: ENTITY_TYPE_FILE,
					Pattern:    "./dists/%s/.keep",
				},
			},
		},
		Library: libraryMetainfo{
			Structure: []StructureEntity{
				{
					EntityType: ENTITY_TYPE_DIR,
					Pattern:    "./libs/%s",
				},
				{
					EntityType: ENTITY_TYPE_FILE,
					Pattern:    "./libs/%s/main.go",
				},
			},
		},
	}
}
