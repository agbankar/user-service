package config

/*
Port: 8080
DbUrl:
DbUserName: root
DbPawssword: test
DbName: work
DbDialect: mysql
DbPort: 3306

*/
type Config struct {
	Port       string `yaml:"Port"`
	DbUserName string `yaml:"DbUserName"`
	DbPassword string `yaml:"DbPassword"`
	DbUrl      string `yaml:"DbUrl"`
	DbPort     string `yaml:"DbPort"`
}

func (c *Config) getPort() string {
	return c.Port
}
func (c *Config) getDbUserName() string {
	return c.DbUserName
}
func (c *Config) getDbPassword() string {
	return c.DbPassword
}
func (c *Config) getDbUrl() string {
	return c.DbUrl
}

func (c *Config) getDbPort() string {
	return c.DbPort
}
