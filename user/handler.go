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
		helper.ErrorHandler(c, err)
		return
	}

	newUser, errRegister := h.userService.Register(input)
	if errRegister != nil {
		helper.ErrorHandler(c, errRegister)
		return
	}

	tokenString, errToken := helper.NewAuthTokenService().GenerateToken(newUser)
	if errToken != nil {
		helper.ErrorHandler(c, errToken)
		return
	}

	output := FormatToAuthOutput(newUser, tokenString)

	c.JSON(http.StatusOK, helper.WrapperResponse(http.StatusOK, "OK", "Register success!", output))
}

func (h *UserHandler) Login(c *gin.Context) {
	var input LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		helper.ErrorHandler(c, err)
		return
	}

	user, errLogin := h.userService.Login(input)
	if errLogin != nil {
		helper.ErrorHandler(c, errLogin)
		return
	}

	formattedUser := FormatToUserOutput(user)

	tokenString, errToken := helper.NewAuthTokenService().GenerateToken(formattedUser)
	if errToken != nil {
		helper.ErrorHandler(c, errToken)
		return
	}

	output := FormatToAuthOutput(user, tokenString)

	c.JSON(http.StatusOK, helper.WrapperResponse(http.StatusOK, "OK", "Login success!", output))
}

func (h *UserHandler) GetMe(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(UserOutput)

	c.JSON(http.StatusOK, helper.WrapperResponse(http.StatusOK, "OK", "Get me success!", currentUser))
}
