package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	PGDB DBConfig `mapstructure:"pg"`
}

type DBConfig struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
}

func Load() *Config {

	config := Config{}
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error occurs while reading the config. ", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("error occurs while unmarshalling the config. ", err)
	}
	return &config
}
