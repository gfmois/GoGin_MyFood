package Controller

import (
	"gx_myfood/Models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type updateRequest struct {
	Nombre     string `json:"nombre"`
	Contraseña string `json:"contraseña"`
	Telefono   string `json:"telefono"`
}

func GetProfile(c *gin.Context) {
	client_model, _ := c.Get("client_model")
	c.JSON(http.StatusOK, client_model)
}

func UpdateProfile(c *gin.Context) {
	var newClient Models.Client
	var oldClient Models.Client
	if err := c.BindJSON(&newClient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client_model, _ := c.Get("client_model")

	oldClient = client_model.(Models.Client)
	if newClient.Nombre != "" {
		oldClient.Nombre = newClient.Nombre
	}
	if newClient.Contraseña != "" {
		passwordEncoded, _ := bcrypt.GenerateFromPassword([]byte(newClient.Contraseña), bcrypt.DefaultCost)
		oldClient.Contraseña = string(passwordEncoded)
	}
	if newClient.Telefono != "" {
		oldClient.Telefono = newClient.Telefono
	}

	if err := Models.UpdateProfile(&oldClient); err != nil {
		c.AbortWithStatus(http.StatusNotImplemented)
		return
	}

	c.JSON(http.StatusOK, oldClient)
}
