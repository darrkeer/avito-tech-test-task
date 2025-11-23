package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/darrkeer/avito-tech-test-task/models"
)

func (h *Handler) TeamAdd(w http.ResponseWriter, r *http.Request) {
	var r_team models.ReadTeam
	if err := json.NewDecoder(r.Body).Decode(&r_team); err != nil {
		WriteErrorResponse(w, "BAD_INPUT", "could not parse input json", http.StatusBadRequest)
		return
	}

	err := h.repository.AddTeam(&r_team)

	switch {
	case errors.Is(err, errors.New("TEAM_EXISTS")):
		WriteErrorResponse(w, "TEAM_EXISTS", "team_name already exists", http.StatusBadRequest)

	case err != nil:
		WriteErrorResponse(w, "UNKNOWN", "unknown error", http.StatusBadRequest)

	default:
	}

	resp := map[string]interface{}{
		"team": r_team,
	}
	WriteResponse(w, &resp)
}
