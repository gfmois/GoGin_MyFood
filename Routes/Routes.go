package Routes

import (
	Controller "gx_myfood/Controllers"
	"gx_myfood/Middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	MakeRoutes(r)
	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		reservesGrp := api.Group("/reservas")
		{
			reservesGrp.GET("/image", Controller.GetImage)
			reservesGrp.Use(Middlewares.AuthMiddleware(true)).GET("", Controller.GetReserves)
			reservesGrp.GET("/getBannedDays", Controller.GetBannedDays)
			reservesGrp.Use(Middlewares.AuthMiddleware(true)).GET(":id_reserva", Controller.GetReserve)
			reservesGrp.Use(Middlewares.AuthMiddleware(true)).GET("/pdf/:id_reserva", Controller.GetPDFReserve)
			reservesGrp.Use(Middlewares.AuthMiddleware(true)).POST("", Controller.CreateReserve)
		}

		productsGrp := api.Group("/productos")
		{
			productsGrp.GET("", Controller.GetProducts)
			productsGrp.GET("/filtro", Controller.GetFilteredProducts)
			productsGrp.GET("/search/:producto", Controller.SearchProducts)
			productsGrp.GET(":slug_producto", Controller.GetProductDetails)
			productsGrp.GET("/alergenos", Controller.GetAllergens)
			productsGrp.GET("/alergenos/:id_alergeno", Controller.GetAlergen)
		}

		categoriesGrp := api.Group("/categorias")
		{
			categoriesGrp.GET("", Controller.GetCategories)
		}

		testGrp := api.Group("/test").Use(Middlewares.AuthMiddleware(true))
		{
			testGrp.GET("", func(c *gin.Context) {
				client_model, _ := c.Get("client_model")
				c.JSON(200, client_model)
			})
		}

		clientsGrp := api.Group("/client").Use(Middlewares.AuthMiddleware(true))
		{
			clientsGrp.GET("/profile", Controller.GetProfile)
			clientsGrp.PUT("/profile", Controller.UpdateProfile)
		}

		authGrp := api.Group("/auth")
		{
			authGrp.POST("/register", Controller.Register)
			authGrp.POST("/login", Controller.Login)
		}

		ordersGrp := api.Group("/pedidos").Use(Middlewares.AuthMiddleware(true))
		{
			ordersGrp.GET("", Controller.GetOrders)
			ordersGrp.POST("", Controller.AddOrder)

		}
	}

	return r
}

func MakeRoutes(r *gin.Engine) {
	cors := func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}
		c.Next()
	}
	r.Use(cors)
}
