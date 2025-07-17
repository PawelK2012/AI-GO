package aichat

import (
	"PawelK2012/go-chat/src/repository"
	"fmt"
)

type AIChat struct {
	repository *repository.Repository
}

func New(repository *repository.Repository) *AIChat {
	aichat := &AIChat{
		repository: repository,
	}
	fmt.Println("initialising ai chat service")
	return aichat
}
