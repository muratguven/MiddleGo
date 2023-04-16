package utils

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of application
type Config struct {
	AppMode          string `mapstructure:"APP_MODE"`
	ConnectionString string `mapstructure:"DB_SOURCE"`
}

// LoadConfig reads configuration from file
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
