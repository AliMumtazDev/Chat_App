package routes

import (
	"github.com/AliMumtazDev/Go_Chat_App/auth"
)

func (r *SocketRouter) SocketRoutes() {
	protected := r.Engine.Group("/protected")
	protected.Use(auth.WSMiddleware())
	{
		protected.GET("/ws", r.RegisterWebSocketRoute)
	}
	backend := r.Engine.Group("/backend")
	backend.Use(auth.BackendWSMiddleware())
	{
		backend.GET("/ws", r.SaveBackendConnection)
	}
}
