package Controller

import (
	"math"
	"net/http"
	"strings"

	"gx_myfood/Models"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []Models.Product

	if err := Models.GetProducts(&products); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	for i := range products {
		categorias := []string{}
		alergenos := []string{}
		for j := range products[i].Categorias {
			categorias = append(categorias, products[i].Categorias[j].Id_categoria)
		}
		for k := range products[i].Alergenos {
			alergenos = append(alergenos, products[i].Alergenos[k].Id_alergeno)
		}
		products[i].C_categorias = categorias
		products[i].C_alergenos = alergenos
	}

	c.JSON(http.StatusOK, gin.H{
		"productos": products,
	})
}

func GetFilteredProducts(c *gin.Context) {
	var products []Models.Product
	var count_products []Models.Product
	var categoriaValores []string
	categoria := c.Query("categorias")
	if categoria != "" {
		categoriaValores = strings.Split(categoria, ",")
	}
	orden := c.Query("orden")
	rango := c.Query("rango")
	rangoValores := strings.Split(rango, ",")
	paginacion := c.Query("paginacion")
	if err := Models.GetFilteredProducts(&products, categoriaValores, orden, rangoValores[0], rangoValores[1], paginacion); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err := Models.GetFilteredProducts(&count_products, categoriaValores, orden, rangoValores[0], rangoValores[1], ""); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	for i := range products {
		categorias := []string{}
		alergenos := []string{}
		for j := range products[i].Categorias {
			categorias = append(categorias, products[i].Categorias[j].Id_categoria)
		}
		for k := range products[i].Alergenos {
			alergenos = append(alergenos, products[i].Alergenos[k].Id_alergeno)
		}
		products[i].C_categorias = categorias
		products[i].C_alergenos = alergenos
	}
	numPages := int(math.Ceil(float64(len(count_products)) / float64(12)))
	c.JSON(http.StatusOK, gin.H{
		"productos": products,
		"pages":     numPages,
	})
}

func SearchProducts(c *gin.Context) {
	var products []Models.Product
	producto := c.Param("producto")

	if err := Models.SearchProducts(&products, producto); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	for i := range products {
		categorias := []string{}
		alergenos := []string{}
		for j := range products[i].Categorias {
			categorias = append(categorias, products[i].Categorias[j].Id_categoria)
		}
		for k := range products[i].Alergenos {
			alergenos = append(alergenos, products[i].Alergenos[k].Id_alergeno)
		}
		products[i].C_categorias = categorias
		products[i].C_alergenos = alergenos
	}

	c.JSON(http.StatusOK, products)
}

func GetProductDetails(c *gin.Context) {
	var product Models.Product
	slug_producto := c.Param("slug_producto")

	if len(slug_producto) != 0 {
		if err := Models.GetProductBySlug(&product, slug_producto); err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		categorias := []string{}
		alergenos := []string{}
		for j := range product.Categorias {
			categorias = append(categorias, product.Categorias[j].Id_categoria)
		}
		for k := range product.Alergenos {
			alergenos = append(alergenos, product.Alergenos[k].Id_alergeno)
		}
		product.C_categorias = categorias
		product.C_alergenos = alergenos

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
		return
	}

	for i := range allergens {
		productos := []string{}
		for j := range allergens[i].Productos {
			productos = append(productos, allergens[i].Productos[j].Id_producto)
		}
		allergens[i].C_productos = productos
	}

	c.JSON(http.StatusOK, gin.H{
		"allergens": allergens,
	})
}

func GetAlergen(c *gin.Context) {
	var allergen Models.Allergen
	id_allergen := c.Param("id_alergeno")

	if err := Models.GetAllergen(&allergen, id_allergen); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	productos := []string{}
	for i := range allergen.Productos {
		productos = append(productos, allergen.Productos[i].Id_producto)
	}
	allergen.C_productos = productos
	c.JSON(http.StatusOK, gin.H{
		"alergeno": allergen,
	})
}
