package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPrice(t *testing.T) {
	// Test for a valid cryptocurrency (Bitcoin)
	req := httptest.NewRequest("GET", "/price/bitcoin", nil)
	w := httptest.NewRecorder()
	GetPrice(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

}
