package utils

import (
	"mrcli/lib/config"
	"mrcli/lib/logger"
)

func Header() {
	metadata := config.GetMetadata()
	logger.Log("", "%s %s\n\n", metadata.Name, metadata.Version)
}
