package routers

import (
	"EasyGinLite/middleware"
	v "EasyGinLite/version/v1"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouter(mode string) *gin.Engine {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic err: ", err)
		}
	}()
	gin.SetMode(mode)
	router := gin.Default()
	router.Use(middleware.Cors())
	//这里可以考虑是否是直接内容渲染=》页面直接渲染--可以说比较老的方式了
	router.GET("/", v.IndexHtml)

	//这里就是一个接口，GET /api/v1/index 可以直接调用，鉴权的东西我没有加，后期看情况吧
	api := router.Group("/api")
	v1 := api.Group("/v1")
	{
		v1.GET("/index", v.Index)
	}

	return router
}
