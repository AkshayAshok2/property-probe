package handler

import (
	"PropertyProbe/platform/properties"
	"errors"
	"gorm-test/database"
	"gorm-test/models"
	"net/http"
	"strconv"

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

func (repository *PropertyRepo) GetProperties(c *gin.Context) {
	var property []properties.Property
	err := properties.GetProperties(repository.Db, &property)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, property)
}

func (repository *PropertyRepo) GetProperty(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var property properties.Property
	err := properties.GetProperty(repository.Db, &property, id)
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
	id, _ := strconv.Atoi(c.Param("id"))
	err := properties.GetProperty(repository.Db, &property, id)
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
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteProperty(repository.Db, &property, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
