package moment

import "time"

var LOC = time.Now().Location()

func TruncateByDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, LOC)
}
