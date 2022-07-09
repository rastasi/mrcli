package utils

import (
	"mrcli/lib/logger"
	"mrcli/lib/metadata"
)

func Header() {
	info := metadata.GetMetadata()
	logger.Log(info.Name, "Version: %s\n\n", info.Version)
}
