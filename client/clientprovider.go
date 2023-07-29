package client

import (
	"github.com/sashabaranov/go-openai"
	"gpt-asker/config"
)

type ClientProvider struct {
	cli *openai.Client
}

var clientProvider *ClientProvider

func getClientProvider() *ClientProvider {
	if clientProvider == nil {
		clientProvider = &ClientProvider{}
		clientProvider.cli = new(openai.Client)
	}
	return clientProvider
}

func (p *ClientProvider) createClient() *openai.Client {
	cfg := config.GetConfig()
	p.cli = openai.NewClient(cfg.APIKey)
	return p.cli
}

func GetClient() *openai.Client {
	provider := getClientProvider()
	return provider.createClient()
}
