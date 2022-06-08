package controller

import "github.com/gin-gonic/gin"

func Route(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
}
