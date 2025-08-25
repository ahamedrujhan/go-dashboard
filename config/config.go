package config

import "github.com/spf13/viper"

type Config struct {
	Port     string
	Host     string
	Username string
	Password string
	Db_name  string
}

func LoadConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	return Config{
		Port:     viper.GetString("port"),
		Host:     viper.GetString("host"),
		Username: viper.GetString("username"),
		Password: viper.GetString("password"),
		Db_name:  viper.GetString("db_name"),
	}
}
