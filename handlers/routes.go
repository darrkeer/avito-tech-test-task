package handlers

import (
	"net/http"

	"github.com/darrkeer/avito-tech-test-task/repository"
)

type Handler struct {
	repository *repository.Repository
}

func New(r *repository.Repository) *Handler {
	return &Handler{
		repository: r,
	}
}

func (h *Handler) Start() {
	http.HandleFunc("/users/new-dummy", h.RegisterUser)
}
