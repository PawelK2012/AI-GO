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
func (c *OAIClient) CompletionsNew(ctx context.Context, q, system_prompt, model string) (*openai.ChatCompletion, error) {

	//messages = [{"role": "system", "content": system_prompt}] + history + [{"role": "user", "content": message}]
	param := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(q),
			//openai.SystemMessage(system_prompt),
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

// just experimenting with NewStreaming API
func (c *OAIClient) StreamingNew(ctx context.Context, sysprompt, question string, tools []openai.ChatCompletionToolParam) {

	messages := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage(sysprompt),
		openai.UserMessage(question),
	}

	params := openai.ChatCompletionNewParams{
		Messages: messages,
		Seed:     openai.Int(0),
		Model:    "llama3.2",
		Tools:    tools,
	}
	stream := c.oai_client.Chat.Completions.NewStreaming(ctx, params)
	acc := openai.ChatCompletionAccumulator{}

	for stream.Next() {
		chunk := stream.Current()

		acc.AddChunk(chunk)

		// When this fires, the current chunk value will not contain content data
		if _, ok := acc.JustFinishedContent(); ok {
			println()
			println("finish-event: Content stream finished")
		}

		if refusal, ok := acc.JustFinishedRefusal(); ok {
			println()
			println("finish-event: refusal stream finished:", refusal)
			println()
		}

		if tool, ok := acc.JustFinishedToolCall(); ok {
			println("finish-event: tool call stream finished:", tool.Index, tool.Name, tool.Arguments)
		}

		// It's best to use chunks after handling JustFinished events.
		// Here we print the delta of the content, if it exists.
		if len(chunk.Choices) > 0 && chunk.Choices[0].Delta.Content != "" {
			print(chunk.Choices[0].Delta.Content)
		}
	}

	if err := stream.Err(); err != nil {
		panic(err)
	}

	if acc.Usage.TotalTokens > 0 {
		println("Total Tokens:", acc.Usage.TotalTokens)
	}
}
