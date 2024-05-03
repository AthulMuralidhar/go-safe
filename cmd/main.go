package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// item represents data about a record item.
type Item struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

func main() {
	router := gin.Default()
	router.GET("/items", getItems)
	router.GET("/items/:id", getAlbumByID)
	router.POST("/items", postItems)

	router.Run("localhost:8080")
}

// getItems responds with the list of all items as JSON.
func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}

// postItems adds an item from JSON received in the request body.
func postItems(c *gin.Context) {
	var newItem Item

	// Call BindJSON to bind the received JSON to
	// newItem.
	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	fmt.Println(newItem)

	// Add the new item to the slice.
	items = append(items, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
}

// getAlbumByID locates the item whose ID value matches the id
// parameter sent by the client, then returns that item as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of items, looking for
	// an item whose ID value matches the parameter.
	for _, a := range items {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}
