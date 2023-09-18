package http

import (
	"github.com/gin-gonic/gin"
	"realtimeChat/internal/service"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes(wsh *WSHandler) *gin.Engine {
	handler := gin.Default()
	auth := handler.Group("")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
		auth.GET("logout", h.logout)

		auth.POST("/ws/createRoom", wsh.CreateRoom)
		auth.GET("/ws/joinRoom/:roomId", wsh.JoinRoom)
		auth.GET("/ws/getRooms", wsh.GetRooms)
		auth.GET("/ws/getClients/:roomId", wsh.GetClients)
		//auth.GET("/ws/joinRoom/:roomID", wsh.JoinRoom)
	}

	return handler
}
