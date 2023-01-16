package Middlewares

import (
	"gx_myfood/Config"
	"gx_myfood/Models"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

//La función toma una cadena de token como entrada y devuelve una cadena de token sin el prefijo "BEARER " y un error (si es que ocurre alguno).
func stripBearerPrefixFromTokenString(tok string) (string, error) {
	if len(tok) > 5 && strings.ToUpper(tok[0:7]) == "BEARER " {
		return tok[7:], nil
	}
	return tok, nil
}

// Se utiliza para extraer el token JWT de la cabecera "Authorization" de una solicitud HTTP
var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
	request.HeaderExtractor{"Authorization"},
	stripBearerPrefixFromTokenString,
}

// Extrae de la cabecera el token o de los parametros el access_token
var AuthExtractor = &request.MultiExtractor{
	AuthorizationHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

//Se encarga de obtener información del usuario a través del id del cliente y almacenarlo en la variable myClient, luego esa información se guarda en el contexto para su uso futuro.
func UpdateContextClient(c *gin.Context, client_id string) {
	var myClient Models.Client
	if client_id != "" {
		Models.GetUserInfo(&myClient, client_id)
	}
	c.Set("client_id", client_id)
	c.Set("client_model", myClient)
}

// Este middleware se encarga de validar un token JWT enviado en las cabeceras de la petición, en caso de ser válido actualiza el contexto con el id del cliente, en caso contrario y si auto401 es true devuelve un estatus 401 de no autorizado, si auto401 es false simplemente no hace nada.
func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		UpdateContextClient(c, "")
		token, err := request.ParseFromRequest(c.Request, AuthExtractor, func(_ *jwt.Token) (interface{}, error) {
			b := ([]byte(Config.NBSecretPassword))
			return b, nil
		})

		if err != nil {
			if auto401 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			}
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			exp := int64(claims["exp"].(float64))
			if time.Now().Unix() > exp {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			}
			UpdateContextClient(c, claims["sub"].(string))
		}

	}
}
