package routers

import (
	handler "demo/handler/user"
	mw "demo/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()
	//测试接口
	commonTest := router.Group("/common")
	{
		commonTest.GET("/test", mw.Logger(), handler.Test)

		commonTest.GET("/user", handler.Get)
		commonTest.POST("/user", handler.Add)
		commonTest.PUT("/user", handler.UpdateData)
		commonTest.PATCH("/user", handler.PatchData)

		commonTest.DELETE("/user", handler.DelData)

		commonTest.OPTIONS("/user", handler.OptionData)

	}

	return router

}
