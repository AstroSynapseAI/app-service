package discord

import (
	"fmt"
	"github.com/AstroSynapseAI/app-service/engine/chains"
	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

type DiscordPlugin struct {
	clients []DiscordClient
}

type DiscordClient struct {
	session      *discordgo.Session
	activePlugin *models.ActivePlugin
}

func NewDiscordPlugin(pluginRecord models.Plugin) *DiscordPlugin {
	plugin := &DiscordPlugin{
		clients: []DiscordClient{},
	}

	for _, activePlugin := range pluginRecord.ActivePlugins {
		session, err := discordgo.New("Bot " + activePlugin.Token)
		if err != nil {
			fmt.Println("Failed to create Discord session:", err)
		}

		client := DiscordClient{
			session:      session,
			activePlugin: &activePlugin,
		}
		plugin.clients = append(plugin.clients, client)
	}

	return plugin
}

func (plugin *DiscordPlugin) OpenConnection(db *database.Database) {
	chain := chains.NewAsaiChain(db)

	// Create a channel of os.Signals for each client.
	stops := make([]chan os.Signal, len(plugin.clients))

	for i, client := range plugin.clients {
		if !client.activePlugin.IsActive {
			return
		}

		avatarID := client.activePlugin.AvatarID
		handler := NewAvatarHandler(avatarID, chain)
		client.session.AddHandler(handler.Messages)

		err := client.session.Open()
		if err != nil {
			fmt.Printf("Failed to create Discord session for client with AvatarID %v: %v\n", avatarID, err)
		}

		stops[i] = make(chan os.Signal, 1)
		signal.Notify(stops[i], syscall.SIGTERM)

		// Launch a goroutine for each client.
		go func(stop chan os.Signal, client *DiscordClient) {
			<-stop

			// Cleanly close down the Discord session.
			err := client.session.Close()
			if err != nil {
				fmt.Printf("Failed to close Discord session for client with AvatarID %v: %v\n", client.activePlugin.AvatarID, err)
			}

		}(stops[i], &client)
	}
}
