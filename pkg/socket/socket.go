package socket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

const (
	publicWebsocketURL  = "wss://api.upbit.com/websocket/v1"
	privateWebsocketURL = "wss://api.upbit.com/websocket/v1/private"
	pongWait            = 120 * time.Second
	pingPeriod          = (pongWait * 9) / 10
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

	return decoder.Decode(v)
}

func (p *PublicWebSocket) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.conn.Close()
}

type PrivateWebSocket struct {
	conn      *websocket.Conn
	mu        sync.Mutex
	accessKey string
	secretKey string
	closeCh   chan struct{}
}

func NewPrivateWebSocket(accessKey, secretKey string) (*PrivateWebSocket, error) {
	token, err := generateJWT(accessKey, secretKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate JWT: %v", err)
	}
	headers := make(map[string][]string)
	headers["Authorization"] = []string{"Bearer " + token}
	conn, _, err := websocket.DefaultDialer.Dial(privateWebsocketURL, headers)
	if err != nil {
		return nil, err
	}

	closeCh := make(chan struct{})
	go func() {
		err := ping(conn, closeCh)
		if err != nil {
			fmt.Println("Error sending PING frame:", err)
		} else {
			fmt.Println("PING frame sent")
		}
	}()

	return &PrivateWebSocket{
		conn:      conn,
		accessKey: accessKey,
		secretKey: secretKey,
		closeCh:   closeCh,
	}, nil
}

func generateJWT(accessKey, secretKey string) (string, error) {
	claims := jwt.MapClaims{
		"access_key": accessKey,
		"nonce":      uuid.New().String(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func (p *PrivateWebSocket) Subscribe(typeField TypeField, format string) error {
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
	if typeField.IsOnlySnapshot {
		request[1].IsOnlyRealtime = typeField.IsOnlySnapshot
	}
	if typeField.IsOnlyRealtime {
		request[1].IsOnlyRealtime = typeField.IsOnlyRealtime
	}
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

func (p *PrivateWebSocket) ReadMessage(v interface{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	_, message, err := p.conn.ReadMessage()
	if err != nil {
		return err
	}
	reader := bytes.NewReader(message)
	decoder := json.NewDecoder(reader)
	return decoder.Decode(v)
}

func (p *PrivateWebSocket) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	close(p.closeCh)
	return p.conn.Close()
}

func ping(conn *websocket.Conn, closeCh chan struct{}) error {
	pingTicker := time.NewTicker(pingPeriod)
	defer pingTicker.Stop()
	for {
		select {
		case <-pingTicker.C:
			err := conn.WriteMessage(websocket.PingMessage, []byte("PING"))
			if err != nil {
				fmt.Println("Error sending PING frame:", err)
			}
		case <-closeCh:
		}
	}
}
