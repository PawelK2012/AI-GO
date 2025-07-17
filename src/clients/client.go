package clients

import (
	"context"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/responses"
)

type ClientInterface interface {
	CompletionsNew(ctx context.Context, q, model string) (*openai.ChatCompletion, error)
	ResponsesNew(ctx context.Context, q, model string) (*responses.Response, error)
}
