package main

import (
	"log"

	"github.com/bigpaulie/go-todo/pkg/models"
	"github.com/bigpaulie/go-todo/pkg/routing"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// initialize database
	models.SetupDatabase()

	// initialize gin framework
	router := gin.Default()

	// Setup Middleware
	setupCorsMiddleWare(router)

	// initialize routes
	routing.SetupApiRouting(router)

	log.Fatalln(router.Run())
}

func setupCorsMiddleWare(e *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowCredentials = true

	e.Use(cors.New(config))
}
