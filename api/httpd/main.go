package main

import (
	"PropertyProbe/database"
	"PropertyProbe/httpd/handler"
	"PropertyProbe/platform/properties"
	"PropertyProbe/platform/search"

	"github.com/gin-gonic/gin"
)

func main() {

	searchHistory := search.New()
	propertyRepo := handler.New()

	// Connect to the database
	db := database.InitDb()
	if db == nil {
		panic("Failed to connect to the database")
	}

	// Clear the properties table
	err := database.ClearDB()
	if err != nil {
		panic(err)
	}

	properties.AddPropertiesToDatbase()

	r := gin.Default()

	api := r.Group("/api")
	{
		// api.GET("/ping", handler.PingGet())
		api.GET("/search", handler.SearchGet(searchHistory))
		api.POST("/search", handler.SearchPost(searchHistory))
		api.POST("/properties", propertyRepo.CreateProperty)
		api.GET("/properties/:zipcode", propertyRepo.GetZipCodeProperties)
		api.GET("/properties", propertyRepo.GetAllProperties)
		api.GET("/properties/address/:address", propertyRepo.GetProperty)
		api.GET("properties/zipcodes", propertyRepo.GetUniqueZipCodes)
		api.PUT("/properties/address/:address", propertyRepo.UpdateProperty)
		api.DELETE("/properties/address/:address", propertyRepo.DeleteProperty)
	}

	r.Run("0.0.0.0:5000")
}
