package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestQuery(t *testing.T) {
	req, err := http.NewRequest("GET", "/?kind=suburb&q=well", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	query(w, req)

	resp := w.Result()
	if actual, expected := resp.StatusCode, http.StatusOK; actual != expected {
		t.Errorf("Status code is wrong. Expected %d, got %d.", expected, actual)
	}

	if actual, expected := resp.Header.Get("Content-Type"), "application/json"; actual != expected {
		t.Errorf("Content type is wrong. Expected %s, got %s.", expected, actual)
	}
}
