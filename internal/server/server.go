package server

import (
	"fmt"
	"net/http"
	"newsApp/internal/newsclients"
	"newsApp/internal/speech"
	"time"
)

// To change the news provider, edit the two 'newsclients' statements below

type Server struct {
	nc     newsclients.NewsApiService
	port   int
	speech *speech.Speech
}

func NewServer(port int) *http.Server {
	NewSvr := &Server{
		port: port,
		nc:   newsclients.NewNewsApiService(),
	}

	speech, err := speech.NewSpeechClient()
	if err != nil {

	}
	NewSvr.speech = speech
	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewSvr.port),
		Handler:      NewSvr.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
