package config

import "github.com/spf13/viper"

func init() {
	viper.SetConfigName("visitor")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("/etc/visitor")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
