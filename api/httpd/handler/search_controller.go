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

type SearchPostRequest struct {
	SearchTerm string `json:"search_term"`
}

func SearchPost(searches search.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := SearchPostRequest{}
		c.Bind(&requestBody)

		searchterm := search.Search{
			SearchTerm: requestBody.SearchTerm,
		}

		searches.Add(searchterm)

		c.Status(http.StatusNoContent)
	}
}
