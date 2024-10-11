package socket

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/websocket"
	"sync"
)

const (
	publicWebsocketURL = "wss://api.upbit.com/websocket/v1"
)

type PublicWebSocket struct {
	conn *websocket.Conn
	mu   sync.Mutex
}

func NewPublicWebSocket() (*PublicWebSocket, error) {
	conn, _, err := websocket.DefaultDialer.Dial(publicWebsocketURL, nil)
	if err != nil {
		return nil, err
	}
	return &PublicWebSocket{conn: conn}, nil
}

func (p *PublicWebSocket) Subscribe(typeField TypeField, format string) error {
	request := p.parseParams(typeField, format)

	message, err := json.Marshal(request)
	if err != nil {
		return err
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	err = p.conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		return err
	}
	return nil
}

func (p *PublicWebSocket) parseParams(typeField TypeField, format string) []Request {
	request := []Request{
		{
			Ticket: typeField.Ticket,
		},
		{
			Type:  typeField.Type,
			Codes: typeField.Codes,
		},
		{
			Format: "DEFAULT",
		},
	}
	if format != "" {
		request[2].Format = format
	}
	return request
}

func (p *PublicWebSocket) ReadMessage(v interface{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	_, message, err := p.conn.ReadMessage()
	if err != nil {
		return err
	}

	reader := bytes.NewReader(message)
	decoder := json.NewDecoder(reader)

	if err := decoder.Decode(v); err != nil {
		return err
	}

	return nil
}

func (p *PublicWebSocket) Close() error {
	return p.conn.Close()
}
