package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
	Runtime  Runtime  `mapstructure:"runtime"`
}

type Server struct {
	Host string `mapstructure:"host"`
}

type Database struct {
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	DbName string `mapstructure:"dbname"`
}

type Runtime struct {
	RuntimeSetup string `mapstructure:"runtime_setup"`
}

func LoadConfig() (config *Config, err error) {
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()
	viper.SetConfigFile("config.toml")
	viper.MergeInConfig()
	viper.SetConfigFile("config-override/config.toml")
	viper.MergeInConfig()
	config = &Config{}
	err = viper.Unmarshal(config)
	return
}
