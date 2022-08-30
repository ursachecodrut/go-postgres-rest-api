package config

import "github.com/spf13/viper"

type Config struct {
	Port  string `mapstructure:"PORT"`
	DBUrl string `mapstructure:"DB_URL"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/common/config.env")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return

}
