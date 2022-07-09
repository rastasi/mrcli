package utils

import (
	"mrcli/lib/config"
	"mrcli/lib/logger"
)

func Header() {
	info := config.GetMetadata()
	logger.Log(info.Name, "Version: %s\n\n", info.Version)
}
