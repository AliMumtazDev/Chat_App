package main

// @title Go Chat App API
// @version 1.0
// @description This is a chat application API built with Go and Gin.
// @host localhost:8002
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
import (
	"log"
	"os"

	authserviceimpl "github.com/AliMumtazDev/Go_Chat_App/api/auth_service"
	userserviceimpl "github.com/AliMumtazDev/Go_Chat_App/api/message_service"
	"github.com/AliMumtazDev/Go_Chat_App/database/mongodb"
	"github.com/AliMumtazDev/Go_Chat_App/database/postgresdb"
	routes "github.com/AliMumtazDev/Go_Chat_App/router"
	connection "github.com/AliMumtazDev/Go_Chat_App/socket_clint"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	key := os.Getenv("BACKEND_WS_KEY")
	if err != nil {
		log.Panicf("Error loading .env file: %s", err)
	}
	connpostgres, err := postgresdb.PostgresConn()
	if err != nil {
		log.Panicf("Error connecting to PostgreSQL: %s", err)
	}
	// Run DB migrations on startup to ensure tables exist
	if err := postgresdb.Migrate(connpostgres); err != nil {
		log.Panicf("Error running migrations: %s", err)
	}
	connmongo, err := mongodb.MOngoConn()
	if err != nil {
		log.Panicf("Error connecting to MongoDB: %s", err)
	}
	userdb := postgresdb.NewStorage(connpostgres)
	messagedb := mongodb.NewStorage(connmongo)
	authService := authserviceimpl.NewAuthService(authserviceimpl.NewAuthServiceImpl{
		UserAuth: userdb,
	})

	userService := userserviceimpl.NewUserService(messagedb)
	go func() {
		wsURL := os.Getenv("WEBSOCKET_URL")
		if wsURL == "" {
			// Prefer Docker service name by default; override in .env for local dev
			wsURL = "ws://websocket:8004/backend/ws"
		}
		log.Printf("Connecting to WebSocket at %s", wsURL)
		connection.ConnectToWebSocketServer(wsURL, key)
		log.Println("Connected to WebSocket server")
	}()
	httpRouter := routes.NewRouter(authService, userService)
	if err := httpRouter.Engine.Run(":8005"); err != nil {
		log.Fatalf("HTTP server failed to start: %s", err)
	}
}
