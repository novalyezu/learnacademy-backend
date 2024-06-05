package router

import (
	"github.com/novalyezu/learnacademy-backend/helper"
	"github.com/novalyezu/learnacademy-backend/middleware"
	"github.com/novalyezu/learnacademy-backend/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(app *gin.Engine, db *gorm.DB) {
	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)
	authTokenService := helper.NewAuthTokenService()
	authTokenMiddleware := middleware.NewAuthTokenMiddleware(authTokenService, userService)

	apiV1 := app.Group("/v1")

	apiV1.POST("/auth/register", userHandler.Register)
	apiV1.POST("/auth/login", userHandler.Login)

	apiV1.GET("/me", authTokenMiddleware.VerifyToken(), userHandler.GetMe)
}
