package clients

import "github.com/openai/openai-go"

type ClientInterface interface {
	Ask(q string) (*openai.ChatCompletion, error)
}
