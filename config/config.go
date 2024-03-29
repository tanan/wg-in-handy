package config

import (
	"log/slog"
	"sync"

	"github.com/caarlos0/env/v10"
)

type Config struct {
	WGEndpoint         string   `env:"WG_ENDPOINT"`
	WGInterfaceAddress string   `env:"WG_INTERFACE_ADDRESS"`
	WGNetwork          string   `env:"WG_NETWORK"`
	WGListenPort       int      `env:"WG_LISTEN_PORT"`
	WGAllowedIPs       []string `env:"WG_ALLOWED_IPS" envSeparator:","`
	WGWorkDir          string   `env:"WG_WORKDIR"`
	WGConfigPath       string   `env:"WG_CONFIG_PATH"`
}

func NewConfig() Config {
	return Config{
		WGEndpoint:         "10.0.0.1:51820",
		WGListenPort:       51820,
		WGInterfaceAddress: "10.0.0.1",
		WGNetwork:          "10.0.0.0/22",
		WGAllowedIPs:       []string{"0.0.0.0/0"},
		WGWorkDir:          "/etc/wireguard",
		WGConfigPath:       "",
	}
}

var cfg Config

var newConfigOnceFunc = sync.OnceFunc(func() {
	cfg = NewConfig()
})

var loadConfigOnceFunc = sync.OnceFunc(func() {
	if err := env.Parse(&cfg); err != nil {
		slog.Error("failed to parse environment variables", slog.String("error", err.Error()))
	}
})

func Get() Config {
	newConfigOnceFunc()
	loadConfigOnceFunc()
	return cfg
}
