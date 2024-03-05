package engine

import (
	"errors"
	"fmt"

	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/repositories"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/llms/openai"
)

type Config struct {
	DB     *database.Database
	Avatar models.Avatar
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

	cnf.Avatar = avatar
}

func (cnf *Config) GetDB() *database.Database {
	return cnf.DB
}

func (cnf *Config) GetAvatarLLM() llms.LanguageModel {
	avatarLLM := cnf.Avatar.LLM
	activeLLMs := cnf.Avatar.ActiveLLMs
	if len(activeLLMs) == 0 {
		fmt.Println("No active LLMs")
		return nil
	}

	var activeLLM models.ActiveLLM
	// extract active llm where activeLLM.llmID == avatarLLM.ID
	for _, active := range activeLLMs {
		if active.LLM.ID == avatarLLM.ID {
			activeLLM = active
		}
	}

	llm, err := loadActiveLLM(activeLLM)
	if err != nil {
		fmt.Println("Error loading LLM:", err)
		return nil
	}
	return llm
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

func (cnf *Config) GetAgents() []AgentConfig {
	activeAgents := cnf.Avatar.ActiveAgents
	var configs []AgentConfig

	for _, activeAgent := range activeAgents {
		configs = append(configs, NewActiveAgent(cnf.Avatar, activeAgent))
	}
	return configs
}

func (cnf *Config) GetTools() []ToolConfig {
	activeTools := cnf.Avatar.ActiveTools
	var configs []ToolConfig
	for _, activeTool := range activeTools {
		configs = append(configs, NewActiveTool(activeTool))
	}
	return configs
}

// Active Agent Config
type ActiveAgent struct {
	Avatar      models.Avatar
	ActiveAgent models.ActiveAgent
}

var _ AgentConfig = (*ActiveAgent)(nil)

func NewActiveAgent(avatar models.Avatar, activeAgent models.ActiveAgent) *ActiveAgent {
	return &ActiveAgent{
		Avatar:      avatar,
		ActiveAgent: activeAgent,
	}
}

func (cnf *ActiveAgent) GetAgentName() string {
	return cnf.ActiveAgent.Agent.Name
}

func (cnf *ActiveAgent) GetAgentSlug() string {
	return cnf.ActiveAgent.Agent.Slug
}

func (cnf *ActiveAgent) GetAgentLLM() llms.LanguageModel {
	agentLLM := cnf.ActiveAgent.LLM
	activeLLMs := cnf.Avatar.ActiveLLMs

	var activeLLM models.ActiveLLM

	for _, active := range activeLLMs {
		if active.LLM.ID == agentLLM.ID {
			activeLLM = active
		}
	}

	llm, err := loadActiveLLM(activeLLM)
	if err != nil {
		fmt.Println("Error loading LLM:", err)
		return nil
	}
	return llm
}

func (cnf *ActiveAgent) GetAgentPrimer() string {
	return cnf.ActiveAgent.Primer
}

func (cnf *ActiveAgent) IsAgentPublic() bool {
	return cnf.ActiveAgent.IsPublic
}

func (cnf *ActiveAgent) IsAgentActive() bool {
	return cnf.ActiveAgent.IsActive
}

func (cnf *ActiveAgent) GetAgentTools() []AgentToolConfig {
	activeTools := cnf.ActiveAgent.ActiveAgentTools
	var configs []AgentToolConfig
	for _, activeTool := range activeTools {
		configs = append(configs, NewActiveAgentTool(activeTool))
	}
	return configs
}

// Active Tool Config
type ActiveTool struct {
	activeTool models.ActiveTool
}

var _ ToolConfig = (*ActiveTool)(nil)

func NewActiveTool(tool models.ActiveTool) *ActiveTool {
	return &ActiveTool{
		activeTool: tool,
	}
}

func (cnf *ActiveTool) GetName() string {
	return cnf.activeTool.Tool.Name
}

func (cnf *ActiveTool) GetSlug() string {
	return cnf.activeTool.Tool.Slug
}

func (cnf *ActiveTool) GetToken() string {
	return cnf.activeTool.Token
}

func (cnf *ActiveTool) IsPublic() bool {
	return cnf.activeTool.IsPublic
}

func (cnf *ActiveTool) IsActive() bool {
	return cnf.activeTool.IsActive
}

type ActiveAgentTool struct {
	activeAgentTool models.ActiveAgentTool
}

var _ AgentToolConfig = (*ActiveAgentTool)(nil)

func NewActiveAgentTool(agentTool models.ActiveAgentTool) *ActiveAgentTool {
	return &ActiveAgentTool{
		activeAgentTool: agentTool,
	}
}

func (cnf *ActiveAgentTool) GetName() string {
	return cnf.activeAgentTool.Tool.Name
}

func (cnf *ActiveAgentTool) GetSlug() string {
	return cnf.activeAgentTool.Tool.Slug
}

func (cnf *ActiveAgentTool) GetToken() string {
	return cnf.activeAgentTool.Token
}

func (cnf *ActiveAgentTool) IsPublic() bool {
	return cnf.activeAgentTool.IsPublic
}

func (cnf *ActiveAgentTool) IsActive() bool {
	return cnf.activeAgentTool.IsActive
}

func loadActiveLLM(activeLLM models.ActiveLLM) (llms.LanguageModel, error) {

	switch activeLLM.LLM.Slug {
	case "mistral":
		LLM, err := ollama.New(
			ollama.WithModel("mistral"),
			ollama.WithServerURL("http://host.docker.internal:11434/"),
		)

		if err != nil {
			fmt.Println("Error setting mistral:", err)
			return nil, err
		}

		return LLM, nil
	case "gpt-4":
		LLM, err := openai.NewChat(
			openai.WithToken(activeLLM.Token),
			openai.WithModel("gpt-4"),
		)

		if err != nil {
			fmt.Println("Error setting gpt-4:", err)
			return nil, err
		}

		return LLM, nil
	default:
		fmt.Println("Unknown LLM:", activeLLM.LLM.Slug)
		return nil, errors.New("unknown LLM")
	}
}
