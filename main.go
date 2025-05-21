package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonItem map[string]any

var item JsonItem

func main() {
	router := gin.Default()
	router.GET("/items", getItems)
	router.PUT("/items", putItem)
	router.PATCH("/items", patchItem)

	router.Run("localhost:8080")
}

func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, item)
}

func putItem(c *gin.Context) {
	var newItem JsonItem

	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	item = newItem
	c.IndentedJSON(http.StatusCreated, newItem)
}

func patchItem(c *gin.Context) {
	var patchItem JsonItem

	if err := c.BindJSON(&patchItem); err != nil {
		return
	}

	for k, v := range patchItem {
		item[k] = v
	}

	c.IndentedJSON(http.StatusCreated, item)
}
