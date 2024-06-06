package router

import (
	"github.com/gin-gonic/gin"
	"github.com/riandiangga12/Golangtraining/Training Golang/handler"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/", handler.RootHandler)
	r.POST("/", handler.PostHandler)
	

}
