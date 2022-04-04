package utils

import (
	customErrors "PLANEXA_backend/errors"
	"github.com/google/uuid"
	"unicode"
)

func CheckPassword(pass string) error {
	if len(pass) <= 6 {
		return customErrors.ErrShortPassword
	}

	for i := 0; i < len(pass); i++ {
		if pass[i] > unicode.MaxASCII {
			return customErrors.ErrLatinPassword
		}
	}
	return nil
}

func GenerateSessionToken() string {
	return uuid.NewString()
}
