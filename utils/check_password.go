package utils

import (
	"errors"
	"unicode"
)

func CheckPassword(pass string) error {
	if len(pass) <= 6 {
		return errors.New("the password should be longer than 6 characters")
	}

	for i := 0; i < len(pass); i++ {
		if pass[i] > unicode.MaxASCII {
			return errors.New("the password should contains Latin characters and numbers")
		}
	}
	return nil
}
