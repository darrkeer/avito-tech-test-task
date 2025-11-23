package tests

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"reflect"
	"testing"

	"github.com/darrkeer/avito-tech-test-task/handlers"
	"github.com/darrkeer/avito-tech-test-task/repository"
	"github.com/jmoiron/sqlx"
)

func InitTestDB(t *testing.T) *sql.DB {
	t.Helper()

	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	tx, err := db.Beginx()
	if err != nil {
		t.Fatalf("failed to start transaction: %v", err)
	}

	t.Cleanup(func() {
		tx.Rollback()
		db.Close()
	})

	return db.DB
}

func InitTestHandler(t *testing.T) *handlers.Handler {
	t.Helper()

	repo := repository.New(InitTestDB(t))
	handler := handlers.New(repo)

	return handler
}

func AssertStatusCode(t *testing.T, response *http.Response, expected int) {
	t.Helper()

	if response.StatusCode != expected {
		t.Errorf("expected status code %d, got %d", expected, response.StatusCode)
	}
}

func AssertJSONBody(t *testing.T, response *http.Response, expected map[string]interface{}) {
	t.Helper()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	defer response.Body.Close()

	var got map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &got); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("JSON mismatch.\nExpected: %+v\nGot: %+v", expected, got)
	}
}
