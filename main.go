package main

import (
	conf "EasyGinLite/config"
	routers "EasyGinLite/router"
	"EasyGinLite/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化设置：./config文件夹下
	//conf.ini 就是存放各种数据连接的信息
	conf.InitConfig()

	//初始化全局变量
	utils.Global()

	//调试模式 gin.DebugMode 有的没的GIN的那些调试信息会显示
	//发布模式 gin.ReleaseMode 以上信息就没了
	router := routers.InitRouter(gin.DebugMode)

	//这里就是就是定义前端的html资源，我用了通配符
	router.LoadHTMLGlob("./dist/*.html")

	//这里定义的是静态资源文件，也就是url中的/assets映射我本地的文件夹目录./dist/assets
	router.Static("/assets", "./dist/assets")

	router.Run("0.0.0.0:6666")
}
