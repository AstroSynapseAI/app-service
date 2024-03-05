package server

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/AstroSynapseAI/app-service/app"
	"github.com/AstroSynapseAI/app-service/controllers"
	"github.com/AstroSynapseAI/app-service/controllers/ws"
	"github.com/AstroSynapseAI/app-service/engine/plugins"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/rest"
	"github.com/bwmarrin/discordgo"
)

type AsaiServer struct {
	Plugins       *plugins.PluginLoader
	discordClient *discordgo.Session
}

func NewAsaiServer() *AsaiServer {
	server := &AsaiServer{
		Plugins: plugins.NewLoader(),
	}

	return server
}

func (server *AsaiServer) Run(db *database.Database) error {
	server.Plugins.LoadConfig(db)

	router := rest.NewRouter()
	router.Mux.StrictSlash(true)

	// Serve API controllers
	router.Load(app.NewRoutes(db))

	// Serve WebSocket
	wsManager := ws.NewManager(db)
	router.Mux.HandleFunc("/ws/chat", wsManager.Handler)

	// Serve Websites
	webCtrl := controllers.NewWebController(router)
	webCtrl.Run()

	//
	// PORT is defined by heroku env
	port := os.Getenv("PORT")

	// If PORT is not defined, server is running locally
	if port == "" {
		err := http.ListenAndServe(":8082", router.Mux)
		if err != nil {
			fmt.Println("Failed to listen:", err)
			return err
		}
		return nil
	}

	// If PORT is defined, server is running on Heroku
	// Create a TCP listener with heroku port
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return err
	}

	// Serve the HTTP server
	err = http.Serve(listener, router.Mux)
	if err != nil {
		fmt.Println("Failed to serve:", err)
		return err
	}

	// Open all plugins connections
	server.Plugins.OpenConnection(db)

	return nil
}
