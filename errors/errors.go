package customErrors

import (
	"errors"
	"net/http"
)

var (
	ErrBadInputData = errors.New("bad input data")

	ErrUnauthorized     = errors.New("user is not authorized")
	ErrUsernameExist    = errors.New("this username already exists")
	ErrUsernameNotExist = errors.New("this username doesn`t exists")
	ErrUserNotFound     = errors.New("this user is not found")
	ErrUserHasntBoards  = errors.New("this user hasn`t boards")

	ErrShortPassword = errors.New("password should be longer than 6 characters")
	ErrLatinPassword = errors.New("password should contains Latin characters and numbers")

	ErrBoardNotFound = errors.New("this board is not found")

	ErrListNotFound = errors.New("this list is not found")

	ErrTaskNotFound = errors.New("this task is not found")
	ErrAccess       = errors.New("user doesn't have access")
)

var errorToCode = map[error]int{
	ErrBadInputData: http.StatusBadRequest,

	ErrUnauthorized:     http.StatusUnauthorized,
	ErrUsernameExist:    http.StatusConflict,
	ErrUsernameNotExist: http.StatusBadRequest,
	ErrUserNotFound:     http.StatusNotFound,

	ErrShortPassword: http.StatusBadRequest,
	ErrLatinPassword: http.StatusBadRequest,

	ErrBoardNotFound: http.StatusNotFound,

	ErrListNotFound: http.StatusNotFound,

	ErrTaskNotFound:    http.StatusNotFound,
	ErrUserHasntBoards: http.StatusForbidden,
	ErrAccess:          http.StatusForbidden,
}

func ConvertErrorToCode(err error) (code int) {
	code, isErrorExist := errorToCode[err]
	if !isErrorExist {
		code = http.StatusInternalServerError
	}
	return
}
