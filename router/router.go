package router

import (
	"github.com/novalyezu/learnacademy-backend/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(app *gin.Engine, db *gorm.DB) {
	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)

	apiV1 := app.Group("/v1")

	apiV1.POST("/auth/register", userHandler.Register)
	apiV1.POST("/auth/login", userHandler.Login)
}
