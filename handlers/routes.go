package handlers

import (
	"encoding/json"
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
	http.HandleFunc("/team/add", h.TeamAdd)
}

func WriteResponse(w http.ResponseWriter, resp *map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func WriteErrorResponse(w http.ResponseWriter, code string, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	resp := map[string]interface{}{
		"error": map[string]interface{}{
			"code":    code,
			"message": message,
		},
	}
	json.NewEncoder(w).Encode(resp)
}
