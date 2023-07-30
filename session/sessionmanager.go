package session

import (
	"bufio"
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"gpt-asker/config"
	"log"
	"os"
	"strings"
)

type SessionManager struct {
	cli *openai.Client
	msg []openai.ChatCompletionMessage
}

func NewSessionManager(client *openai.Client) *SessionManager {
	return &SessionManager{cli: client, msg: make([]openai.ChatCompletionMessage, 0)}
}

func (manager *SessionManager) SendMessage(message string) (response string) {
	message = strings.Replace(message, "\r", "", -1)
	manager.msg = append(manager.msg, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	})
	resp, err := manager.cli.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo, //TODO: resolve parameter
			Messages: manager.msg,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	response = resp.Choices[0].Message.Content
	manager.msg = append(manager.msg, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: response,
	})
	return
}

func (manager *SessionManager) Spin() {
	cfg := config.GetConfig()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(cfg.UserName, ": ")
		input, _ := reader.ReadString('\n')
		fmt.Println("GPT :", manager.SendMessage(input))
	}
}
