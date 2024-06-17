package kanyerest

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	httpClient := &http.Client{Timeout: 5 * time.Second}

	cl := NewYeRestClient(httpClient)

	got, err := cl.GetQuote(ctx)
	if err != nil {
		t.Errorf("GetQuote() error = %s", err)

		return
	}

	if got == nil {
		t.Errorf("GetQuote() result should not be nil")
	}
}
