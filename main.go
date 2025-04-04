package main

import (
	"log"
	"api.finance.com/config"
	"api.finance.com/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
	godotenv.Load();
	r := gin.Default()
	apiConfig := config.LoadConfig();
	db, err := config.GetDBConnection()
	if (err != nil) {
		log.Fatal(err)
	}

	routes.SetupRoutes(r, apiConfig, db)

	r.Run(apiConfig.ServerPort)
}