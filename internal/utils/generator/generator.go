package generator

import "github.com/google/uuid"

type Gen struct{}

func New() *Gen {
	return &Gen{}
}

func (g *Gen) Generate() string {
	return uuid.NewString()
}
