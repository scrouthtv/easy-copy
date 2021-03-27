package main

import "time"

var lastSize uint64 = 0

// in what intervals to measure time (millis)
const ms int = 500

var (
	sizePerSecond float32   = 0
	lastTime      time.Time = time.Now()
)

func speedLoop() {
	ticker := time.NewTicker(time.Duration(ms) * time.Millisecond)
	for tick := range ticker.C {
		var seconds float32 = float32(tick.Sub(lastTime).Seconds())
		if seconds != 0 {
			sizePerSecond = float32(done_size-lastSize) / float32(seconds)
			lastSize = done_size
		}
	}
	ticker.Stop()
}
