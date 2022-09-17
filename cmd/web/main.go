package main

import (
	"log"

	"github.com/bigpaulie/go-todo/pkg/models"
	"github.com/bigpaulie/go-todo/pkg/routing"
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

	// initialize routes
	routing.SetupApiRouting(router)

	log.Fatalln(router.Run())
}
