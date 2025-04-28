package routes

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app

	route.Use(gzip.Gzip(gzip.DefaultCompression))

	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"ping": "pong"})
	})
}
