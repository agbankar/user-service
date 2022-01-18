package config

type Config struct {
	Port string
}

func (c *Config) getPort() string {
	return c.Port
}
