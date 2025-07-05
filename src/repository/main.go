package repository

import "PawelK2012/go-chat/src/clients"

type Repository struct {
	OAI clients.ClientInterface
}

func New() *Repository {
	openaiclient := clients.NewOAIClient()
	return &Repository{OAI: openaiclient}
}
