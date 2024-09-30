package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	PGDB              DBConfig          `mapstructure:"pg"`
	GoogleOAuthConfig GoogleOAuthConfig `mapstructure:"googleOAuth"`
	Cookie            Cookie            `mapstructure:"cookie"`
	Environment       string            `mapstructure:"environment"`
	AppName           string            `mapstructure:"appName"`
}
type Cookie struct {
	CookieNameAuth string        `mapstructure:"cookieNameAuth"`
	Secret         string        `mapstructure:"secret"`
	Expires        time.Duration `mapstructure:"expires"`
}
type DBConfig struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
}

type GoogleOAuthConfig struct {
	ClientID         string   `mapstructure:"clientID"`
	ClientSecret     string   `mapstructure:"clientSecret"`
	RedirectURL      string   `mapstructure:"redirectURL"`
	Scopes           []string `mapstructure:"scopes"`
	Endpoint         Endpoint `mapstructure:"endpoint"`
	OauthStateString string   `mapstructure:"oauthStateString"`
}

type Endpoint struct {
	AuthURL  string `mapstructure:"authURL"`
	TokenURL string `mapstructure:"tokenURL"`
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
