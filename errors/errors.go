package customErrors

import (
	"errors"
	"net/http"
)

var (
	ErrBadInputData  = errors.New("bad input data")
	ErrUnauthorized  = errors.New("user is not authorized")
	ErrUsernameExist = errors.New("this username already exists")
)

var errorToCode = map[error]int{
	ErrBadInputData:  http.StatusBadRequest,
	ErrUnauthorized:  http.StatusUnauthorized,
	ErrUsernameExist: http.StatusConflict,
}

func ConvertErrorToCode(err error) (code int) {
	code, isErrorExist := errorToCode[err]
	if !isErrorExist {
		code = http.StatusInternalServerError
	}
	return
}
