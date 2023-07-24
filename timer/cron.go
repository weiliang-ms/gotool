package timer

import (
	"github.com/robfig/cron/v3"
)

func StartCronTimer(spec string, f func()) {
	c := cron.New()
	c.AddFunc(spec, f)
	c.Start()
	select {}
}
