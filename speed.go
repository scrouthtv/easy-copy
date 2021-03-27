package main

import "time"

var last_size uint64 = 0

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
			sizePerSecond = float32(done_size-last_size) / float32(seconds)
			last_size = done_size
		}
	}
	ticker.Stop()
}
