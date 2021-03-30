package main

import (
	"aDeploy/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

/**
路由控制
*/
func MapRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
	}))

	router.LoadHTMLGlob("web/dist/*.html")          // 添加入口index.html
	router.LoadHTMLFiles("web/static/*/*")          // 添加资源路径
	router.Static("/static", "./web/dist/static")   // 添加资源路径
	router.StaticFile("/", "./web/dist/index.html") //前端接口

	// 创建根路由
	apiRoot := router.Group("/aDeployApi")
	apiRoot.Use()

	authApi := apiRoot.Group("/image")
	authApi.POST("/getList", controller.GetList)
	return router
}
