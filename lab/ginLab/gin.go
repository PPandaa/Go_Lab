package ginLab

import "github.com/gin-gonic/gin"

func Server() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", hello)
	r.GET("/showParameter/:user", showParameter)
	r.GET("/showQuery", showQuery)
	r.POST("/showBody", showBody)

	r.Run(":8080")

}
