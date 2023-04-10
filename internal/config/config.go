package config

type Config struct {
	Server Server
	Db     Db
}
type Db struct {
	Password string
	UserName string
	Url      string
	Port     string
	Driver   string
	Name     string
	LogLevel int
}
type Server struct {
	Port string
}

func (c *Config) getPort() string {
	return c.Server.Port
}
func (c *Config) getDbUserName() string {
	return c.Db.UserName
}
func (c *Config) getDbPassword() string {
	return c.Db.Password
}
func (c *Config) getDbUrl() string {
	return c.Db.Url
}

func (c *Config) getDbPort() string {
	return c.Db.Port
}
func (c *Config) getDriver() string {
	return c.Db.Driver
}
func (c *Config) getName() string {
	return c.Db.Name
}
func (c *Config) geLogLevel() int {
	return c.Db.LogLevel
}
