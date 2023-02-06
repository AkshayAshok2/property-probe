package main

import (
	"PropertyProbe/httpd/handler"
	"PropertyProbe/platform/properties"

	"github.com/gin-gonic/gin"
)

func main() {

	//DB placeholder
	housing := properties.New()

	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/ping", handler.PingGet())
		api.GET("/properties", handler.PropertiesGet(housing))
		api.POST("/properties", handler.PropertiesPost(housing))
	}

	r.Run("0.0.0.0:5000")
}
