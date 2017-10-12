package template

import (
	"testing"
	"time"
)

func TestNextMonday(t *testing.T) {

	testCases := []struct {
		input    string
		excepted string
	}{
		{"2017-10-10", "2017-10-16"},
		{"2017-10-09", "2017-10-16"},
		{"2017-10-08", "2017-10-09"},
		{"2017-10-07", "2017-10-09"},
		{"2017-10-06", "2017-10-09"},
	}

	for _, c := range testCases {
		tt, _ := time.Parse("2006-01-02", c.input)
		nt := NextMonday(tt)
		result := nt.Format("2006-01-02")

		if result != c.excepted {
			t.Error("should be:", c.excepted, "when", c.input, "but get:", result)
		}
	}
}

func TestBulletJournal(t *testing.T) {
	NextWeekBulletJournal()
}
