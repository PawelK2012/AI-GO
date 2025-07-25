package clients

import (
	"context"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/responses"
)

type ClientInterface interface {
	CompletionsNew(ctx context.Context, q, system_prompt, model string) (*openai.ChatCompletion, error)
	ResponsesNew(ctx context.Context, q, model string) (*responses.Response, error)
	StreamingNew(ctx context.Context, sysprompt, question string, tools []openai.ChatCompletionToolParam)
}
