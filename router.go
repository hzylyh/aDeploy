/*
 * @Description: 路由文件
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2021-04-13 18:40:28
 * @LastEditors: John Holl
 * @LastEditTime: 2021-05-10 17:06:50
 */
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

	// 跨域配置
	apiRoot.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
	}))

	// deployment
	deployApi := apiRoot.Group("/deployment")
	deployApi.POST("/create", controller.CreateDeployment)

	// service
	serviceApi := apiRoot.Group("/service")
	serviceApi.POST("/create", controller.CreateService)

	// image
	imageApi := apiRoot.Group("/image")
	imageApi.POST("/getList", controller.GetImageList)

	return router
}
