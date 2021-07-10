package progress

import (
	"errors"
	"time"
)

var Start time.Time

var (
	lastSize         uint64 = 0
	lastWatchdogSize uint64 = 0
)

const (
	speedTicker    = 500       // how often the speed should be recalculated (ms)
	watchdogTicker = 60 * 1000 // how fast the watchdog should run
)

var (
	SizePerSecond float32   = 0
	lastTime      time.Time = time.Now()
)

func WatchSpeed() {
	ticker := time.NewTicker(time.Duration(speedTicker) * time.Millisecond)
	for tick := range ticker.C {
		var seconds float32 = float32(tick.Sub(lastTime).Seconds())
		if seconds != 0 {
			SizePerSecond = float32(DoneSize-lastSize) / seconds
			lastSize = DoneSize
		}
	}

	ticker.Stop()
}

var ErrStall = errors.New("less than 8 bytes transferred in 1 minute")

func Watchdog(err chan error) {
	ticker := time.NewTicker(time.Duration(watchdogTicker) * time.Millisecond)
	for range ticker.C {
		if DoneSize-lastWatchdogSize < 8 {
			err <- ErrStall
		}
	}

	ticker.Stop()
}
