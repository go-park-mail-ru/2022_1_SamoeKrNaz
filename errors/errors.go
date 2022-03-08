package customErrors

import (
	"errors"
	"net/http"
)

var (
	ErrBadInputData  = errors.New("bad input data")
	ErrUnauthorized  = errors.New("user is not authorized")
	ErrUsernameExist = errors.New("this username already exists")
	ErrShotPassword  = errors.New("password should be longer than 6 characters")
	ErrLatinPassword = errors.New("password should contains Latin characters and numbers")
)

var errorToCode = map[error]int{
	ErrBadInputData:  http.StatusBadRequest,
	ErrUnauthorized:  http.StatusUnauthorized,
	ErrUsernameExist: http.StatusConflict,
	ErrShotPassword:  http.StatusBadRequest,
	ErrLatinPassword: http.StatusBadRequest,
}

func ConvertErrorToCode(err error) (code int) {
	code, isErrorExist := errorToCode[err]
	if !isErrorExist {
		code = http.StatusInternalServerError
	}
	return
}
