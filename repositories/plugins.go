package repositories

import (
	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/crud/orms/gorm"
	db "gorm.io/gorm"
)

type PluginsRepository struct {
	Repo   *gorm.Repository[models.Plugin]
	Active *gorm.Repository[models.ActivePlugin]
}

func NewPluginsRepository(db *database.Database) *PluginsRepository {
	return &PluginsRepository{
		Repo:   gorm.NewRepository[models.Plugin](db, models.Plugin{}),
		Active: gorm.NewRepository[models.ActivePlugin](db, models.ActivePlugin{}),
	}
}

func (plugin *PluginsRepository) FetchAll() []models.Plugin {
	var plugins []models.Plugin
	query := plugin.Repo.DB.Preload("ActivePlugins")

	result := query.Find(&plugins)
	if result.Error != nil {
		return []models.Plugin{}
	}
	return plugins
}

func (plugin *PluginsRepository) SaveActivePlugin(avatarData models.ActivePlugin) error {
	result := plugin.Active.DB.Save(&avatarData)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (plugin *PluginsRepository) ToggleActivePlugin(avatarID uint, pluginID uint, active bool) error {
	var activePlugin models.ActivePlugin

	result := plugin.Active.DB.Where("avatar_id = ? AND plugin_id = ?", avatarID, pluginID).First(&activePlugin)
	if result.Error == db.ErrRecordNotFound {
		activePlugin.PluginID = pluginID
		activePlugin.AvatarID = avatarID
	}

	activePlugin.IsActive = active
	result = plugin.Active.DB.Save(&activePlugin)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (plugin *PluginsRepository) GetActivePlugins(avatarID uint) []models.ActivePlugin {
	query := plugin.Active.DB
	query = query.Preload("Plugin")
	query = query.Where("avatar_id = ?", avatarID)

	var activePlugins []models.ActivePlugin
	result := query.Find(&activePlugins)
	if result.Error != nil {
		return []models.ActivePlugin{}
	}

	return activePlugins
}

func (plugin *PluginsRepository) GetActivePlugin(avatarID uint, pluginID uint) (models.ActivePlugin, error) {
	query := plugin.Active.DB
	query = query.Preload("Plugin")
	query = query.Where("avatar_id = ? AND plugin_id = ?", avatarID, pluginID)

	var activePlugin models.ActivePlugin

	result := query.First(&activePlugin)
	if result.Error != nil {
		return models.ActivePlugin{}, result.Error
	}

	return activePlugin, nil
}
