package generator

import "github.com/google/uuid"

type gen struct{}

func New() *gen {
	return &gen{}
}

func (g *gen) Generate() string {
	return uuid.NewString()
}
