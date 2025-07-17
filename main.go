package main

import (
	"PawelK2012/go-chat/src/repository"
	aichat "PawelK2012/go-chat/src/services/ai_chat"
)

func main() {
	// Instantiate clients repository to open the OpenAI connection
	repository := repository.New()

	// comment out if you would like to try Judge service
	// judge := judge.New(repository)
	// judge.JudgeLLMResult()

	aichat := aichat.New(repository)
	aichat.StartChat()

}
