package config

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Env        string `mapstructure:"env"`
	Storage    string `mapstructure:"storage"`
	LogFile    string `mapstructure:"logfile"`
	HTTPServer struct {
		Address      string        `mapstructure:"address"`
		Timeout      time.Duration `mapstructure:"timeout"`
		Idle_timeout time.Duration `mapstructure:"time"`
	} `mapstructure:"http"`
}

func MustLoad() *Config {
	configPath := flag.String("config", "./config/config.yaml", "path to configure the project")
	flag.Parse()
	if *configPath == "" {
		log.Fatal("could not get config path")
	}

	//check if config file exists

	if _, err := os.Stat(*configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exists: %v", err)
	}
	viper.SetConfigType("yaml")
	viper.SetConfigFile(*configPath)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("viper could not read config file:%v", err)
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("viper could not unmarshal to config struct:%v", err)

	}
	return &cfg
}
