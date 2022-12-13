package Controller

import (
	"fmt"
	"gx_myfood/Common"
	"gx_myfood/Models"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetReserves(c *gin.Context) {
	var reserves []Models.Reserve

	if err := Models.GetReserves(&reserves); err != nil {
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

		if err := Models.GetBannedDays(&bannedDays, comensales, servicio); err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, bannedDays)
		return
	}

	c.AbortWithStatus(http.StatusBadRequest)
	return
}

func GetPdfImage(c *gin.Context) {
	buff, err := ioutil.ReadFile("Assets/Images/test.jpg")

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.Header("Content-Type", "image/jpeg")
	c.Writer.Write(buff)
}

func GetReserve(c *gin.Context) {
	var reserva Models.Reserve
	id_reserva := c.Param("id_reserva")

	if err := Models.GetReserve(&reserva, id_reserva); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H{
		"reserva": reserva,
	})
}

func CreateReserve(c *gin.Context) {
	var reserva Models.Reserve
	c.BindJSON(&reserva)
	reserva.Id_cliente = "asdf"
	reserva.Estado = "Pendiente"
	reserva.SetId_Reserva(Common.GenerateRandomIdWithLength(10))
	Models.CreateReserve(&reserva)
	c.JSON(200, reserva)
}
