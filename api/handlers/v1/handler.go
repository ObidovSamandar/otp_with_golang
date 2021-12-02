package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/obidovsamandar/go-task-auth/api/models"
)

const (
	//ErrorCodeInvalidURL ...
	ErrorCodeInvalidURL = "INVALID_URL"
	//ErrorCodeInvalidJSON ...
	ErrorCodeInvalidJSON = "INVALID_JSON"
	//ErrorCodeInternal ...
	ErrorCodeInternal = "INTERNAL"
	//ErrorCodeUnauthorized ...
	ErrorCodeUnauthorized = "UNAUTHORIZED"
	//ErrorCodeAlreadyExists ...
	ErrorCodeAlreadyExists = "ALREADY_EXISTS"
	//ErrorCodeNotFound ...
	ErrorCodeNotFound = "NOT_FOUND"
	//ErrorCodeInvalidCode ...
	ErrorCodeInvalidCode = "INVALID_CODE"
	//ErrorBadRequest ...
	ErrorBadRequest = "BAD_REQUEST"
	//ErrorCodeForbidden ...
	ErrorCodeForbidden = "FORBIDDEN"
	//ErrorCodeNotApproved ...
	ErrorCodeNotApproved = "NOT_APPROVED"
	//ErrorCodeWrongClub ...
	ErrorCodeWrongClub = "WRONG_CLUB"
	//ErrorCodePasswordsNotEqual ...
	ErrorCodePasswordsNotEqual = "PASSWORDS_NOT_EQUAL"
)

func handleInternalWithMessage(c *gin.Context, err error, message string) bool {
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Error: models.ResponseError{
				Code:    ErrorCodeInternal,
				Message: message,
			},
		})
		return true
	}

	return false
}

func handleBadRequestErrWithMessage(c *gin.Context, err error, message string) bool {
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			Error: models.ResponseError{
				Code:    ErrorCodeInvalidJSON,
				Message: "Invalid Json",
			},
		})
		return true
	}
	return false
}
