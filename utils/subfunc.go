package utils

import (
	"github.com/google/uuid"
)

func CheckUuid(uuidString string) bool {
	err := uuid.Validate(uuidString)
	return err == nil
}
