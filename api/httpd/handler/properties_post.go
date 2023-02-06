package handler

import (
	"PropertyProbe/platform/properties"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PropertiesPostRequest struct {
	Owner   string `json:"owner"`
	Address string `json:"address"`
}

func PropertiesPost(housing properties.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := PropertiesPostRequest{}
		c.Bind(&requestBody)

		property := properties.Property{
			Owner:   requestBody.Owner,
			Address: requestBody.Address,
		}

		housing.Add(property)

		c.Status(http.StatusNoContent)
	}
}
