package handler

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"gpt-asker/client"
)

func asker() {
	resp, err := client.GetClient().CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello! I'm using go API to chat with you",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
