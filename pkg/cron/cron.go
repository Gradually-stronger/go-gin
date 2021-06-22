package cron

import (
	"github.com/robfig/cron/v3"
)

var CronJobs *cron.Cron

func init() {
	CronJobs = cron.New()
	CronJobs.Start()
}
