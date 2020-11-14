package anagramservice

// Config store configuration of server
type Config struct {
	BindAddr string `toml:"bind_addr"`
}

// NewConfig create and return Config
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
