package Controller

import (
	"net/http"

	"gx_myfood/Models"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []Models.Product

	if err := Models.GetProducts(&products); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H{
		"productos": products,
	})
}

func GetProductById(c *gin.Context) {
	var product Models.Product
	id_product := c.Param("id_product")

	if len(id_product) != 0 {
		if err := Models.GetProductById(&product, id_product); err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"producto": product,
		})
		return
	}

	c.AbortWithStatus(http.StatusBadRequest)
	return
}

func GetAllergens(c *gin.Context) {
	var allergens []Models.Allergen

	if err := Models.GetAllergens(&allergens); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H{
		"allergenos": allergens,
	})
}

func GetAlergeno(c *gin.Context) {
	var allergen Models.Allergen
	id_allergen := c.Param("id_allergen")

	if err := Models.GetAllergeno(&allergen, id_allergen); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H{
		"alergeno": allergen,
	})
}
