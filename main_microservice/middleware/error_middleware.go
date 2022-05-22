package middleware

import (
	customErrors "PLANEXA_backend/errors"
	"PLANEXA_backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CheckError(c *gin.Context) {
	if c.Errors != nil {
		newErr := new(models.CustomError)
		newErr.CustomErr = c.Errors.Last().Error()
		errJson, err := newErr.MarshalJSON()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		c.Data(customErrors.ConvertErrorToCode(c.Errors.Last()), "application/json; charset=utf-8", errJson)
		return
	}
	return
}
