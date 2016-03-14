package timing

import (
	"time"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

// GetTime returns the time duration since the start of the program
func GetTime() time.Duration {
	return time.Since(startTime)
}
