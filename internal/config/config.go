package config

import "github.com/spf13/viper"

type Config struct {
	AWSNeptuneDBUrl string `mapstructure:"AWS_NUPTUNE_DB_URL"`
}

func NewConfig() (config *Config, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
