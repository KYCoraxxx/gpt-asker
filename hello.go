package main

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	yaml "gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type config struct {
	APIKey string `yaml:"APIKey"`
}

func getConfig() config {
	var cfg config

	content, err := ioutil.ReadFile("./application.yaml")
	if err != nil {
		log.Fatal(err)
	}

	if err = yaml.Unmarshal(content, &cfg); err != nil {
		log.Fatal(err)
	}

	return cfg
}

func main() {
	cfg := getConfig()

	client := openai.NewClient(cfg.APIKey)
	resp, err := client.CreateChatCompletion(
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
