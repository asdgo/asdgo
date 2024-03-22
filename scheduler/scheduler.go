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

func (s *Scheduler) EveryMinute(f func()) {
	s.Start(time.Minute, f)
}

func (s *Scheduler) EveryFiveMinutes(f func()) {
	s.Start(5*time.Minute, f)
}

func (s *Scheduler) EveryFifteenMinutes(f func()) {
	s.Start(15*time.Minute, f)
}

func (s *Scheduler) EveryThirtyMinutes(f func()) {
	s.Start(30*time.Minute, f)
}

func (s *Scheduler) EveryHour(f func()) {
	s.Start(time.Hour, f)
}
