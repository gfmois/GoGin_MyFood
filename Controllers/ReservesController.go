package Controller

import (
	"net/http"

	"gx_myfood/Models"

	"github.com/gin-gonic/gin"
)

func GetReserves(c *gin.Context) {
	var reserves []Models.Reserve
	err := Models.GetReserves(&reserves)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H{
		"reserves": reserves,
	})
}

func GetBannedDays(c *gin.Context) {
	var bannedDays []string
	comensales, servicio := c.Query("comensales"), c.Query("servicio")

	if len(comensales) != 0 || len(servicio) != 0 {
		err := Models.GetBannedDays(&bannedDays, comensales, servicio)

		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, bannedDays)
		return
	}

	c.AbortWithStatus(http.StatusBadRequest)
	return
}
