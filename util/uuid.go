package util

import (
	"github.com/google/uuid"
)

func GetUuid() string {
	return uuid.NewString()
}

func ValidateUuid(guid string) bool {
	_, parseError := uuid.Parse(guid)

	if parseError != nil {
		return false
	}

	return true
}
