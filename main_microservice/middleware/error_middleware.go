package middleware

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func CheckError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := errors.Unwrap(c.Errors.Last())
			newErr := new(models.CustomError)
			newErr.CustomErr = err.Error()
			errJson, err1 := newErr.MarshalJSON()
			if err1 != nil {
				fmt.Println(err1.Error())
				return
			}
			c.Data(customErrors.ConvertErrorToCode(err), "application/json; charset=utf-8", errJson)
			return
		}
	}
}
