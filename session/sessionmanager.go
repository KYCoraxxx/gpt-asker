package session

import (
	"bufio"
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"gpt-asker/client"
	"gpt-asker/config"
	"log"
	"os"
	"strings"
)

type SessionManager struct {
	cli *openai.Client
	msg []openai.ChatCompletionMessage
}

var sessionManager *SessionManager

func getSessionManager() *SessionManager {
	if sessionManager == nil {
		sessionManager = new(SessionManager)
		sessionManager.cli = client.GetClient()
		sessionManager.msg = make([]openai.ChatCompletionMessage, 0)
	}
	return sessionManager
}

func SendMessage(message string) (response string) {
	message = strings.Replace(message, "\r", "", -1)
	manager := getSessionManager()
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

func Spin() {
	cfg := config.GetConfig()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(cfg.UserName, ": ")
		input, _ := reader.ReadString('\n')
		fmt.Println("GPT :", SendMessage(input))
	}
}
