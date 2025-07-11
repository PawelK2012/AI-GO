package main

import (
	"PawelK2012/go-chat/src/repository"
	"PawelK2012/go-chat/src/services/judge"
)

func main() {
	// Instantiate clients repository to open the OpenAI connection
	repository := repository.New()
	judge := judge.New(repository)

	judge.JudgeLLMResult()
}
