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

func (h *Handler) InitRoutes() *gin.Engine {
	handler := gin.Default()
	auth := handler.Group("")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}

	return handler
}
