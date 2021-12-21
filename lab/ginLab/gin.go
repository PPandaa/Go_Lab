package ginLab

import (
	"GoLab/lab/ginLab/handler"

	"github.com/gin-gonic/gin"
)

func Server() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(JSONMiddleware())

	// router
	r.GET("/", handler.Hello)
	r.GET("/showParameter/:user", handler.ShowParameter)
	r.GET("/showQuery", handler.ShowQuery)
	r.POST("/showBody", handler.ShowBody)

	r.Run(":8080")

}

func JSONMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}

}
