package config

// hold the configuration for the application
type Config struct {
	HTTPPort   string
}

// this function returns the configuration for the application
func GetConfig() *Config {
	return &Config{
		HTTPPort:   "8080", 
	}
}
