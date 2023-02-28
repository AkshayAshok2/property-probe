// package handler

// import (
// 	"PropertyProbe/platform/properties"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func PropertiesGet(housing properties.Getter) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		results := housing.GetAll()
// 		c.JSON(http.StatusOK, results)
// 	}
// }