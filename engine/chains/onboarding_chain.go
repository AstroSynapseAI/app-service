package chains

import (
	"context"
	"fmt"
	"os"

	"github.com/AstroSynapseAI/app-service/engine"
	"github.com/AstroSynapseAI/app-service/engine/memory"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/prompts"
	"github.com/tmc/langchaingo/schema"
)

const (
	PromptTemplate = `
		Your task is to monitor onboarding conversations between Asai and a new users. You are given the conversation history, onboarding script and latest human input. 

		Based on the conversation history, onboarding script and latest human input you need to decide: 
		- should the Asai begin the onbarding process of the new user?
		- should Asai continue with the onboarding process?
		- should Asai stop the onboarding process?		
		
		Onboarding process should start: 
		- if the conversation history is empty
		- if user explictly requested it

		Onboarding process should not start:
		- if the conversation history is NOT empty

		Onboarding process should end:
		- if all steps of the onboarding scripts are concluded
		- if user explicitly requested it

		Onboarding process should continue:
		- until the AI hasn't categorized the new user 
		- if user is inquisitive about the Astro Synapse company or ASAI

		Your response will always be: 
		Yes - if the onboarding process should begin
		Yes - if the onboarding process should continue
		No - if the onboarding process should stop

		ANSWER ONLY WITH "Yes" OR "No"!

		BEGIN!

		CONVERSATION HISTORY:
		{{.history}}

		ONBOARDING SCRIPT:
		{{.script}}

		USER INPUT: {{.input}}
	`
)

type OnboardingChain struct {
	Chain  *chains.LLMChain
	memory *memory.AsaiMemory
}

func NewOnboardingChain(config engine.AvatarConfig, memory *memory.AsaiMemory) (*OnboardingChain, error) {
	prompt := prompts.NewPromptTemplate(PromptTemplate, []string{"history", "input", "script"})

	if config == nil {
		return nil, fmt.Errorf("config cannot be nil")
	}

	chain := chains.NewLLMChain(config.GetAvatarLLM(), prompt)

	onbaording := &OnboardingChain{
		Chain:  chain,
		memory: memory,
	}

	return onbaording, nil
}

func (onboarding *OnboardingChain) Call(ctx context.Context, input string) (string, error) {
	inputValues := map[string]any{}
	inputValues["input"] = input

	history, err := schema.GetBufferString(onboarding.memory.Messages(), "Human", "Asai")
	if err != nil {
		return "", err
	}
	inputValues["history"] = history

	tmplContent, err := os.ReadFile("./engine/documents/onboarding_script.txt")
	if err != nil {
		fmt.Println("Error reading onboarding script:", err)
		return "", err
	}

	script := string(tmplContent)
	inputValues["script"] = script

	response, err := chains.Call(
		ctx,
		onboarding.Chain,
		inputValues,
		chains.WithTemperature(0.1),
	)
	if err != nil {
		return "", err
	}

	answer := response["text"].(string)
	return answer, nil
}
