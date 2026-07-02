package utils

type Duration int64

const (
	Seconds Duration = 1
	Minutes Duration = 60 * Seconds
	Hours   Duration = 60 * Minutes
	Days    Duration = 24 * Hours
	Weeks   Duration = 7 * Days
)

var defaultDurations = []Duration{Weeks, Days, Hours, Minutes, Seconds}

func (d Duration) Label() string {
	switch d {
	case Seconds:
		return "s"
	case Minutes:
		return "m"
	case Hours:
		return "h"
	case Days:
		return "d"
	case Weeks:
		return "w"
	default:
		return ""
	}
}
