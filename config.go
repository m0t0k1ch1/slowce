package main

type Config struct {
	Port   int               `json:"port"`
	Rules  map[string]string `json:"rules"`
	Slack  *SlackConfig      `json:"slack"`
	Consul *ConsulConfig     `json:"consul"`
}

type SlackConfig struct {
	Token string `json:"token"`
}

type ConsulConfig struct {
	BinPath string `json:"bin_path"`
}
