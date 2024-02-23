package engine

import (
	"fmt"

	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/repositories"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

type Config struct {
	DB     *database.Database
	Avatar *models.Avatar
}

var _ AvatarConfig = (*Config)(nil)

func NewConfig(db *database.Database) *Config {
	return &Config{
		DB: db,
	}
}

func (cnf *Config) LoadConfig(avatarID uint) {
	avatars := repositories.NewAvatarsRepository(cnf.DB)

	avatar, err := avatars.Fetch(avatarID)
	if err != nil {
		fmt.Println("Error loading avatar:", err)
		return
	}

	cnf.Avatar = &avatar
}

func (cnf *Config) GetDB() *database.Database {
	return cnf.DB
}

func (cnf *Config) GetAvatarLLM() llms.LanguageModel {
	avatarLLM := cnf.Avatar.LLM

	switch avatarLLM.Slug {
	case "gpt-4":
		LLM, err := openai.NewChat(openai.WithModel("gpt-4"))
		if err != nil {
			fmt.Println("Error setting gpt-4:", err)
			return nil
		}
		return LLM
	default:
		fmt.Println("Unknown LLM:", avatarLLM.Slug)
		return nil
	}
}

func (cnf *Config) GetAvatarName() string {
	return cnf.Avatar.Name
}

func (cnf *Config) GetAvatarPrimer() string {
	return cnf.Avatar.Primer
}

func (cnf *Config) GetAvatarMemorySize() int {
	return 4048
}

func (cnf *Config) AvatarIsPublic() bool {
	return cnf.Avatar.IsPublic
}

func (cnf *Config) GetLLM() LLMConfig {
	return nil
}

func (cnf *Config) GetAgents() []AgentConfig {
	return nil
}

func (cnf *Config) GetTools() []ToolConfig {
	return nil
}

func (cnf *Config) GetPlugins() []PluginConfig {
	return nil
}
