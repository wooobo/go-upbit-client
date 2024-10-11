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
