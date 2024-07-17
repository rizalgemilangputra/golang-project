package route

import "github.com/gin-gonic/gin"

type RouteConfig struct {
	App *gin.Engine
}

func (config *RouteConfig) Setup() {
	config.SetupGuestRoute()
}

func (config *RouteConfig) SetupGuestRoute() {
	config.App.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
