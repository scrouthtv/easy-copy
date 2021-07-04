package main

import "time"

var (
	lastSize         uint64 = 0
	lastWatchdogSize uint64 = 0
)

const (
	speedTicker    = 500     // how often the speed should be recalculated
	watchdogTicker = 60*1000 // how fast the watchdog should run
)

var (
	sizePerSecond float32   = 0
	lastTime      time.Time = time.Now()
)

func speedLoop() {
	ticker := time.NewTicker(time.Duration(speedTicker) * time.Millisecond)
	for tick := range ticker.C {
		var seconds float32 = float32(tick.Sub(lastTime).Seconds())
		if seconds != 0 {
			sizePerSecond = float32(doneSize-lastSize) / float32(seconds)
			lastSize = doneSize
		}
	}

	ticker.Stop()
}

func watchdog() {
	ticker := time.NewTicker(time.Duration(watchdogTicker) * time.Millisecond)
	for _ = range ticker.C {
		if doneSize - lastWatchdogSize < 8 {
			errStall()
		}
	}
}
