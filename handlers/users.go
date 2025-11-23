package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/darrkeer/avito-tech-test-task/models"
)

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	newUser := models.NewUser(0, "Mike", true)

	if err := h.repository.InsertUser(newUser); err != nil {
		http.Error(w, "failed to insert user", http.StatusInternalServerError)
		return
	}

	resp := map[string]interface{}{
		"status": "ok",
		"user":   newUser,
	}
	json.NewEncoder(w).Encode(resp)
}
