package main

import (
	"awesomeProject/controller"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	controller.Route(r)
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
