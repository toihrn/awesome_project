package application

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"github.com/toxb11/awesome_project/infra/ai"
	"github.com/toxb11/awesome_project/vo"
	"log"
	"strconv"
)

func CompleteChat(ctx context.Context, req *vo.CompleteChatRequest) (*vo.CompleteChatResponse, error) {
	chatCompletionResponse, err := ai.OpenAIClient.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: req.Message,
			},
		},
		User: strconv.Itoa(int(req.UserId)),
	})
	if err != nil {
		log.Printf("[CompleteChat] err: %v\n", err)
		return nil, err
	}
	log.Printf("[CompleteChat] resp from openai: %v", chatCompletionResponse)
	var ansMsg string
	if len(chatCompletionResponse.Choices) > 0 {
		ansMsg = chatCompletionResponse.Choices[0].Message.Content
	}
	resp := &vo.CompleteChatResponse{
		Answer: ansMsg,
	}
	return resp, nil
}

//func CompleteChatStream(ctx context.Context, req *vo.CompleteChatRequest) (*vo.CompleteChatResponse, error) {
//	chatCompletionStream, err := ai.OpenAIClient.CreateChatCompletionStream(ctx, openai.ChatCompletionRequest{
//		Model: openai.GPT3Dot5Turbo,
//		Messages: []openai.ChatCompletionMessage{
//			{
//				Role:    openai.ChatMessageRoleUser,
//				Content: req.Message,
//			},
//		},
//		User: strconv.Itoa(int(req.UserId)),
//	})
//	if err != nil {
//		return nil, err
//	}
//
//}
