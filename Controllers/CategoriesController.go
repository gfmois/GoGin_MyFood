package Controller

import (
	"gx_myfood/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var categories []Models.Category

	if err := Models.GetCategories(&categories); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	for i := range categories {
		productos := []string{}
		for j := range categories[i].Productos {
			productos = append(productos, categories[i].Productos[j].Id_producto)
		}
		categories[i].C_productos = productos
	}

	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}
