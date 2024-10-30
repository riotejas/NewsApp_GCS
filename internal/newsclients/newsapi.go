package newsclients

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	cfg "newsApp/internal/config"
)

var NewsApiName = "newsapi"

type NewsApiService interface {
	Scan(context.Context) ([]byte, error)
	Sources(context.Context) ([]byte, error)
}

type newsApi struct {
	url string
	params
}

func NewNewsApiService() NewsApiService {
	nc := &newsApi{}
	slog.Info("Loading apiKey", "client", NewsApiName)
	config := cfg.NewConfig()
	err := config.LoadConfig(NewsApiName)
	if err != nil {
		panic(err)
	}
	nc.apiKey = config.ApiKey
	nc.url = config.Url
	nc.country = config.Country
	nc.language = config.Language
	return nc
}

func (nc *newsApi) Scan(ctx context.Context) ([]byte, error) {
	requestURL := fmt.Sprintf(
		"%stop-headlines?sources=techcrunch",
		nc.url)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("X-Api-Key", nc.apiKey)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending news headlines request: %w", err)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return resBody, nil
}

func (nc *newsApi) Sources(ctx context.Context) ([]byte, error) {
	requestURL := fmt.Sprintf(
		"%stop-headlines/sources?country=%s&language=%s",
		nc.url, nc.country, nc.language)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating nc request: %w", err)
	}
	req.Header.Set("X-Api-Key", nc.apiKey)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending news sources request: %w", err)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading nc response body: %w", err)
	}
	return resBody, nil
}
