package handler

import (
	"PropertyProbe/platform/search"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSearchGet(t *testing.T) {
	// Create a new search repository and add some search terms
	repo := search.New()
	searchTerm1 := search.Search{SearchTerm: "test 1"}
	searchTerm2 := search.Search{SearchTerm: "test 2"}
	repo.Add(searchTerm1)
	repo.Add(searchTerm2)

	// Create a new Gin router and add the SearchGet handler
	router := gin.Default()
	router.GET("/search", SearchGet(repo))

	// Create a new HTTP request for the /search endpoint
	req, err := http.NewRequest("GET", "/search", nil)
	assert.NoError(t, err)

	// Create a new HTTP response recorder
	rr := httptest.NewRecorder()

	// Send the HTTP request to the server
	router.ServeHTTP(rr, req)

	// Check that the HTTP status code is OK
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check that the HTTP response body contains the expected search terms
	expected := []search.Search{searchTerm1, searchTerm2}
	var actual []search.Search
	err = json.NewDecoder(rr.Body).Decode(&actual)
	assert.NoError(t, err)
	assert.Equal(t, len(expected), len(actual))
	for i, v := range actual {
		assert.Equal(t, expected[i].SearchTerm, v.SearchTerm)
	}
}

func TestSearchPost(t *testing.T) {
	// Create a new search repository
	repo := search.New()

	// Create a new Gin router and add the SearchPost handler
	router := gin.Default()
	router.POST("/search", SearchPost(repo))

	// Create a new HTTP request for the /search endpoint
	payload := `{"search_term": "test"}`
	req, err := http.NewRequest("POST", "/search", strings.NewReader(payload))
	assert.NoError(t, err)

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP response recorder
	rr := httptest.NewRecorder()

	// Send the HTTP request to the server
	router.ServeHTTP(rr, req)

	// Check that the HTTP status code is NoContent
	assert.Equal(t, http.StatusNoContent, rr.Code)

	// Check that the search term was added to the repository
	results := repo.GetAll()
	assert.Equal(t, 1, len(results))
	assert.Equal(t, "test", results[0].SearchTerm)
}
