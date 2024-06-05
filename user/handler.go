package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/novalyezu/learnacademy-backend/helper"
)

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var input RegisterInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ValidationErrorResponse(err)
		c.JSON(http.StatusBadRequest, helper.WrapperResponse(http.StatusBadRequest, "BadRequest", "input validation errors", gin.H{"errors": errors}))
		return
	}

	newUser, errRegister := h.userService.Register(input)
	if errRegister != nil {
		c.JSON(http.StatusBadRequest, helper.WrapperResponse(http.StatusBadRequest, "Error", errRegister.Error(), nil))
		return
	}

	tokenString, errToken := helper.NewAuthTokenService().GenerateToken(newUser)
	if errToken != nil {
		c.JSON(http.StatusInternalServerError, helper.WrapperResponse(http.StatusInternalServerError, "InternalServerError", errToken.Error(), nil))
		return
	}

	output := FormatToAuthOutput(newUser, tokenString)

	c.JSON(http.StatusOK, helper.WrapperResponse(http.StatusOK, "OK", "Register success!", output))
}

func (h *UserHandler) Login(c *gin.Context) {
	var input LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ValidationErrorResponse(err)
		c.JSON(http.StatusBadRequest, helper.WrapperResponse(http.StatusBadRequest, "BadRequest", "input validation errors", gin.H{"errors": errors}))
		return
	}

	user, errLogin := h.userService.Login(input)
	if errLogin != nil {
		c.JSON(http.StatusUnauthorized, helper.WrapperResponse(http.StatusUnauthorized, "Unauthorized", errLogin.Error(), nil))
		return
	}

	formattedUser := FormatToUserOutput(user)

	tokenString, errToken := helper.NewAuthTokenService().GenerateToken(formattedUser)
	if errToken != nil {
		c.JSON(http.StatusInternalServerError, helper.WrapperResponse(http.StatusInternalServerError, "InternalServerError", errToken.Error(), nil))
		return
	}

	output := FormatToAuthOutput(user, tokenString)

	c.JSON(http.StatusOK, helper.WrapperResponse(http.StatusOK, "OK", "Login success!", output))
}

func (h *UserHandler) GetMe(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(User)

	output := FormatToUserOutput(currentUser)

	c.JSON(http.StatusOK, helper.WrapperResponse(http.StatusOK, "OK", "Get me success!", output))
}
