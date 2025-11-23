package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

var body1 = `{
	"team_name": "payments",
	"members": [
		{
			"user_id": "u1",
			"username": "Alice",
			"is_active": true
		},
		{
			"user_id": "u2",
			"username": "Bob",
			"is_active": true
		}
	]
}`

func TestTeamAddSimple(t *testing.T) {
	handler := initTestHandler(t)

	req := httptest.NewRequest(http.MethodPost, "/team/add", bytes.NewBufferString(body1))
	rr := httptest.NewRecorder()

	handler.TeamAdd(rr, req)

	response := rr.Result()

	resBody, _ := makeJSONFrom(body1)
	expected := map[string]interface{}{"team": resBody}

	assertStatusCode(t, response, http.StatusOK)
	assertJSONBody(t, response, expected)
}
