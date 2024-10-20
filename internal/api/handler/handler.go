package handler

import "ddos/internal/generator"

type Handler struct {
	Generator *generator.Generator
}

func NewHandler(gene *generator.Generator) *Handler {
	return &Handler{}
}
