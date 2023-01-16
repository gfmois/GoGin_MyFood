package Controller

import (
	"encoding/base64"
	"gx_myfood/Common"
	"gx_myfood/Models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func responseClient(client Models.Client) gin.H {
	return gin.H{
		"token":  Common.GenToken(client.Id_cliente),
		"type":   "Bearer",
		"nombre": client.Nombre,
		"email":  client.Email,
		"avatar": client.Avatar,
	}
}

func Register(c *gin.Context) {
	var client Models.Client
	c.BindJSON(&client)
	client.Id_cliente = Common.GenerateRandomIdWithLength(10)
	client.Avatar = "https://api.multiavatar.com/" + base64.StdEncoding.EncodeToString([]byte(client.Nombre)) + ".png"
	passwordEncoded, _ := bcrypt.GenerateFromPassword([]byte(client.Contraseña), bcrypt.DefaultCost)
	client.Contraseña = string(passwordEncoded)

	if err := Models.Register(&client); err != nil {
		c.AbortWithStatus(http.StatusNotImplemented)
		return
	}

	c.JSON(http.StatusOK, responseClient(client))
}

func Login(c *gin.Context) {
	var client Models.Client
	c.BindJSON(&client)

	if err := Models.Login(&client); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, responseClient(client))
}
