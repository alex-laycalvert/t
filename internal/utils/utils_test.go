package utils

import (
	"testing"
)

func TestFormatElapsedTime(t *testing.T) {
	elapsed := int64(1*Weeks + 2*Days + 3*Hours + 4*Minutes + 5*Seconds)
	want := "1w 2d 3h 4m 5s"
	got := FormatElapsedTime(elapsed)

	if got != want {
		t.Errorf(`FormatElapsedTime(%d) = %q, want match for %#q`, elapsed, got, want)
	}
}

func TestFormatElapsedTimeSomeEmptyDurationValues(t *testing.T) {
	elapsed := int64(1*Weeks + 2*Hours + 3*Seconds)
	want := "1w 2h 3s"
	got := FormatElapsedTime(elapsed)

	if got != want {
		t.Errorf(`FormatElapsedTime(%d) = %q, want match for %#q`, elapsed, got, want)
	}
}

func TestFormatElapsedTimeCustomDurations(t *testing.T) {
	elapsed := int64(1*Weeks + 2*Hours + 3*Seconds)
	want := "7d 120m 3s"
	got := FormatElapsedTime(elapsed, Days, Minutes, Seconds)

	if got != want {
		t.Errorf(`FormatElapsedTime(%d, Days, Minutes, Seconds) = %q, want match for %#q`, elapsed, got, want)
	}
}

func TestFormatElapsedTimeCustomDurationsUnordered(t *testing.T) {
	elapsed := int64(1*Weeks + 2*Hours + 3*Seconds)
	want := "7d 120m 3s"
	got := FormatElapsedTime(elapsed, Days, Minutes, Seconds)

	if got != want {
		t.Errorf(`FormatElapsedTime(%d, Minutes, Seconds, Days) = %q, want match for %#q`, elapsed, got, want)
	}
}

func TestFormatElapsedTimeOnlyLastDuration(t *testing.T) {
	elapsed := int64(3 * Seconds)
	want := "3s"
	got := FormatElapsedTime(elapsed)

	if got != want {
		t.Errorf(`FormatElapsedTime(%d) = %q, want match for %#q`, elapsed, got, want)
	}
}
