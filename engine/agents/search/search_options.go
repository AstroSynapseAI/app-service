package search

import (
	"github.com/AstroSynapseAI/app-service/engine"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/schema"
)

type SearchAgentOptions func(agent *SearchAgent)

func WithMemory(memory schema.Memory) SearchAgentOptions {
	return func(agent *SearchAgent) {
		agent.Memory = memory
	}
}

func WithPrimer(primer string) SearchAgentOptions {
	return func(agent *SearchAgent) {
		agent.Primer = primer
	}
}

func WithLLM(llm llms.LanguageModel) SearchAgentOptions {
	return func(agent *SearchAgent) {
		agent.LLM = llm
	}
}

func WithToolsConfig(tools []engine.AgentToolConfig) SearchAgentOptions {
	return func(agent *SearchAgent) {
		agent.ToolsConfg = tools
	}
}
