package router

import (
	"net/http"

	"github.com/novalyezu/learnacademy-backend/community"
	"github.com/novalyezu/learnacademy-backend/helper"
	"github.com/novalyezu/learnacademy-backend/middleware"
	"github.com/novalyezu/learnacademy-backend/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(app *gin.Engine, db *gorm.DB) {
	fileService := helper.NewFileService()

	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)
	authTokenService := helper.NewAuthTokenService()
	authTokenMiddleware := middleware.NewAuthTokenMiddleware(authTokenService, userService)

	communityRepository := community.NewCommunityRepository(db)
	communityService := community.NewCommunityService(communityRepository)
	communityHandler := community.NewCommunityHandler(communityService, fileService)

	app.GET("/api", HealthCheck)

	apiV1 := app.Group("/api/v1")

	apiV1.POST("/auth/register", userHandler.Register)
	apiV1.POST("/auth/login", userHandler.Login)

	apiV1.GET("/me", authTokenMiddleware.VerifyToken(), userHandler.GetMe)

	apiV1.POST("/communities", authTokenMiddleware.VerifyToken(), communityHandler.CreateCommunity)
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]any{
		"status":  "OK",
		"message": "Server is up and running",
	})
}
