// Package gigasecond provides a library that adds gigaseconds to a time.Time object.
package gigasecond

import "time"

// AddGigasecond accepts a time.Time arguement and adds a gigasecond, if successful,
// the resulting time.Time is returned
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
