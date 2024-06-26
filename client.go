package kanyerest

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

const apiUrl = `https://api.kanye.rest`

type Quote struct {
	Quote string `json:"quote"`
}

type Client interface {
	GetQuote(ctx context.Context) (*Quote, error)
}

type client struct {
	httpClient *http.Client
}

func NewYeRestClient(httpClient *http.Client) Client {
	return &client{httpClient}
}

func (c *client) GetQuote(ctx context.Context) (*Quote, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiUrl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	quote := new(Quote)
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(quote); err != nil {
		return nil, err
	}

	return quote, nil
}
