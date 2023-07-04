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
	RateLimiter RateLimiter `mapstructure:"ratelimtier"`
}
type HTTPServer struct {
	Address      string        `mapstructure:"address"`
	Timeout      time.Duration `mapstructure:"timeout"`
	Idle_timeout time.Duration `mapstructure:"time"`
}

type Storage struct {
	User                string `mapstructure:"user"`
	Password            string `mapstructure:"password"`
	Host                string `mapstructure:"host"`
	Port                string `mapstructure:"port"`
	DBName              string `mapstructure:"dbname"`
	SSLMode             string `mapstructure:"sslmode"`
	PoolMaxConns        int    `mapstructure:"pool_max_conns"`
	PoolMinConns        int    `mapstructure:"pool_min_conns"`
	PoolMaxConnIdleTime string `mapstructure:"pool_max_conn_idle_time"`
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

type RateLimiter struct {
	MaxReq     int           `mapstructure:"maxreq"`
	Expiration time.Duration `mapstructure:"expiration"`
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
