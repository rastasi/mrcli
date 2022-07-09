package utils

import (
	"mrcli/lib/config"
	"mrcli/lib/logger"
)

func Header() {
	info := config.GetMetadata()
	logger.Log("", "%s %s\n\n", info.Name, info.Version)
}
