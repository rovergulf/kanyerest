package kanyerest

import (
	"net/http"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	httpClient := &http.Client{Timeout: 5 * time.Second}

	cl := NewYeRestClient(httpClient)

	got, err := cl.GetQuote()
	if err != nil {
		t.Errorf("GetQuote() error = %s", err)

		return
	}

	if got == nil {
		t.Errorf("GetQuote() result should not be nil")
	}
}
