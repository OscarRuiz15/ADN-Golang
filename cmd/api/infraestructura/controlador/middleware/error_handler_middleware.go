package middleware

import (
	"ADN_Golang/cmd/api/dominio/exception"
	"ADN_Golang/pkg/apierror"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

const internalServerErrorMessage = "an error occurred during the processing of your request"

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err == nil {
			return
		}
		cause := errors.Cause(err.Err)

		if _, ok := cause.(exception.NotFound); ok {
			throwException(c, http.StatusNotFound, err.Err)
			return
		}

		if _, ok := cause.(exception.Duplicity); ok {
			throwException(c, http.StatusBadRequest, err.Err)
			return
		}

		if _, ok := cause.(exception.DataRequired); ok {
			throwException(c, http.StatusBadRequest, err.Err)
			return
		}

		log.Println("middleware error 500", cause)
		throwException(c, http.StatusInternalServerError, errors.New(internalServerErrorMessage))
	}
}

func throwException(ctx *gin.Context, status int, err error) {
	restErr := apierrors.NewApiError(err.Error(), http.StatusText(status), status)
	log.Println(restErr.Message())
	ctx.JSON(restErr.Status(), restErr)
}
