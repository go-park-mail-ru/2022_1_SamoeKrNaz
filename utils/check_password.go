package utils

import (
	customErrors "PLANEXA_backend/errors"
	"unicode"
)

func CheckPassword(pass string) error {
	if len(pass) <= 6 {
		return customErrors.ErrPassword
	}

	for i := 0; i < len(pass); i++ {
		if pass[i] > unicode.MaxASCII {
			return customErrors.ErrPassword
		}
	}
	return nil
}
