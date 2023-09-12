package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, "hello")
}
