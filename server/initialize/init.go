package initialize

import (
	"taylors/config"
	"taylors/global"
)

func Init() {
	err := config.Init()
	if err != nil {
		panic("配置错误")
	}

	Log()
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		Mysql()
	}
	DBTables()

	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		Redis()
	}

	Cron()
	Crawler()
}

// 程序结束前
func Close() {
	_ = global.GVA_DB.Close()
}
