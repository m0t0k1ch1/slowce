package main

import (
	"encoding/json"
	"errors"
	"flag"
	"log"
	"os"
)

const (
	DefaultConfigPath = "config.json"
)

var (
	ErrInvalidOutgointWebHooksToken = errors.New("invalid Outgoing WebHooks token")
	ErrUnknownMessage               = errors.New("unknown message")
)

func main() {
	var configPath = flag.String("conf", DefaultConfigPath, "path to your config file")
	flag.Parse()

	configFile, err := os.Open(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	if err := json.NewDecoder(configFile).Decode(&config); err != nil {
		log.Fatal(err)
	}

	NewApp(config).Run()
}
