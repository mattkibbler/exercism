// Package gigasecond handles all gigasecond stuff
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond adds one gigasecond to a provided time
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(1000000000 * time.Second)
	return t
}
