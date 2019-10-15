package etc

import (
	"github.com/caarlos0/env/v6"
	"time"
)

type APIConfig struct {
	Addr         string        `env:"SEED_HTTP_ADDR" envDefault:":8080"`
	ReadTimeout  time.Duration `env:"SEED_HTTP_READ_TIMEOUT" envDefault:"15s"`
	WriteTimeout time.Duration `env:"SEED_HTTP_WRITE_TIMEOUT" envDefault:"15s"`
}

func GetAPIConfig() (cfg APIConfig, err error) {
	err = env.Parse(&cfg)
	return
}
