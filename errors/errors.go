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

	ErrShortPassword = errors.New("password should be longer than 6 characters")
	ErrLatinPassword = errors.New("password should contains Latin characters and numbers")

	ErrBoardNotFound = errors.New("this board is not found")

	ErrListNotFound = errors.New("this list is not found")

	ErrTaskNotFound = errors.New("this task is not found")

	ErrCheckListNotFound = errors.New("this checklist is not found")

	ErrCheckListItemNotFound = errors.New("this checklistitem is not found")

	ErrCommentNotFound = errors.New("this comment is not found")

	ErrNoAccess = errors.New("user doesn't have access")
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

	ErrTaskNotFound: http.StatusNotFound,

	ErrCheckListNotFound: http.StatusNotFound,

	ErrCheckListItemNotFound: http.StatusNotFound,

	ErrCommentNotFound: http.StatusNotFound,

	ErrNoAccess: http.StatusForbidden,
}

func ConvertErrorToCode(err error) (code int) {
	code, isErrorExist := errorToCode[err]
	if !isErrorExist {
		code = http.StatusInternalServerError
	}
	return
}
