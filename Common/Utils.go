package Common

import (
	"strings"

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
