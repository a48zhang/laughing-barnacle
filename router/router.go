package router

import (
	"github.com/gin-gonic/gin"
	"main/handler/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("api/v1")
	v1.Use(middleware.Validator())
	return r
}
