package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/novalyezu/learnacademy-backend/router"
	"github.com/novalyezu/learnacademy-backend/user"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// For Vercel Serverless Functions

var app *gin.Engine

func init() {
	app = gin.Default()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Etc/UTC",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DATABASE"), os.Getenv("POSTGRES_PORT"))
	db, errGorm := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errGorm != nil {
		log.Fatal(errGorm.Error())
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", user.PasswordValidator)
	}

	router.InitRouter(app, db)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
