package main

import (
	"PropertyProbe/httpd/handler"
	"PropertyProbe/platform/search"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB TESTING
var dsn = "go:Gators123@tcp(cen3031-project.mysql.database.azure.com:3306)/listings?charset=utf8mb4&parseTime=True&loc=Local"
var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

type Property struct {
	Owner   string
	Address string
}

func GoDatabaseCreate(w http.ResponseWriter, r *http.Request) {
	Property := Property{
		Owner:   "Julien",
		Address: "1013 Fieldstone Dr"}

	db.Create(&Property)
	if err := db.Create(&Property).Error; err != nil {
		log.Fatalln((err))
	}

	json.NewEncoder(w).Encode(Property)

	fmt.Println("Fields Added", Property)
}

func main() {
	// DB TESTING
	// http.HandleFunc("/testDB", GoDatabaseCreate)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	searchHistory := search.New()

	r := gin.Default()

	api := r.Group("/api")
	{
		// api.GET("/ping", handler.PingGet())
		api.GET("/search", handler.SearchGet(searchHistory))
		api.POST("/search", handler.SearchPost(searchHistory))
	}

	r.Run("0.0.0.0:5000")
}
