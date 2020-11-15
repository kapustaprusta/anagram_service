package anagramservice

// Config stores server configuration
type Config struct {
	BindAddr string `toml:"bind_addr"`
}

// NewConfig returns pointer to the Config
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
