package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	handler := InitTestHandler(t)

	req := httptest.NewRequest(http.MethodGet, "/users/new-dummy", nil)
	rr := httptest.NewRecorder()

	handler.RegisterUser(rr, req)

	response := rr.Result()

	expected := map[string]interface{}{
		"status": "ok",
		"user": map[string]interface{}{
			"id":        float64(0),
			"name":      "Mike",
			"is_active": true,
		},
	}

	AssertStatusCode(t, response, http.StatusOK)
	AssertJSONBody(t, response, expected)
}
