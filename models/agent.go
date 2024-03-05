package models

import (
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"gorm.io/gorm"
	"regexp"
)

type Agent struct {
	gorm.Model
	Name         string        `json:"name,omitempty"`
	Description  string        `json:"description,omitempty"`
	Slug         string        `json:"slug,omitempty"`
	Primer       string        `json:"primer,omitempty"`
	ActiveAgents []ActiveAgent `gorm:"foreignKey:AgentID;" json:"active_agents"`
}

func (*Agent) SeedModel(db *database.Database) error {
	seeder := "seed_agents"
	rex := regexp.MustCompile(`(?m)^\s+`)
	result := db.Adapter.Gorm().Where("seeder_name = ?", seeder).First(&DBSeeder{})
	if result.Error == gorm.ErrRecordNotFound {
		var agents []Agent = []Agent{
			{
				Name:        "Search Agent",
				Slug:        "search-agent",
				Description: "Utilizes search engines such as Google, DuckDuckGo, and Metaphor for automated web searches.",
				Primer: rex.ReplaceAllString(`
				CURRENT DATE: {{.today}}

				Search Assistant is trained to search the web based on user input and conversation history using the following tools:
				
				{{.tool_descriptions}}
				
				Based on the user input and conversation history, it decides whether an internet search should be performed.
				
				Use the following format:
				
				Question: The input question you must answer 
				Thought: You should always think about what to do 
				Action: The action to take, should be one of [ {{.tool_names}} ] 
				Action Input: The input to the action 
				Observation: The result of the action 
				...(This Thought/Action/Action Input/Observation can repeat N times) 
				Thought: I now know the final answer
				Final Answer: 
					Summary: [ Summarize the final answer here ]
					Most Relevant Links:
					- Link 1: Description of Link 1
					- Link N: Description of Link N
					...(Depending on relevance you can add none, or N number of Links)
				
				Begin!
				
				Conversation History: {{.history}}
				
				User input: {{.input}}
				Thought:{{.agent_scratchpad}}
				`, ""),
			},
			{
				Name:        "Browser Agent",
				Slug:        "browser-agent",
				Description: "Equipped with the capability to scrape, read website contents, and interact with web pages and web applications.",
				Primer: rex.ReplaceAllString(`
				Please write a detailed report of the following website and its pages that will not exceed 4048 tokens:
		
				"{{.context}}"
		
				If query is provided, focus on the content related to the query. 
		
				Query: {{.query}}
		
				Structure the report in the following format:
		
				WEBSITE SUMMARY:
				[Place the summary of the entire website here]
		
				PAGE SUMMARIES:
				- [Page 1 Title]: [Summary of Page 1]
				- [Page N Title]: [Summary of Page N]
				...(Create a summary for every sub-page on the website)
		
				LINK INDEX:
				- Link 1: [Description of Link 1]
				- Link N: [Description of Link N]
				...(Depending on relevance, you can add none or N number of links)
		
				FINAL THOUGHTS:
				[Place any final thoughts or a concluding summary here]`, ""),
			},
		}

		if result := db.Adapter.Gorm().Create(&agents); result.Error != nil {
			return result.Error
		}

		if result := db.Adapter.Gorm().Create(&DBSeeder{SeederName: seeder}); result.Error != nil {
			return result.Error
		}
	}

	return nil
}
