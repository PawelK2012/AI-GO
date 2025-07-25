package aichat

import (
	"context"
	"fmt"
	"os"
	"strings"
)

func (c *AIChat) AskQuestion(user_question string) (string, error) {
	fmt.Print("starting chat...")
	ctx := context.Background()

	//override with youe name
	name := "Pawel Kaim"

	// reading files content
	cv := getFileData("src/services/ai_chat/files/cv.txt")
	summary := getFileData("src/services/ai_chat/files/summary.txt")

	// building prompt
	var system_prompr strings.Builder
	prompt := fmt.Sprintf("You are acting as %s. You are answering questions on %s 's website, particularly questions related to %s's career, background, skills and experience. Your responsibility is to represent %s for interactions on the website as faithfully as possible. You are given a summary of %s's background and LinkedIn profile which you can use to answer questions. Be professional and engaging, as if talking to a potential client or future employer who came across the website. If you don't know the answer, say so. ", name, name, name, name, name)
	system_prompr.WriteString(prompt)

	formatted_summary := fmt.Sprintf("\n\n## Summary:\n %s \n\n## LinkedIn Profile:\n %s \n\n", summary, cv)
	system_prompr.WriteString(formatted_summary)

	prompt_ctx := fmt.Sprintf("With this context, please chat with the user, always staying in character as %s.", name)
	system_prompr.WriteString(prompt_ctx)

	//fmt.Print(system_prompr.String())

	// currently Ollam is not supporting OpenAI Responses API https://github.com/ollama/ollama/issues/9659
	//r, err := c.repository.OAI.ResponsesNew(ctx, system_prompr.String(), "llama3.2")
	r, err := c.repository.OAI.CompletionsNew(ctx, user_question, system_prompr.String(), "llama3.2")
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	// fmt.Print("====== RESPONSE ====")
	// fmt.Print(r.Choices[0].Message.Content)
	return r.Choices[0].Message.Content, nil

}

func getFileData(p string) string {
	data, err := os.ReadFile(p)
	check(err)
	return string(data)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
