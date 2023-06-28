package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog/log"
)

type common struct {
	Port string `env:"PORT" envDefault:"3300"`
}

type cron struct {
	EnableCron   bool   `env:"ENABLE_CRON" envDefault:"false"`
	IntervalCron string `env:"INTERVAL_CRON" envDefault:"* * * * *"`
}

var (
	Common common
	Cron   cron
)

func init() {
	env.Parse(&Common)
	env.Parse(&Cron)

	log.Debug().Interface("common", Common).Send()
	log.Debug().Interface("cron", Cron).Send()

}
