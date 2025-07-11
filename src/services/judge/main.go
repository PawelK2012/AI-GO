package judge

import (
	"PawelK2012/go-chat/src/repository"
	"fmt"
)

type Judge struct {
	repository *repository.Repository
}

func New(repository *repository.Repository) *Judge {
	judge := &Judge{
		repository: repository,
	}
	fmt.Println("initialising chat service")
	return judge
}
