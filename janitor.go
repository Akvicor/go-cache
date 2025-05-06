package cache

import (
	"time"
)

// janitor interface use for any cache
type janitorInterface interface {
	DeleteExpired()
	SetJanitor(*janitor)
	StopJanitor()
}

// stop janitor
func stopJanitor(c janitorInterface) {
	c.StopJanitor()
}

// run janitor
func runJanitor(c janitorInterface, ci time.Duration) {
	j := &janitor{
		interval: ci,
		stop:     make(chan bool),
	}
	c.SetJanitor(j)
	go j.run(c)
}

// expired item cleaner
type janitor struct {
	interval time.Duration
	stop     chan bool
}

// clean up expired data
func (j *janitor) run(c janitorInterface) {
	ticker := time.NewTicker(j.interval)
	for {
		select {
		case <-ticker.C:
			c.DeleteExpired()
		case <-j.stop:
			ticker.Stop()
			return
		}
	}
}
