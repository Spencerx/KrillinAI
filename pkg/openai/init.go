package openai

import (
	"github.com/sashabaranov/go-openai"
	"krillin-ai/config"
	"net/http"
)

type Client struct {
	client *openai.Client
}

func NewClient(baseUrl, apiKey, proxyAddr string) *Client {
	cfg := openai.DefaultConfig(apiKey)
	if baseUrl != "" {
		cfg.BaseURL = baseUrl
	}

	if proxyAddr != "" {
		transport := &http.Transport{
			Proxy: http.ProxyURL(config.Conf.App.ParsedProxy),
		}
		cfg.HTTPClient = &http.Client{
			Transport: transport,
		}
	}

	client := openai.NewClientWithConfig(cfg)
	return &Client{client: client}
}
