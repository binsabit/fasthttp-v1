package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port    string
	DB      string
	LogFile string
}

var (
	defaults = map[string]interface{}{
		"port":    "8080",
		"db":      "postgres://yerdaulet:pa55word@localhost:5432/prosclad",
		"logfile": "../log/log.txt",
	}
	configPath = []string{"./config"}
)

func Configure() Config {
	for k, v := range defaults {
		viper.SetDefault(k, v)
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	for _, path := range configPath {
		viper.AddConfigPath(path)
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("could not read config file:%v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("could not unmarshal config:%v", err)
	}
	return config

}
