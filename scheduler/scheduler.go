package scheduler

import "time"

type Scheduler struct{}

func (s *Scheduler) Start(interval time.Duration, f func()) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			f()
		}
	}
}
