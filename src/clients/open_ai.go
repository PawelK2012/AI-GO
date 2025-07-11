package clients

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
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

func (c *OAIClient) Ask(ctx context.Context, q, model string) (*openai.ChatCompletion, error) {

	param := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(q),
		},
		Seed:  openai.Int(1),
		Model: model,
	}

	completion, err := c.oai_client.Chat.Completions.New(context.TODO(), param)
	if err != nil {
		fmt.Print("completion failed \n", err.Error())
	}

	return completion, err
}
