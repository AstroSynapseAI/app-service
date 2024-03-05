package search

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/AstroSynapseAI/app-service/engine"
	"github.com/AstroSynapseAI/app-service/engine/tools/google"

	asaiTools "github.com/AstroSynapseAI/app-service/engine/tools"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/tools"
	"github.com/tmc/langchaingo/tools/duckduckgo"
)

var _ tools.Tool = &SearchAgent{}

type SearchAgent struct {
	Memory     schema.Memory
	Primer     string
	LLM        llms.LanguageModel
	ToolsConfg []engine.AgentToolConfig
	Executor   agents.Executor
}

func NewSearchAgent(options ...SearchAgentOptions) (*SearchAgent, error) {
	// create a new search agent
	searchAgent := &SearchAgent{
		Memory: memory.NewSimple(),
	}

	// apply search agent options
	for _, option := range options {
		option(searchAgent)
	}

	if searchAgent.LLM == nil {
		return nil, errors.New("llm is required")
	}

	// create search agent tools
	searchTools := []tools.Tool{}

	for _, tool := range searchAgent.ToolsConfg {
		if tool.GetSlug() == "google-search" && tool.IsActive() {
			google, err := google.New(tool.GetToken(), 10)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}

			searchTools = append(searchTools, google)
		}

		if tool.GetSlug() == "ddg-search" && tool.IsActive() {
			ddg, err := duckduckgo.New(10, duckduckgo.DefaultUserAgent)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			searchTools = append(searchTools, ddg)
		}
	}

	fmt.Println("Search Agent Tools:")
	for _, tool := range searchTools {
		fmt.Println(tool.Name())
	}

	// searchTmplt, err := templates.Load("search.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }

	// create search prompt template
	promptTmplt := prompts.PromptTemplate{
		Template:       searchAgent.Primer,
		TemplateFormat: prompts.TemplateFormatGoTemplate,
		InputVariables: []string{"input", "agent_scratchpad", "today"},
		PartialVariables: map[string]interface{}{
			"today":             time.Now().Format("January 02, 2006"),
			"tool_names":        asaiTools.Names(searchTools),
			"tool_descriptions": asaiTools.Descriptions(searchTools),
			"history":           "",
		},
	}

	// create the search prompt
	agent := agents.NewOneShotAgent(
		searchAgent.LLM,
		searchTools,
		agents.WithMemory(searchAgent.Memory),
		agents.WithPrompt(promptTmplt),
		agents.WithMaxIterations(3),
	)

	// create agents executor chain
	searchAgent.Executor = agents.NewExecutor(agent, searchTools)

	return searchAgent, nil
}

func (agent *SearchAgent) Call(ctx context.Context, input string) (string, error) {
	fmt.Println("Search Agent Running...")
	reponse, err := chains.Run(ctx, agent.Executor, input)
	if err != nil {
		return "Search Agent encountered an error: " + err.Error(), nil
	}
	return reponse, nil
}

func (agent *SearchAgent) Name() string {
	return "Search Agent"
}

func (agent *SearchAgent) Description() string {
	return `
		Search Agent is an agent specialized in searching the web.
		The agent can use DuckDuckGo, SerpApi Google API, and Metaphor Search 
		tools to search the web. Input should be a question or query related to Human input.
		Output should be a summary of the search with the most relevant results.
	`
}
