package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type WSHandler struct {
	Hub *Hub
}

func NewWSHandler(h *Hub) *WSHandler {
	return &WSHandler{Hub: h}
}

type roomInput struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func (h *WSHandler) CreateRoom(c *gin.Context) {
	roomInput := &roomInput{}
	err := c.ShouldBindJSON(&roomInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.Hub.Rooms[roomInput.ID] = &Room{
		ID:      roomInput.ID,
		Name:    roomInput.Name,
		Clients: make(map[string]*Client),
	}

	c.JSON(http.StatusCreated, roomInput)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *WSHandler) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Path /ws/joinRoom/:roomId?userId=xxx&username=xxx
	roomId := c.Param("roomId")
	userId := c.Query("userId")
	username := c.Query("username")

	cl := &Client{
		Conn:     conn,
		Message:  make(chan *Message, 10),
		ID:       userId,
		RoomID:   roomId,
		Username: username,
	}

	m := &Message{
		Content:  "A new user has joined the room",
		RoomID:   roomId,
		Username: username,
	}

	h.Hub.Register <- cl
	h.Hub.Broadcast <- m

	go cl.writeMessage()
	cl.readMessage(h.Hub)
}

type Rooms struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *WSHandler) GetRooms(c *gin.Context) {
	rooms := make([]Rooms, 0)

	for _, r := range h.Hub.Rooms {
		rooms = append(rooms, Rooms{ID: r.ID, Name: r.Name})
	}

	c.JSON(http.StatusOK, rooms)
}

type Clients struct {
	Username string
}

func (h *WSHandler) GetClients(c *gin.Context) {
	clients := make([]Clients, 0)
	roomID := c.Param("roomId")

	_, ok := h.Hub.Rooms[roomID]
	if !ok {
		c.JSON(http.StatusConflict, gin.H{"error": "room does not exist"})
		return
	}

	for _, c := range h.Hub.Rooms[roomID].Clients {
		clients = append(clients, Clients{Username: c.Username})
	}

	c.JSON(http.StatusOK, clients)
}
