package moment

import "time"

func TruncateByDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func TruncateByHour(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, t.Hour(), 0, 0, 0, t.Location())
}

func TruncateByMinute(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, t.Hour(), t.Minute(), 0, 0, t.Location())
}
