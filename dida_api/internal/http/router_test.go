package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRouterCreatesTask(t *testing.T) {
	router := NewRouter()
	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(`{"title":"Inbox task"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", rec.Code)
	}
}
