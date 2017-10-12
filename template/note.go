package format

import "time"

func NextMonday(t time.Time) time.Time {
	diff := 0
	switch t.Weekday() {
	case time.Sunday:
		diff = 1
	case time.Saturday:
		diff = 2
	case time.Friday:
		diff = 3
	case time.Thursday:
		diff = 4
	case time.Wednesday:
		diff = 5
	case time.Tuesday:
		diff = 6
	default:
		diff = 7
	}

	nt := t.Truncate(time.Hour).Add(time.Duration(diff) * 24 * time.Hour)
	return nt
}
