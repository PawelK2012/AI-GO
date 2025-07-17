package clients

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/responses"
)

type OAIClient struct {
	oai_client openai.Client
}

func NewOAIClient() ClientInterface {
	fmt.Print("initialising OpenAI client \n")
	c := openai.NewClient(
		option.WithAPIKey("ollama"),
		option.WithBaseURL("http://localhost:11434/v1"),
	)
	return &OAIClient{
		oai_client: c,
	}
}

// wrapper for OpenAI client .Completions.New()
func (c *OAIClient) CompletionsNew(ctx context.Context, q, model string) (*openai.ChatCompletion, error) {

	param := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(q),
		},
		Seed:  openai.Int(1),
		Model: model,
	}

	completion, err := c.oai_client.Chat.Completions.New(context.TODO(), param)
	if err != nil {
		_, e := fmt.Print("completion failed:", err.Error())
		return nil, e
	}

	return completion, nil
}

// wrapper for OpenAI client .Responses.New()
// currently Ollam is not supporting OpenAI Responses API https://github.com/ollama/ollama/issues/9659
func (c *OAIClient) ResponsesNew(ctx context.Context, q, model string) (*responses.Response, error) {
	params := responses.ResponseNewParams{
		Model:           model,
		Temperature:     openai.Float(0.7),
		MaxOutputTokens: openai.Int(512),
		Input: responses.ResponseNewParamsInputUnion{
			OfString: openai.String(q),
		},
	}

	resp, err := c.oai_client.Responses.New(ctx, params)
	if err != nil {
		_, e := fmt.Print("responses failed:", err.Error())
		return nil, e
	}
	return resp, nil

}
