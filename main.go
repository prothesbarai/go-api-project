package main

import (
	"fmt"
	"go-api-project/database"
	"go-api-project/logger"
	"go-api-project/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func init(){
	err := godotenv.Load()
	if(err != nil){
		logger.AppLogger.Error.Println("Couldn't Load Env file : ",err)
		os.Exit(1)
	}
}

func main() {
	router := gin.Default()
	database.ConnectionDB()
	routes.SetupRoutes(router)
	port := os.Getenv("PORT") // live server দিয়ে দিবে
	fmt.Print("Port From env : ",port)
	if(port == ""){port = "8080"} // default local development
	router.Run(":"+port)
}