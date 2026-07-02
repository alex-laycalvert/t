package utils

import (
	"slices"
	"strconv"
	"strings"
)

const DateTimeLayout = "15:04:05 Mon 02/01/2006"

func FormatElapsedTime(elapsed int64, durations ...Duration) string {
	if len(durations) == 0 {
		durations = defaultDurations
	} else {
		// defaultDurations is already sorted
		slices.Sort(durations)
		slices.Reverse(durations)
	}

	builder := strings.Builder{}
	for _, d := range durations {
		durationAmount := elapsed / int64(d)
		elapsed %= int64(d)

		if durationAmount > 0 {
			if builder.Len() != 0 {
				builder.WriteString(" ")
			}

			builder.WriteString(strconv.FormatInt(durationAmount, 10))
			builder.WriteString(d.Label())
		}
	}

	return builder.String()
}
