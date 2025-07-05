package chat

import (
	"PawelK2012/go-chat/src/repository"
	"fmt"
)

type Chat struct {
	repository *repository.Repository
}

func New(repository *repository.Repository) *Chat {
	chat := &Chat{
		repository: repository,
	}
	fmt.Println("initialising chat service")
	return chat
}
