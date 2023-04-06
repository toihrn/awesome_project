package ai

import "github.com/sashabaranov/go-openai"

var OpenAIClient *openai.Client

const (
	openaiAPIKey = "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"
)

func InitOpenAIClient() {
	OpenAIClient = openai.NewClient(openaiAPIKey)
}
