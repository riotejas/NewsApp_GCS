package newsclients

import (
	"context"
	"fmt"

	"io"
	"log/slog"
	"net/http"
	cfg "newsApp/internal/config"
)

// News from Australia and US
// https://newsdata.io/api/1/latest?apikey=<APIKEY>&country=au,us
// English w/ to/from date
// https://newsdata.io/api/1/archive?apikey=<APIKEY>&q=example&language=en&from_date=2023-01-19&to_date=2023-01-25
// Latest articles on EVs filtered by AI tag "technology"
// https://newsdata.io/api/1/latest?apikey=<APIKEY>&q=electric%20vehicles&tag=technology

var NewsDataName = "newsdata"

type NewsDataService interface {
	Speech(context.Context) ([]byte, error)
	Scan(context.Context) (int, []byte, error)
	Sources(context.Context) (int, []byte, error)
}

type newsData struct {
	url string
	params
}
type params struct {
	apiKey           string
	id               string
	q                string
	qInTitle         string
	qInMeta          string
	timeframe        string
	country          string
	category         string
	excludeCategory  string
	language         string
	domain           string
	domainUrl        string
	excludeDomain    string
	excludeField     string
	priorityDomain   string
	timeZone         string
	fullContent      string
	image            bool
	video            bool
	removeDuplicates string
	size             int
	page             int
}

func NewNewsDataService() NewsDataService {
	nc := &newsData{}
	slog.Info("Loading apiKey")
	config := cfg.NewConfig()
	err := config.LoadConfig(NewsDataName)
	if err != nil {
		panic(err)
	}
	nc.apiKey = config.ApiKey
	nc.country = config.Country
	nc.language = config.Language
	return nc
}

func (nc *newsData) Speech(ctx context.Context) ([]byte, error) {
	return nil, nil
}

func (nc *newsData) Scan(ctx context.Context) (int, []byte, error) {
	requestURL := fmt.Sprintf(
		"%s1/latest?apikey=%s&country=%s&language=%s",
		nc.url, nc.apiKey, nc.country, nc.language)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return 0, nil, fmt.Errorf("error creating request: %w", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("error sending news latest request: %w", err)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, nil, fmt.Errorf("error reading response body: %w", err)
	}

	return res.StatusCode, resBody, nil
}

func (nc *newsData) Sources(ctx context.Context) (int, []byte, error) {
	requestURL := fmt.Sprintf(
		"https://newsdata.io/api/1/sources?apikey=%s&country=%s&language=%s",
		nc.apiKey, nc.country, nc.language)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return 0, nil, fmt.Errorf("error creating nc request: %w", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("error sending news sources request: %w", err)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, nil, fmt.Errorf("error reading nc response body: %w", err)
	}
	return res.StatusCode, resBody, nil
}

func (nc *newsData) Health(ctx context.Context) (map[string]string, error) {
	slog.Info("Checking health")
	// todo: what can we check for health wise?
	return map[string]string{
		"message": "App is healthy",
	}, nil
}
