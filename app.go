package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	config Config
	logger *log.Logger
}

func NewApp(config Config) *App {
	return &App{
		config: config,
		logger: log.New(
			os.Stdout,
			"app: ",
			log.Ldate|log.Ltime,
		),
	}
}

func (app *App) log(v ...interface{}) {
	app.logger.Println(v)
}

func (app *App) logError(err error) {
	app.log("[ERROR]", err)
}

func (app *App) error(w http.ResponseWriter, err error) {
	app.log(err)
	w.WriteHeader(http.StatusInternalServerError)
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	event, err := app.convertToEvent(r)
	if err != nil {
		if err == ErrUnknownMessage {
			app.log("received unknown message")
			return
		}
		app.error(w, err)
		return
	}

	app.log("fire", event.name)
	if err := app.fireEvent(event); err != nil {
		app.error(w, err)
	}
}

func (app *App) Run() {
	addr := fmt.Sprintf(":%d", app.config.Port)
	srv := &http.Server{
		Handler: app,
		Addr:    addr,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	app.log(fmt.Sprintf("listening on %s", addr))

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, syscall.SIGTERM)
	for {
		s := <-signalCh
		if s == syscall.SIGTERM {
			app.log("received SIGTERM")
			srv.Shutdown(context.Background())
		}
	}
}
