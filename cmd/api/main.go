package main

import (
	"fmt"
	"log/slog"
	cfg "newsApp/internal/config"
	"newsApp/internal/server"
)

func main() {
	config := cfg.NewConfig()
	err := config.LoadConfig("")
	if err != nil {
		panic(err)
	}

	slog.Info("Starting server on", "port", config.Port)
	svr := server.NewServer(config.Port)

	err = svr.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

}
