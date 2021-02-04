package initialize

import (
	"taylors/global"
	"taylors/logger"
)

func Log() {
	is_dev := false
	if global.GVA_CONFIG.Log.Environment == "dev" {
		is_dev = true
	}

	logger.Init(global.GVA_CONFIG.Log.FileName, is_dev, global.GVA_CONFIG.Log.Level)
}
