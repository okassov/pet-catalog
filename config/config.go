package config

import "github.com/spf13/viper"

type Config struct {
	PG     PGConfig     `mapstructure:",squash"`
	Server ServerConfig `mapstructure:",squash"`
}

type ServerConfig struct {
	Port string `mapstructure:"SERVER_PORT"`
}

type PGConfig struct {
	PGUrl          string `mapstructure:"PG_URL"`
	PGPort         string `mapstructure:"PG_PORT"`
	PGUser         string `mapstructure:"PG_USER"`
	PGPassword     string `mapstructure:"PG_PASSWORD"`
	PGDatabase     string `mapstructure:"PG_DATABASE"`
	PGMaxPool      int    `mapstructure:"PG_MAX_POOL"`
	PGConnAttempts int    `mapstructure:"PG_CONN_ATTEMPTS"`
	PGConnTimeout  int    `mapstructure:"PG_CONN_TIMEOUT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}



