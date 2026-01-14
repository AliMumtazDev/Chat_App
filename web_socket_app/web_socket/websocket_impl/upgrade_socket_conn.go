package websocket_impl

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/AliMumtazDev/Go_Chat_App/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var ConnLock = sync.Mutex{}

func (w *WebSocketServiceImpl) AddConn(userID string, wsConn *websocket.Conn, c *gin.Context) error {
	uid, err := strconv.Atoi(userID)
	if err != nil {
		log.Println("Error converting userID to int:", err)
		return err
	}
	log.Println("User ID:", uid)
	ConnLock.Lock()
	defer ConnLock.Unlock()

	if existingClient, ok := w.Clients[uid]; ok {
		existingClient.Close()
		delete(w.Clients, uid)
	}
	w.Clients[uid] = wsConn
	log.Println("User connected:", uid)
	fmt.Println("connected clients:", w.Clients)

	go func() {
		defer func() {
			ConnLock.Lock()
			defer ConnLock.Unlock()
			if w.Clients[uid] == wsConn {
				delete(w.Clients, uid)
				wsConn.Close()
				log.Println("User disconnected:", uid)
			}
		}()

		for {
			var incoming models.ServerMesageToSocket
			err := wsConn.ReadJSON(&incoming)
			if err != nil {
				log.Println("Error reading JSON:", err)
				break
			}
			log.Println("Received JSON from", uid, ":", incoming)
			if incoming.Action == "send" {
				recieverID := int(incoming.DestinationID)
				log.Printf("Received message for receiver_id: %d", recieverID)
				if recieverID == uid {
					log.Println("Cannot send message to self")
					continue
				}
				ConnLock.Lock()
				conn, ok := w.Clients[recieverID]
				ConnLock.Unlock()
				if ok {
					message := map[string]any{
						"receiver_id": recieverID,
						"content":     incoming.Content,
					}
					err := conn.WriteJSON(message)
					if err != nil {
						log.Println("Error writing JSON to receiver_id:", err)
					} else {
						log.Println("Message sent to receiver_id:", recieverID)
					}
				} else {
					log.Println("Receiver not connected:", recieverID)
				}
			}
		}
	}()
	return nil
}