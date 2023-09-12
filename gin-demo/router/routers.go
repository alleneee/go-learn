package routers

import (
	"gin-demo/controller"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.Default()
	r.GET("/", controller.GetAll)
	return r
}
