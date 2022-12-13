package Routes

import (
	Controller "gx_myfood/Controllers"

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
			reservesGrp.GET("", Controller.GetReserves)
			reservesGrp.GET("/getBannedDays", Controller.GetBannedDays)
			reservesGrp.GET("/image", Controller.GetPdfImage)
			reservesGrp.GET(":id_reserva", Controller.GetReserve)
			reservesGrp.POST("", Controller.CreateReserve)
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
