package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"realtimeChat/internal/domain"
)

func (h *Handler) signUp(c *gin.Context) {
	user := domain.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := h.service.Authorization.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": userID})
}

type input struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	input := input{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, token, err := h.service.Authorization.CheckUser(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("authorization", token, 3600*6, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"id": userID})
}

func (h *Handler) logout(c *gin.Context) {
	c.SetCookie("authorization", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}
