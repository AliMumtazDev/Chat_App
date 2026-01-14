package routes

import (
	"log"

	userserviceimpl "github.com/AliMumtazDev/Go_Chat_App/api/message_service"
	"github.com/AliMumtazDev/Go_Chat_App/database/mongodb"
	socketinterface "github.com/AliMumtazDev/socket/web_socket"
	"github.com/AliMumtazDev/socket/web_socket/websocket_impl"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type SocketRouter struct {
	Engine         *gin.Engine
	WebSocket      socketinterface.WebSocketService
	Messageservice userserviceimpl.UserService
	MongoDB        mongodb.Storage
}

func NewRouter() *SocketRouter {

	connMongo, err := mongodb.MOngoConn()
	if err != nil {
		log.Fatalf("MongoDB connection error: %s", err)
	}
	messagedb := mongodb.NewStorage(connMongo)

	websockets := websocket_impl.NewWebSocketService(messagedb)
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowWebSockets:  true,
	}))
	router := &SocketRouter{
		Engine:    engine,
		WebSocket: websockets,
		// Messageservice: userService,
		MongoDB: messagedb,
	}
	router.SocketRoutes()
	return router
}
