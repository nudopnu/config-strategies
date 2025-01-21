package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
	Runtime  Runtime  `mapstructure:"runtime"`
}

type Server struct {
	Host      string `mapstructure:"host"`
	JwtSecret string `mapstructure:"jwt_secret"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DbName   string `mapstructure:"dbname"`
	Password string `mapstructure:"password"`
}

type Runtime struct {
	RuntimeSetup string `mapstructure:"runtime_setup"`
}

func LoadConfig() (config *Config, err error) {
	log.Printf("TESTT")
	envPrefix := "APP"
	loadSecrets(envPrefix)
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()

	viper.SetConfigFile("config.toml")
	viper.MergeInConfig()
	viper.SetConfigFile("config-override.toml")
	viper.MergeInConfig()

	config = &Config{}
	if err := viper.Unmarshal(config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	return
}

func loadSecrets(envPrefix string) {
	secrets := map[string]string{
		"database.password": "DATABASE_PASSWORD",
		"server.jwt_secret": "SERVER_JWT_SECRET",
	}
	for key, envVar := range secrets {
		// bind secret to environment value
		// this is necessary, because it is not present in config
		// and viper thus not knows about this
		envSecret := fmt.Sprintf("%s_%s", envPrefix, envVar)
		viper.BindEnv(key, envSecret)
		// look for *_FILE environment variables
		envSecretFile := fmt.Sprintf("%s_%s_FILE", envPrefix, envVar)
		secretFile, ok := os.LookupEnv(envSecretFile)
		log.Printf("found %s:%s", envSecretFile, secretFile)
		if !ok {
			continue
		}
		// check if corresponding variable is already set
		_, ok = os.LookupEnv(envSecret)
		if ok {
			continue
		}
		// load secret from file into viper
		value, err := os.ReadFile(secretFile)
		if err != nil {
			log.Fatal(err)
		}
		viper.Set(key, string(value))
	}
}
