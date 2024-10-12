package socket

import (
	"github.com/google/uuid"
	"log"
	"testing"
)

func TestWebsocket_Public(t *testing.T) {
	ws, err := NewPublicWebSocket()
	if err != nil {
		log.Fatal("Error connecting to WebSocket:", err)
	}
	defer func(ws *PublicWebSocket) {
		err := ws.Close()
		if err != nil {
			log.Fatal("Error closing WebSocket:", err)
		}
	}(ws)

	typeField := TypeField{
		Ticket: uuid.New().String(),
		Type:   TypeTicker,
		Codes:  []string{"KRW-BTC", "KRW-ETH"},
	}
	err = ws.Subscribe(typeField, "DEFAULT")
	if err != nil {
		log.Fatal("Error subscribing to channels:", err)
	}

	for {
		got := new(TickerResponse)
		err := ws.ReadMessage(got)
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}

		log.Printf("Ticker: %s, TradePrice: %f\n", got.Code, got.TradePrice)
	}
}

func TestWebSocket_Private(t *testing.T) {
	accessKey := ""
	secretKey := ""

	ws, err := NewPrivateWebSocket(accessKey, secretKey)
	if err != nil {
		log.Fatal("Error connecting to WebSocket:", err)
	}
	defer func(ws *PrivateWebSocket) {
		err := ws.Close()
		if err != nil {
			log.Fatal("Error closing WebSocket:", err)
		}
	}(ws)

	types := TypeField{
		Ticket: uuid.New().String(),
		Type:   TypeMyOrder,
		Codes:  []string{},
	}

	err = ws.Subscribe(types, "DEFAULT")
	if err != nil {
		log.Fatal("Error subscribing to channels:", err)
	}

	for {
		got := new(MyOrderResponse)
		err := ws.ReadMessage(got)
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}
		log.Printf("MyOrder: %s, UUID: %s, State: %s\n", got.Code, got.UUID, got.State)
	}
}
func TestWebSocket_Private_Asset(t *testing.T) {
	accessKey := ""
	secretKey := ""

	ws, err := NewPrivateWebSocket(accessKey, secretKey)
	if err != nil {
		log.Fatal("Error connecting to WebSocket:", err)
	}
	defer func(ws *PrivateWebSocket) {
		err := ws.Close()
		if err != nil {
			log.Fatal("Error closing WebSocket:", err)
		}
	}(ws)

	types := TypeField{
		Ticket: uuid.New().String(),
		Type:   TypeMyAsset,
	}

	err = ws.Subscribe(types, "DEFAULT")
	if err != nil {
		log.Fatal("Error subscribing to channels:", err)
	}

	for {
		got := new(MyOrderResponse)
		err := ws.ReadMessage(got)
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}
		log.Printf("MyOrder: %s, UUID: %s, State: %s\n", got.Code, got.UUID, got.State)
	}
}
