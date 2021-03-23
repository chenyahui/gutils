package moment

import "time"

func TruncateByDay(t time.Time) time.Time {
	return t.Truncate(24 * time.Hour).Add(-8 * time.Hour)
}
