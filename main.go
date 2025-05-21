package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonItem map[string]any
type Items map[string]JsonItem

var items = make(Items, 0)

func main() {
	router := gin.Default()
	router.GET("/", getItems)
	router.GET("/:key", getByKey)
	router.PUT("/:key", putItem)
	router.PATCH("/:key", patchItem)

	router.Run("localhost:8080")
}

func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}

func getByKey(c *gin.Context) {
	key := c.Param("key")

	c.IndentedJSON(http.StatusOK, items[key])
}

func putItem(c *gin.Context) {
	key := c.Param("key")

	var newItem JsonItem

	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	items[key] = newItem
	c.IndentedJSON(http.StatusCreated, newItem)
}

func patchItem(c *gin.Context) {
	key := c.Param("key")

	var patchItem JsonItem

	if err := c.BindJSON(&patchItem); err != nil {
		return
	}

	item := items[key]

	for k, v := range patchItem {
		item[k] = v
	}

	c.IndentedJSON(http.StatusCreated, item)
}
