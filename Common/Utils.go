package Common

import (
	"gx_myfood/Config"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

func GenerateRandomId() string {
	var uuid = strings.Split(uuid.NewString(), "-")
	var newUuid string
	for _, value := range uuid {
		newUuid += value
	}
	return newUuid
}

func GenerateRandomIdWithLength(length int) string {
	var uuid = GenerateRandomId()
	return uuid[1:length]
}

func GenToken(id_client string) string {
	jwt_token := jwt.New(jwt.GetSigningMethod("HS256"))
	jwt_token.Claims = jwt.MapClaims{
		"sub": id_client,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token, _ := jwt_token.SignedString([]byte(Config.NBSecretPassword))
	return token
}
