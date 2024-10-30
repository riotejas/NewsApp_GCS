package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port     int
	ApiKey   string
	Url      string
	Country  string
	Language string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) LoadConfig(vendor string) error {
	v := viper.New()

	// read environment variables.  env vars have precedence over config file
	v.AutomaticEnv()

	// read config file
	v.AddConfigPath(".")
	v.SetConfigType("yaml")
	v.SetConfigName("config")
	err := v.ReadInConfig()
	if err != nil {
		return err
	}
	c.Port = v.GetInt("apiServer.port")

	if vendor == "" {
		return nil
	}
	c.ApiKey = v.GetString(fmt.Sprintf("%s_api_key", vendor))
	c.Url = v.GetString(fmt.Sprintf("%s.url", vendor))
	c.Country = v.GetString(fmt.Sprintf("%s.country", vendor))
	c.Language = v.GetString(fmt.Sprintf("%s.language", vendor))

	if !v.IsSet("apiServer.port") || c.Port == 0 {
		return fmt.Errorf("port not set")
	}
	if !v.IsSet(fmt.Sprintf("%s_api_key", vendor)) || c.ApiKey == "" {
		return fmt.Errorf("apiKey not set")
	}

	return nil
}
