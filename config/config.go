package config

// Config ...
type Config struct {
	DB *DBConfig
}

// DBConfig ...
type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

// GetConfig ...
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "root",
			Password: "root",
			Name:     "famouspersons",
			Charset:  "utf8",
		},
	}
}
