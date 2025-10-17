package utils

import "time"

// return the current time in GMT + 0 timezone
func GMTTime() time.Time {
	return time.Now().In(time.FixedZone("GMT", 0))
}
