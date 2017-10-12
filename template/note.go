package template

import (
	"fmt"
	"time"
)

// NextMonday find date of next Monday
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

// BulletJournal print bullet journal template
func BulletJournal(selectTime time.Time) {

	dateFormat := "2006-01-02 Mon"

	nextMon := NextMonday(selectTime)
	nextSun := nextMon.Add(6 * 24 * time.Hour)
	nextSat := nextMon.Add(5 * 24 * time.Hour)
	nextFri := nextMon.Add(4 * 24 * time.Hour)
	nextThu := nextMon.Add(3 * 24 * time.Hour)
	nextWed := nextMon.Add(2 * 24 * time.Hour)
	nextTue := nextMon.Add(1 * 24 * time.Hour)

	_, weekNumber := nextMon.ISOWeek()

	format := `Week %d
%s
%s
%s
%s
%s
%s
%s
`
	fmt.Printf(
		format,
		weekNumber,
		nextSun.Format(dateFormat),
		nextSat.Format(dateFormat),
		nextFri.Format(dateFormat),
		nextThu.Format(dateFormat),
		nextWed.Format(dateFormat),
		nextTue.Format(dateFormat),
		nextMon.Format(dateFormat),
	)
}

// NextWeekBulletJournal print next week bullet journal template
func NextWeekBulletJournal() {
	BulletJournal(time.Now())
}
