package main

import "time"

var lastSize uint64 = 0

// ms specifies how often the speed should be recalculated.
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
			sizePerSecond = float32(doneSize-lastSize) / float32(seconds)
			lastSize = doneSize
		}
	}
	ticker.Stop()
}
