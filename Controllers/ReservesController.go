package Controller

import (
	"encoding/json"
	"fmt"
	"gx_myfood/Common"
	"gx_myfood/Models"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetReserves(c *gin.Context) {
	var reserves []Models.Reserve
	client_id, _ := c.Get("client_id")

	if err := Models.GetReserves(&reserves, client_id.(string)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	for i, n := range reserves {
		reserves[i].Fecha = strings.Split(n.Fecha, "T")[0]
	}

	c.JSON(http.StatusOK, gin.H{
		"reservas": reserves,
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
		for i, n := range bannedDays {
			bannedDays[i] = strings.Split(n, "T")[0]
		}
		c.JSON(http.StatusOK, bannedDays)
		return
	}

	c.AbortWithStatus(http.StatusBadRequest)
	return
}

func GetImage(c *gin.Context) {
	buff, err := ioutil.ReadFile("Assets/Images/test.jpg")

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.Header("Content-Type", "image/jpeg")
	c.Writer.Write(buff)
}

func GetReserve(c *gin.Context) {
	var reserva Models.Reserve
	id_reserva := c.Param("id_reserva")

	if err := Models.GetReserve(&reserva, id_reserva); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	reserva.Fecha = strings.Split(reserva.Fecha, "T")[0]

	c.JSON(http.StatusOK, reserva)
}

type Reserve struct {
	Fecha        string `json:"fecha"`
	Tipo         string `json:"tipo"`
	N_comensales int    `json:"n_comensales"`
}

func GetPDFReserve(c *gin.Context) {
	var reserva Models.Reserve
	id_reserva := c.Param("id_reserva")

	if err := Models.GetReserve(&reserva, id_reserva); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	reserve := Reserve{strings.Split(reserva.Fecha, "T")[0], reserva.Tipo, reserva.N_comensales}
	reserveJSON, _ := json.Marshal(reserve)
	var list []map[string]interface{}
	var reserveMap map[string]interface{}
	json.Unmarshal(reserveJSON, &reserveMap)
	for key, value := range reserveMap {
		if key == "fecha" || key == "tipo" || key == "n_comensales" {
			tempMap := make(map[string]interface{})
			tempMap[key] = value
			list = append(list, tempMap)
		}
	}
	c.JSON(http.StatusOK, list)
}

func CreateReserve(c *gin.Context) {
	var input Models.ReserveInput
	client_id, _ := c.Get("client_id")
	c.BindJSON(&input)
	var reserva Models.Reserve = Models.Reserve(input)
	reserva.Id_cliente = client_id.(string)
	reserva.Estado = "Pendiente"
	f_date, _ := time.Parse("02/01/2006", reserva.Fecha)
	reserva.Fecha = f_date.Format("2006-01-02")
	reserva.SetId_Reserva(Common.GenerateRandomIdWithLength(10))

	if err := Models.CreateReserve(&reserva); err != nil {
		c.AbortWithStatus(http.StatusNotImplemented)
		return
	}
	c.JSON(http.StatusOK, reserva.Id_reserva)
}
