package main

import (
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	handler(w, r)
	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Errorf("Expected 200")
	}
}
