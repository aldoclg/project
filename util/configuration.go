package util

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (*Configuration, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	viper.ReadInConfig()

	c := new(Configuration)

	err := viper.Unmarshal(c)

	return c, err
}
