package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type NotFoundError struct {
	Message string
}

func NewNotFoundError(msg string) NotFoundError {
	return NotFoundError{Message: msg}
}

func (e NotFoundError) Error() string {
	return e.Message
}

type BadRequestError struct {
	Message string
}

func NewBadRequestError(msg string) BadRequestError {
	return BadRequestError{Message: msg}
}

func (e BadRequestError) Error() string {
	return e.Message
}

type UnauthorizedError struct {
	Message string
}

func NewUnauthorizedError(msg string) UnauthorizedError {
	return UnauthorizedError{Message: msg}
}

func (e UnauthorizedError) Error() string {
	return e.Message
}

func ErrorHandler(c *gin.Context, err any) {
	if validationError(c, err) {
		return
	}

	if badRequestError(c, err) {
		return
	}

	if unauthorizedError(c, err) {
		return
	}

	if notFoundError(c, err) {
		return
	}

	internalServerError(c, err)
}

func internalServerError(c *gin.Context, err any) {
	c.JSON(http.StatusInternalServerError, WrapperResponse(http.StatusInternalServerError, "InternalServerError", err.(error).Error(), nil))
}

func notFoundError(c *gin.Context, err any) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		c.JSON(http.StatusNotFound, WrapperResponse(http.StatusNotFound, "NotFound", exception.Error(), nil))
		return true
	} else {
		return false
	}
}

func badRequestError(c *gin.Context, err any) bool {
	exception, ok := err.(BadRequestError)
	if ok {
		c.JSON(http.StatusBadRequest, WrapperResponse(http.StatusBadRequest, "BadRequest", exception.Error(), nil))
		return true
	} else {
		return false
	}
}

func unauthorizedError(c *gin.Context, err any) bool {
	exception, ok := err.(UnauthorizedError)
	if ok {
		c.JSON(http.StatusUnauthorized, WrapperResponse(http.StatusUnauthorized, "Unauthorized", exception.Error(), nil))
		return true
	} else {
		return false
	}
}

func validationError(c *gin.Context, err any) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		var errors []string
		for _, e := range exception {
			switch e.Tag() {
			case "required":
				errors = append(errors, e.Field()+": field is required")
			case "email":
				errors = append(errors, e.Field()+": invalid email")
			case "password":
				errors = append(errors, e.Field()+": password min 6, uppercase, lowercase, number, special char")
			default:
				errors = append(errors, e.Field()+": "+e.Error())
			}
		}
		c.JSON(http.StatusBadRequest, WrapperResponse(http.StatusBadRequest, "BadRequest", "input validation errors", gin.H{"errors": errors}))
		return true
	} else {
		return false
	}
}
