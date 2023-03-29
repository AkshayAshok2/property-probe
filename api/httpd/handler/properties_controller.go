package handler

import (
	"PropertyProbe/database"
	"PropertyProbe/platform/properties"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PropertyRepo struct {
	Db *gorm.DB
}

func New() *PropertyRepo {
	db := database.InitDb()
	db.AutoMigrate(&properties.Property{})
	return &PropertyRepo{Db: db}
}

func (repository *PropertyRepo) CreateProperty(c *gin.Context) {
	var property properties.Property
	c.BindJSON(&property)
	err := properties.CreateProperty(repository.Db, &property)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, property)
}

func (repository *PropertyRepo) GetAllProperties(c *gin.Context) {
	var property []properties.Property
	err := properties.GetAllProperties(repository.Db, &property)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, property)
}

func (repository *PropertyRepo) GetZipCodeProperties(c *gin.Context) {
	var property []properties.Property
	zipcode := (c.Param("zipcode"))
	err := properties.GetZipCodeProperties(repository.Db, &property, zipcode)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, property)
}

func (repository *PropertyRepo) GetProperty(c *gin.Context) {
	address := (c.Param("address"))
	var property properties.Property
	err := properties.GetProperty(repository.Db, &property, address)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, property)
}

func (repository *PropertyRepo) UpdateProperty(c *gin.Context) {
	var property properties.Property
	address := c.Param("address")
	err := properties.GetProperty(repository.Db, &property, address)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&property)
	err = properties.UpdateProperty(repository.Db, &property)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, property)
}

// delete user
func (repository *PropertyRepo) DeleteProperty(c *gin.Context) {
	var property properties.Property
	address := c.Param("address")
	err := properties.DeleteProperty(repository.Db, &property, address)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Property deleted successfully"})
}
