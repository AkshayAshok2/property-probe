package handler

import (
	"PropertyProbe/platform/properties"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PropertiesPostRequest struct {
	gorm.Model
	AuctionType     string `json:"auction_type"`
	JudgementAmount string `json:"judgement_amount"`
	Address         string `json:"property_address"`
	AssessedValue   string `json:"assessedvalue"`
}

func PropertiesPost(housing properties.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := PropertiesPostRequest{}
		c.Bind(&requestBody)

		property := properties.Property{
			AuctionType:     requestBody.AuctionType,
			JudgementAmount: requestBody.JudgementAmount,
			Address:         requestBody.Address,
			AssessedValue:   requestBody.AssessedValue,
		}

		housing.Add(property)

		c.Status(http.StatusNoContent)
	}
}
