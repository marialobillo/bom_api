package handler

import (
	"github.com/marialobillo/bom_api/internal/service"

)

type PartHandler struct {
	partService *service.Createpart
}

func NewPartHandler(partService *service.CreatePart) *PartHandler {
	return &PartHandler{partService: partService}
}

