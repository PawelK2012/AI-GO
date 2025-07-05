package main

import (
	"PawelK2012/go-chat/src/repository"
	"PawelK2012/go-chat/src/services/chat"
)

func main() {
	// Instantiate clients repository to open the OpenAI connection
	repository := repository.New()
	chat := chat.New(repository)
	// q := "Please propose a hard, challenging question to assess someone's IQ. Respond only with the question."
	q := "What kind of houseplant is easy to take care of?"
	chat.GetChat(q)
}
