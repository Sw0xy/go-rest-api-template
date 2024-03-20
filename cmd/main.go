package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Sw0xy/go-rest-api-template/api/route"
	"github.com/Sw0xy/go-rest-api-template/bootstrap"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: false,
	})

	app := bootstrap.Initialize()

	defer bootstrap.CloseDb(app.DB.Client())
	r := mux.NewRouter()
	route.Setup(r, app.DB, time.Second*15)

	srv := &http.Server{
		Addr:         "localhost:3010",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error(err)
		}
	}()

	log.Info("server started")

	// Graceful Shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	srv.Shutdown(ctx)
	log.Info("shutting down")
	os.Exit(0)
}
