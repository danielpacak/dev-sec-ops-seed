package main

import (
	"context"
	"github.com/danielpacak/dev-sec-ops-seed/pkg/etc"
	"github.com/danielpacak/dev-sec-ops-seed/pkg/http/api"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run() (err error) {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(false)
	log.SetFormatter(&log.JSONFormatter{})

	log.WithFields(log.Fields{
		"version":  version,
		"commit":   commit,
		"built_at": date,
	}).Debug("Starting seed")

	apiHandler := api.NewAPIHandler(api.BuildInfo{
		Version: version,
		Commit:  commit,
		Date:    date,
	})

	apiConfig, err := etc.GetAPIConfig()
	if err != nil {
		return
	}

	server := http.Server{
		Handler:      apiHandler,
		Addr:         apiConfig.Addr,
		ReadTimeout:  apiConfig.ReadTimeout,
		WriteTimeout: apiConfig.WriteTimeout,
	}

	shutdownComplete := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, os.Kill)
		captured := <-sigint
		log.WithField("signal", captured.String()).Debug("Trapped os signal")

		log.Debug("Graceful shutdown started")
		if err := server.Shutdown(context.Background()); err != nil {
			log.WithError(err).Error("Error while shutting down server")
		}
		log.Debug("Graceful shutdown completed")
		close(shutdownComplete)
	}()

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Error: %v", err)
		}
		log.Debug("ListenAndServe returned")
	}()
	<-shutdownComplete
	return
}
