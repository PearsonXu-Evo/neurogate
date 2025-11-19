package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	App    AppConfig
}

type ServerConfig struct {
	Port int
	Mode string
}

type AppConfig struct {
	Name    string
	Version string
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	viper.SetEnvPrefix("neurogate")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		log.Fatalf("Unable to decode into struct: %s", err)
	}

	return &c
}
