package config

type Config struct {
	Endpoint         string   `env:"ENDPOINT"`
	ListenPort       int      `env:"LISTEN_PORT"`
	InterfaceAddress string   `env:"INTERFACE_ADDRESS"`
	AllowedIPs       []string `env:"ALLOWED_IPS" envSeparator:","`
	WorkDir          string   `env:"WORKDIR"`
	WGConfigPath     string   `env:"WG_CONFIG_PATH"`
}

func NewConfig() *Config {
	return &Config{
		Endpoint:         "10.0.0.1/24",
		ListenPort:       51820,
		InterfaceAddress: "10.0.0.1/24",
		AllowedIPs:       []string{"0.0.0.0/0"},
		WorkDir:          "/etc/wireguard",
		WGConfigPath:     "",
	}
}
