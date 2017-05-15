package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
)

type Event struct {
	name string
}

func (app *App) convertToEvent(r *http.Request) (*Event, error) {
	text := r.Form.Get("text")
	triggerWord := r.Form.Get("trigger_word")

	re := regexp.MustCompile(fmt.Sprintf(`^(%s)`, triggerWord))

	message := re.ReplaceAllString(text, "")
	message = strings.TrimSpace(message)

	eventName, ok := app.config.Rules[message]
	if !ok {
		return nil, ErrUnknownMessage
	}

	return &Event{
		name: eventName,
	}, nil
}

func (app *App) fireEvent(event *Event) error {
	return exec.Command(app.config.Consul.BinPath, "event", "-name", event.name).Run()
}
