package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lexgalante/go.gin/src/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("unable to load .env file")
	}

	router := gin.Default()

	routes.ConfigureRouters(router)

	router.Run(os.Getenv("APP_PORT"))
}
