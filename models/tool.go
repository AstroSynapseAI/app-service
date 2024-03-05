package models

import (
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"gorm.io/gorm"
)

type Tool struct {
	gorm.Model
	Name             string            `json:"name,omitempty"`
	Description      string            `json:"description,omitempty"`
	Slug             string            `json:"slug,omitempty"`
	Token            string            `json:"token,omitempty"`
	ActiveTools      []ActiveTool      `gorm:"foreignKey:ToolID;" json:"active_tools"`
	ActiveAgentTools []ActiveAgentTool `gorm:"foreignKey:ToolID;" json:"agent_tools"`
}

func (*Tool) SeedModel(db *database.Database) error {
	seeder := "seed_tools"

	result := db.Adapter.Gorm().Where("seeder_name = ?", seeder).First(&DBSeeder{})
	if result.Error == gorm.ErrRecordNotFound {
		var tools []Tool = []Tool{
			{
				Name:        "Google Search",
				Slug:        "google-search",
				Token:       "",
				Description: "Google Search via SerpAPI, SerpAPI token required.",
			},
			{
				Name:        "DuckDuckGo Search",
				Slug:        "ddg-search",
				Token:       "",
				Description: "DuckDuckGo Search API, Free Service.",
			},
			{
				Name:        "Metaphor Search",
				Slug:        "metaphor-search",
				Token:       "",
				Description: "Metaphor Search API, API key required.",
			},
		}

		if result := db.Adapter.Gorm().Create(&tools); result.Error != nil {
			return result.Error
		}

		if result := db.Adapter.Gorm().Create(&DBSeeder{SeederName: seeder}); result.Error != nil {
			return result.Error
		}
	}

	return nil
}
