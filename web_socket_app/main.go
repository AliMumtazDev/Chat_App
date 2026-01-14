package main

import (
	"log"

	routes "github.com/AliMumtazDev/socket/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	webSocketRouter := routes.NewRouter()
	err = webSocketRouter.Engine.Run(":8004")
	if err != nil {
		log.Fatalf("Websocket server failed to start: %s", err)
	}
}
