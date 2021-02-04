package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"taylors/global"
	"taylors/initialize"
	"taylors/logger"
	"taylors/middleware"
	"taylors/router"
	"time"
	//"runtime"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	initialize.Init()
	defer initialize.Close()

	Router := Routers()
	Router.Static("/form-generator", "./resource/page")
	Router.Static("/index", "./resource/web")
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := &http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logger.Debug("server run success on ", address)
	logger.Error(s.ListenAndServe())
}

func Routers() *gin.Engine {
	var Router = gin.Default()
	//Router.Use(middleware.LoadTls())  // 打开就能玩https了
	logger.Debug("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors())
	logger.Debug("use middleware cors")
	//Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	logger.Debug("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	ApiGroup := Router.Group("")
	router.InitUserRouter(ApiGroup)                  // 注册用户路由
	router.InitBaseRouter(ApiGroup)                  // 注册基础功能路由 不做鉴权
	router.InitMenuRouter(ApiGroup)                  // 注册menu路由
	router.InitAuthorityRouter(ApiGroup)             // 注册角色路由
	router.InitApiRouter(ApiGroup)                   // 注册功能api路由
	router.InitFileUploadAndDownloadRouter(ApiGroup) // 文件上传下载功能路由
	router.InitWorkflowRouter(ApiGroup)              // 工作流相关路由
	router.InitCasbinRouter(ApiGroup)                // 权限相关路由
	router.InitJwtRouter(ApiGroup)                   // jwt相关路由
	router.InitSystemRouter(ApiGroup)                // system相关路由
	router.InitCustomerRouter(ApiGroup)              // 客户路由
	router.InitAutoCodeRouter(ApiGroup)              // 创建自动化代码
	router.InitStockRouter(ApiGroup)                 // 股票
	logger.Info("router register success")
	return Router
}
