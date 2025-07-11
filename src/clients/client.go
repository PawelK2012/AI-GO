package clients

import (
	"context"

	"github.com/openai/openai-go"
)

type ClientInterface interface {
	Ask(ctx context.Context, q, model string) (*openai.ChatCompletion, error)
}
