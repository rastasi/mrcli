package layout

import (
	"mrcli/metadata"
)

func Header() {
	metadata := metadata.GetMetadata()
	Log("", "%s %s\n\n", []any{metadata.ProjectName, metadata.Version})
}
