package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/AstroSynapseAI/app-service/engine/chains"
	"github.com/gorilla/websocket"
)

var (
	// pongWait is how long we will await a pong response from client
	pongWait = 10 * time.Second
	// pingInterval has to be less than pongWait, We cant multiply by 0.9 to get 90% of time
	// Because that can make decimals, so instead *9 / 10 to get 90%
	// The reason why it has to be less than PingRequency is becuase otherwise it will
	// send a new Ping before getting response
	pingInterval = (pongWait * 9) / 10
	// Prompt for intializing conversation
	prompt = "New user, has connected. Invoke the onboarding_script.txt and welcome user."
)

type Client struct {
	sessionID  string
	egress     chan []byte
	connection *websocket.Conn
	manager    *Manager
	asaiChain  *chains.AsaiChain
}

func NewClient(conn *websocket.Conn, manager *Manager, chain *chains.AsaiChain) *Client {
	return &Client{
		egress:     make(chan []byte),
		connection: conn,
		manager:    manager,
		asaiChain:  chain,
	}
}

func (client *Client) MaintainConnection(ctx context.Context) {
	ticker := time.NewTicker(pingInterval)

	// Configure Wait time for Pong response, use Current time + pongWait
	// This has to be done here to set the first initial timer.
	if err := client.connection.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Println("Failed to set read deadline:", err)
		return
	}

	client.connection.SetPongHandler(client.PongHandler)

	for {
		// fmt.Println("Ping...")
		err := client.connection.WriteMessage(websocket.PingMessage, []byte{})
		if err != nil {
			fmt.Println("Ping failed:", err)
			return
		}
		// Wait for next tick
		<-ticker.C
	}
}

func (client *Client) PongHandler(pongMsg string) error {
	// Current time + Pong Wait time

	err := client.connection.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		fmt.Println("Failed to set read deadline in pong handler:", err)
		return err
	}

	return nil
}

func (client *Client) ReadMsgs(ctx context.Context) {
	defer func() {
		client.manager.removeClient(client)
	}()

	for {
		_, payload, err := client.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("error reading message: %v", err)
			}
			break
		}

		var request struct {
			SessionId  string `json:"session_id"`
			UserPrompt string `json:"user_prompt"`
		}

		if err = json.Unmarshal(payload, &request); err != nil {
			fmt.Println("error marshalling message: ", err)
			break
		}

		client.sessionID = request.SessionId

		client.asaiChain.SetSessionID(request.SessionId)

		client.asaiChain.Stream = func(ctx context.Context, chunk []byte) {
			client.egress <- chunk
		}

		if request.UserPrompt == "new user connected" {
			request.UserPrompt = prompt
		}

		go func() {
			if err = client.asaiChain.Run(ctx, request.UserPrompt); err != nil {
				fmt.Println("error Asai running chain: ", err)
				// Send an error message
				errMessage, _ := json.Marshal(map[string]string{
					"step": "error",
				})

				client.egress <- errMessage
				return
			}
		}()

	}
}

func (client *Client) SendMsgs(ctx context.Context) {
	defer func() {
		client.manager.removeClient(client)
	}()

	for msg := range client.egress {
		// Write a Regular text message to the connection
		if err := client.connection.WriteMessage(websocket.TextMessage, msg); err != nil {
			fmt.Println("Sending message failed:", err)
			return
		}

		// fmt.Println("Msg sent...")
	}
}
