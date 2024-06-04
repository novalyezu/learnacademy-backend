package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/novalyezu/learnacademy-backend/router"
	"github.com/novalyezu/learnacademy-backend/user"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	app := gin.Default()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jayapura",
		"localhost", "postgres", "123qweasd", "db_learn_academy", 5432)
	db, errGorm := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errGorm != nil {
		log.Fatal(errGorm.Error())
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", user.PasswordValidator)
	}

	app.GET("/", HealthCheck)
	router.InitRouter(app, db)
	app.Run()
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]any{
		"status":  "OK",
		"message": "Server is up and running",
	})
}
