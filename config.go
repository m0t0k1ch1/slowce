package main

type Config struct {
	Port   int               `json:"port"`
	Rules  map[string]string `json:"rules"`
	Consul string            `json:"consul"`
}
