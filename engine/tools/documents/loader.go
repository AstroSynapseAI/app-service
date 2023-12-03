package documents

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/tmc/langchaingo/tools"
)

const (
	DefualtRootPath = "./engine/documents/"

	ErrInvalidInput = "Invalid input"
	ErrFileLoad     = "Error while loading requested file"
)

var _ tools.Tool = &DocumentsTool{}

type DocumentsTool struct {
	RootPath string
}

func NewLoader(options ...DocuemntsToolOption) (*DocumentsTool, error) {
	documents := &DocumentsTool{
		RootPath: DefualtRootPath,
	}

	for _, option := range options {
		option(documents)
	}

	return documents, nil
}

func (agent *DocumentsTool) Name() string {
	return "Library Agent"
}

func (agent *DocumentsTool) Description() string {
	return `
		Library Agent is an agent specialized in fetching and reading documents from your internal knowledge library. It helps you access relevant information to provide accurate and informative responses.
		
		**Usage Example**:
		1. Identify the appropriate document from the library list based on the user's query.
		2. Use the "FileName" parameter to request the document in the JSON format:
		{
			"FileName": "filename.txt"
		}		
		3. Receive the content of the file as output.
		4. Integrate the content into your response in a coherent and natural way.
		
		**Available Documents**:
		- Name: Astro Synapse Impressum
		  - FileName: astrosynapse_impressum.txt
		  - Description: General information about Astro Synapse. Use to answer any information related to Astro Synapse, such as the company's mission, vision, website, etc.
		- Name: ASAI Impressum
		  - FileName: asai_impressum.txt
		  - Description: ASAI Architecture describes the code and technologies used to build and run ASAI. Use this document to understand how ASAI works and to answer any technical questions about ASAI.
		- Name: Onboarding Script
		  - FileName: onboarding_script.txt
		  - Description: Welcome script for new users. Use this document when you are prompted to welcome or onboard a new user to our system.
		
		**Error Handling**:
		If you encounter an error or cannot fetch the requested document, inform the user and offer alternative assistance or information.
	`
}

func (tool *DocumentsTool) Call(ctx context.Context, input string) (string, error) {
	var toolInput struct {
		FileName string
		Query    string
	}

	re := regexp.MustCompile(`(?s)\{.*\}`)
	jsonString := re.FindString(input)

	err := json.Unmarshal([]byte(jsonString), &toolInput)
	if err != nil {
		return fmt.Sprintf("%v: %s", ErrInvalidInput, err), nil
	}

	path := tool.RootPath + toolInput.FileName

	fileContents, err := os.ReadFile(path)
	if err != nil {
		return fmt.Sprintf("%v: %s", ErrFileLoad, err), nil
	}

	return string(fileContents), nil
}
