package handler

import (
	"PropertyProbe/platform/search"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchGet(searches search.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := searches.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
