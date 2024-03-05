package plugins

import (
	"github.com/AstroSynapseAI/app-service/engine/plugins/discord"
	"github.com/AstroSynapseAI/app-service/repositories"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
)

type PluginLoader struct {
	Plugins []Plugin
}

var _ Plugins = (*PluginLoader)(nil)

func NewLoader() *PluginLoader {
	return &PluginLoader{}
}

func (loader *PluginLoader) LoadConfig(db *database.Database) error {
	pluginsRepo := repositories.NewPluginsRepository(db)
	pluginsRecords := pluginsRepo.FetchAll()

	for _, pluginRecord := range pluginsRecords {
		if pluginRecord.Slug == "discord" {
			discordPlugin := discord.NewDiscordPlugin(pluginRecord)
			loader.Plugins = append(loader.Plugins, discordPlugin)
		}
	}

	return nil
}

func (loader *PluginLoader) OpenConnection(db *database.Database) error {

	for _, plugin := range loader.Plugins {
		plugin.OpenConnection(db)
	}

	return nil
}
