package utils

import "github.com/google/uuid"

func IsStringUUID(str string) (bool) {
	_, err := uuid.Parse(str)
	if err != nil {
		return false
	}
  return true
}