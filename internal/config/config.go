package config

type Config struct {
	AppPort  string `mapstructure:"APP_PORT"`
	DBUrl    string `mapstructure:"DB_URL"`
	LogLevel string `mapstructure:"LOG_LEVEL"`
}
