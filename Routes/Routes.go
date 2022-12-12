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
			reservesGrp.GET("/getBannedDays", Controller.GetBannedDays)
		}
	}

	return r
}
