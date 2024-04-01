package config

import (
	"log"
	"sync"
	"time"

	"github.com/spf13/viper"
)

var once sync.Once
var configuration *Configuration

type Configuration struct {
	Database   Database
	Thirdparty Thirdparty
}

type Database struct {
	Driver               string
	Name                 string
	User                 string
	Password             string
	Host                 string
	Port                 int
	AdditionalParameters string
	MaxOpenConns         int
	MaxIdleConns         int
	ConnMaxLifetime      time.Duration
	Debug                bool
}

type Thirdparty struct {
	BaseUrl string
	Key     string
}

func GetConfig() *Configuration {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("../../../.")
		viper.AddConfigPath("../../../../.")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}

		if err := viper.Unmarshal(&configuration); err != nil {
			log.Fatalf("Unable to decode into struct, %v", err)
		}
	})

	return configuration
}
