package google

import (
	"context"
	"errors"
	"fmt"

	"github.com/tmc/langchaingo/tools"
)

type Tool struct {
	client *Client
}

var _ tools.Tool = &Tool{}

func New(apiKey string, maxResults int) (*Tool, error) {
	fmt.Println("Google Search SerpAPI Tool Created...")
	return &Tool{
		client: NewClient(apiKey, maxResults),
	}, nil
}

func (tool Tool) Name() string {
	return "Google Search SerpAPI Tool"
}

func (tool Tool) Description() string {
	return `
	A  wrapper around Google Search SerpAPI Tool.
	Will return a structured list of search results.
	Input should be a search query.`
}

func (tool Tool) Call(ctx context.Context, input string) (string, error) {
	fmt.Println("Google Search SerpAPI Tool Running...")
	result, err := tool.client.Search(ctx, input)
	if err != nil {
		if errors.Is(err, ErrNoGoodResult) {
			return "No good Google search results was found", nil
		}

		if errors.Is(err, ErrAPIError) {
			return "Google SerpAPI responded with an error", nil
		}

		return "", err
	}

	return result, nil

}
