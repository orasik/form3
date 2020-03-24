package form3

import (
	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
	"time"
)

type Config struct {
	Port            int           `env:"PORT" envDefault:"2400"`
	AccountBaseURL  string        `env:"API_BASEURL,required"`
	AccountEndPoint string        `env:"ACCOUNT_ENDPOINT,required"`
	LogFormat       string        `env:"LOG_FORMAT" envDefault:"json"`
	LogLevel        string        `env:"LOG_LEVEL" envDefault:"debug"`
	Timeout         time.Duration `env:"API_TIMEOUT" envDefault:"3s"`
}

func ParseConfig() (Config, error) {
	var cfg Config
	err := env.Parse(&cfg)

	if err != nil {
		return Config{}, err
	}

	if cfg.LogFormat == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}

	switch cfg.LogLevel {
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	default:
		log.SetLevel(log.DebugLevel)
	}

	return cfg, nil
}
