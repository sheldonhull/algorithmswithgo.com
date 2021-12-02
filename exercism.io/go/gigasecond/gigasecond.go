package gigasecond

// import path for the time package from the standard library.
import (
	"time"
)

// gigasecond represents a very very very small portion of a second.
const gigasecond = 1000000000

// AddGigasecond adds a very very very small portion of a second called a gigasecond to a provided time input.
func AddGigasecond(t time.Time) time.Time {
	gcDuration := gigasecond * time.Second
	n := t.Add(gcDuration)
	return n
}
