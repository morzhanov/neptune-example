package config

import "github.com/spf13/viper"

type Config struct {
	AWSAccessKeyID     string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AWSSecretAccessKey string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	AWSRegion          string `mapstructure:"AWS_REGION"`
	AWSNeptuneDBName   string `mapstructure:"AWS_NUPTUNE_DB_NAME"`
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
