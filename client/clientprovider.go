package client

import (
	"github.com/sashabaranov/go-openai"
	"gpt-asker/config"
	"gpt-asker/session"
)

type ClientProvider struct {
	cli     *openai.Client
	manager *session.SessionManager
}

var clientProvider *ClientProvider

func getClientProvider() *ClientProvider {
	if clientProvider == nil {
		clientProvider = &ClientProvider{}
		clientProvider.cli = new(openai.Client)
		clientProvider.manager = new(session.SessionManager)
	}
	return clientProvider
}

func (p *ClientProvider) createClient() *openai.Client {
	cfg := config.GetConfig()
	p.cli = openai.NewClient(cfg.APIKey)
	return p.cli
}

func NewSession() *session.SessionManager {
	provider := getClientProvider()
	provider.createClient()

	provider.manager = session.NewSessionManager(provider.cli)
	return provider.manager
}
