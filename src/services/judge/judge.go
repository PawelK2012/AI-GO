package judge

import (
	"context"
	"fmt"
)

// Judge is a service that takes a user task and sends it to multiple large language models (LLMs) to compare and evaluate their responses.
//
// The goal is to improve the overall quality and reliability of LLM-generated output
func (ch *Judge) JudgeLLMResult() {
	ctx := context.Background()
	var competitors []string
	competitors = append(competitors, "llama3.2", "granite3.3:2b", "gemma3:1b")

	var ans []string

	sys_prompt := "Answer only with the question, no explanation."

	q := "Please come up with a challenging, nuanced question that I can ask a number of LLMs to evaluate their intelligence."

	fmt.Printf("question1: %s\n", q)

	for i, comp := range competitors {
		answer, err := ch.repository.OAI.CompletionsNew(ctx, q, sys_prompt, comp)

		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("\ncompetitor: %s\nanswer:  %s\n", comp, answer.Choices[0].Message.Content)
		fmt.Print("================================ \n")

		a := fmt.Sprintf("# Response from competitor %d \n %s \n", i+1, answer.Choices[0].Message.Content)
		ans = append(ans, a)
	}
	res := ch.Judge(ctx, ans, competitors, q)
	println("judge result: \n", res)
	fmt.Print("================================ \n")

}

// The Judge() evaluates and ranks the results returned by the JudgeLLMResult()
func (ch *Judge) Judge(ctx context.Context, ans, competitors []string, question string) string {
	sys_prompt := "Keep official tone."
	judge := fmt.Sprintf("You are judging a competition between %d competitors. Each model has been given this question: %s Your job is to evaluate each response for clarity and strength of argument, and rank them in order of best to worst. Respond with JSON, and only JSON, with the following format: {{`results`: [`best competitor number`, `second best competitor number`, `third best competitor number`, ...]}} Here are the responses from each competitor: %s Now respond with the JSON with the ranked order of the competitors, nothing else. Do not include markdown formatting or code blocks.", len(competitors), question, ans)

	println("judge\n", judge)
	fmt.Print("================================ \n")
	answer, err := ch.repository.OAI.CompletionsNew(ctx, judge, sys_prompt, "llama3.2")
	if err != nil {
		fmt.Println(err.Error())
	}
	return answer.Choices[0].Message.Content

}
