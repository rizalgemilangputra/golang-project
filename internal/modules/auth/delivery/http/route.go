package http

import "github.com/gin-gonic/gin"

func NewRoute(app *gin.Engine) *gin.Engine {
	app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return app
}
