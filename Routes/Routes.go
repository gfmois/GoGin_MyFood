package Routes

import (
	Controller "gx_myfood/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		reservesGrp := api.Group("/reservas")
		{
			reservesGrp.GET("", Controller.GetReserves)
			reservesGrp.GET(":id_reserva", Controller.GetReserva)
			reservesGrp.GET("/getBannedDays", Controller.GetBannedDays)
			reservesGrp.GET("/image", Controller.GetPdfImage)
			reservesGrp.POST("", Controller.GetBody)
		}

		productsGrp := api.Group("/products")
		{
			productsGrp.GET("", Controller.GetProducts)
			productsGrp.GET(":id_producto", Controller.GetProductById)
			productsGrp.GET("/alergenos", Controller.GetAllergens)
			productsGrp.GET("/alergenos/:id_alergeno", Controller.GetAlergeno)
			// productsGrp.GET("/categories")
			// productsGrp.GET("/categories/:id_category")
		}
	}

	return r
}
