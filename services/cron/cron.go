package cron

import (
	"go-log-generator/config"
	"go-log-generator/services/generator"

	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

func Routine() {
	log.Info().Msg("cron initialized")

	c := cron.New()

	if config.Cron.EnableCron { // toggle cron activation
		c.AddFunc(config.Cron.IntervalCron, TheJob)
	}

	c.Start()
}
func TheJob() {
	//the job for cron routine
	generator.GenerateLog(100)
}
