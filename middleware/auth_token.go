package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/novalyezu/learnacademy-backend/helper"
	"github.com/novalyezu/learnacademy-backend/user"
)

type AuthTokenMiddleware interface {
	VerifyToken() gin.HandlerFunc
}

type authTokenMiddleware struct {
	authTokenService helper.AuthTokenService
	userService      user.UserService
}

func NewAuthTokenMiddleware(authTokenService helper.AuthTokenService, userService user.UserService) AuthTokenMiddleware {
	return &authTokenMiddleware{
		authTokenService: authTokenService,
		userService:      userService,
	}
}

func (m *authTokenMiddleware) VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if !strings.Contains(authorization, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.WrapperResponse(http.StatusUnauthorized, "Unauthorized", "invalid token", nil))
			return
		}

		var tokenString string
		tokenSplit := strings.Split(authorization, " ")
		if len(tokenSplit) == 2 {
			tokenString = tokenSplit[1]
		}

		token, errToken := m.authTokenService.ValidateToken(tokenString)
		if errToken != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.WrapperResponse(http.StatusUnauthorized, "Unauthorized", "invalid token", nil))
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.WrapperResponse(http.StatusUnauthorized, "Unauthorized", "invalid token", nil))
			return
		}

		rsUser, errUser := m.userService.GetById(claim["payload"].(map[string]any)["id"].(string))
		if errUser != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.WrapperResponse(http.StatusUnauthorized, "Unauthorized", "invalid token", nil))
			return
		}

		userOutput := user.FormatToUserOutput(rsUser)

		c.Set("currentUser", userOutput)
	}
}
