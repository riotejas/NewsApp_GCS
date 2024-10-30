package server

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

func handleError(w http.ResponseWriter, r *http.Request, err error, status int) {
	http.Error(w, err.Error(), status)
}

type NewsResponse struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
		Source struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"source"`
		Author      string    `json:"author"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Url         string    `json:"url"`
		UrlToImage  string    `json:"urlToImage"`
		PublishedAt time.Time `json:"publishedAt"`
		Content     string    `json:"content"`
	} `json:"articles"`
}

func (s *Server) RegisterRoutes() http.Handler {
	// Create a slog logger
	//logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	// Use the slog-chi middleware
	//r.Use(slogchi.New(logger))
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Route("/v1", func(r chi.Router) {

		r.Get("/speech", s.SpeechHandler)

		r.Get("/scan", s.ScanHandler)

		r.Get("/sources", s.SourcesHandler)

		r.Get("/doc", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("TBD"))
		})
	})
	r.Get("/health", s.HealthHandler)

	// simulate an error, allows user to validate scripts
	r.Get("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(400), 400)
	})

	return r
}

func (s *Server) SpeechHandler(w http.ResponseWriter, r *http.Request) {

	body, err := s.nc.Scan(r.Context())
	if err != nil {
		slog.Error("SpeechHandler", "scan error", err)
		handleError(w, r, err, http.StatusInternalServerError)
		return
	}

	var response NewsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		slog.Error("SpeechHandler", "unmarshal error", err)
	}

	var synText string
	for i, article := range response.Articles {
		synText = fmt.Sprintf("%s. Article %d. %s", synText, i+1, article.Description)
	}

	var audioContent []byte
	audioContent, err = s.speech.SpeechClient(synText)
	if err != nil {
		slog.Error("SpeechHandler", "speech error", err)
		handleError(w, r, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", strconv.Itoa(len(audioContent)))
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(audioContent)
}

func (s *Server) ScanHandler(w http.ResponseWriter, r *http.Request) {

	body, err := s.nc.Scan(r.Context())
	if err != nil {
		slog.Error("ScanHandler", "error", err)
		handleError(w, r, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(body)
}

func (s *Server) SourcesHandler(w http.ResponseWriter, r *http.Request) {

	body, err := s.nc.Sources(r.Context())
	if err != nil {
		slog.Error("SourcesHandler", "error", err)
		handleError(w, r, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(body)
}

func (s *Server) HealthHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Checking health")

	res := map[string]string{
		"message": "App is healthy",
	}
	jsonResp, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonResp)
}
