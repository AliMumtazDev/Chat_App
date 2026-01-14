package connection

import (
    "log"
    "time"

    "github.com/gorilla/websocket"
)

var Conn *websocket.Conn

// ConnectToWebSocketServer keeps retrying until it connects, without crashing the app.
func ConnectToWebSocketServer(baseURL string, key string) {
    wsURL := baseURL + "?key=" + key
    backoff := time.Second * 2
    for {
        conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
        if err != nil {
            log.Printf("WebSocket connect failed: %v; retrying in %s", err, backoff)
            time.Sleep(backoff)
            if backoff < 30*time.Second {
                backoff *= 2
            }
            continue
        }
        Conn = conn
        log.Println("Connected to WebSocket server")
        return
    }
}
