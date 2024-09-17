package handler

import (
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var clients = make(map[*websocket.Conn]bool)
var mu sync.Mutex

// WebSocketHandler là handler cho WebSocket
func (w WasteBinHandler) WebSocketHandler() fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		mu.Lock()
		clients[c] = true
		mu.Unlock()
		defer func() {
			mu.Lock()
			delete(clients, c)
			mu.Unlock()
			c.Close()
		}()

		for {
			_, _, err := c.ReadMessage()
			if err != nil {
				break
			}
		}
	})
}

// BroadcastToClients sẽ được gọi khi có dữ liệu cập nhật
func BroadcastToClients(message string) {
	mu.Lock()
	defer mu.Unlock()
	for client := range clients {
		if err := client.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			client.Close()
			delete(clients, client)
		}
	}
}
