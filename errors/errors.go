package customErrors

import (
	"errors"
	"net/http"
)

var (
	ErrBadInputData  = errors.New("bad input data")
	ErrUnauthorized  = errors.New("user is not authorized YESYES")
	ErrUsernameExist = errors.New("this username already exists")
	ErrPassword      = errors.New("password should be longer than 6 characters and contains Latin characters and numbers")
)

var errorToCode = map[error]int{
	ErrBadInputData:  http.StatusBadRequest,
	ErrUnauthorized:  http.StatusUnauthorized,
	ErrUsernameExist: http.StatusConflict,
	ErrPassword:      http.StatusBadRequest,
}

func ConvertErrorToCode(err error) (code int) {
	code, isErrorExist := errorToCode[err]
	if !isErrorExist {
		code = http.StatusInternalServerError
	}
	return
}
