package config

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Env         string      `mapstructure:"env"`
	LogFile     string      `mapstructure:"logfile"`
	Storage     Storage     `mapstructure:"storage"`
	HTTPServer  HTTPServer  `mapstructure:"http"`
	RateLimiter RateLimiter `mapstructure:"ratelimiter"`
}

type HTTPServer struct {
	Address      string        `mapstructure:"address"`
	Timeout      time.Duration `mapstructure:"timeout"`
	Idle_timeout time.Duration `mapstructure:"time"`
}

type Storage struct {
	DBDriver     string        `mapstructure:"db_driver"`
	DSN          string        `mapstructure:"dsn"`
	MaxOpenConns int           `mapstructure:"max_open_conns"`
	MaxIdleConns int           `mapstructure:"max_idle_conns"`
	MaxIdleTime  time.Duration `mapstructure:"max_idle_time"`
	MaxConnLife  time.Duration `mapstructure:"max_life_time"`
}

type RateLimiter struct {
	MaxReq     int           `mapstructure:"maxreq"`
	Expiration time.Duration `mapstructure:"expiration"`
}

type SecurityHeaders struct {
	HSTSMaxAge                int
	HSTSExcludeSubdomains     bool
	HSTSPreload               bool
	ContentSecurityPolicy     string
	XSSProtection             string
	ContentTypeNosniff        string
	XFrameOptions             string
	ReferrerPolicy            string
	CrossOriginEmbedderPolicy string
	CrossOriginOpenerPolicy   string
	CrossOriginResourcePolicy string
	OriginAgentCluster        string
	XDNSPrefetchControl       string
	XDownloadOptions          string
	XPermittedCrossDomain     string
}

func MustLoad() *Config {
	configPath := flag.String("config", "./config.yaml", "path to configure the project")
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
