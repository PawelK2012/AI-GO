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
	c := openai.NewClient(
		option.WithAPIKey("ollama"),
		option.WithBaseURL("http://localhost:11434/v1"),
	)
	return &OAIClient{
		oai_client: c,
	}
}

func (c *OAIClient) Ask(q string) (*openai.ChatCompletion, error) {

	param := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			// openai.UserMessage("Please come up with a challenging, nuanced question that I can ask a number of LLMs to evaluate their intelligence. Answer only with the question, no explanation."),
			openai.UserMessage(q),
		},
		Seed:  openai.Int(1),
		Model: "llama3.2",
	}

	completion, err := c.oai_client.Chat.Completions.New(context.TODO(), param)
	if err != nil {
		fmt.Print("initialising OpenAI clien failed \n")
		panic(err.Error())
	}

	param.Messages = append(param.Messages, completion.Choices[0].Message.ToParam())
	param.Messages = append(param.Messages, openai.UserMessage("How big are those?"))

	completion, err = c.oai_client.Chat.Completions.New(context.TODO(), param)
	if err != nil {
		panic(err.Error())
	}

	return completion, err
}
