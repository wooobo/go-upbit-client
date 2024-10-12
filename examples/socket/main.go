package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"

	wsClient "github.com/wooobo/go-upbit-client/pkg/socket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := wsClient.NewPublicWebSocket()
	if err != nil {
		log.Fatal("Error connecting to WebSocket:", err)
	}
	defer func(ws *wsClient.PublicWebSocket) {
		err := ws.Close()
		if err != nil {
			log.Fatal("Error closing WebSocket:", err)
		}
	}(ws)

	typeField := wsClient.TypeField{
		Ticket: uuid.New().String(),
		Type:   wsClient.TypeTicker,
		Codes:  []string{"KRW-BTC", "KRW-ETH"},
	}
	err = ws.Subscribe(typeField, "DEFAULT")
	if err != nil {
		log.Fatal("Error subscribing to channels:", err)
	}

	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}

	for {
		data := new(wsClient.TickerResponse)
		err := ws.ReadMessage(data)
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}

		err = connection.WriteJSON(data)

		log.Printf("Ticker: %s, TradePrice: %f\n", data.Code, data.TradePrice)
	}
}

func main() {
	http.HandleFunc("/ws/ticker/won/all", serveWs)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "examples/socket/main.html")
	})

	port := "8080"
	url := fmt.Sprintf("http://localhost:%s", port)

	fmt.Printf("Server is starting on %s\n", url)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
