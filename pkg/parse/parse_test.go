package parse_test

import (
	"testing"
	"time"

	"github.com/adamdecaf/tz/pkg/parse"

	"github.com/stretchr/testify/require"
)

func TestParseTime(t *testing.T) {
	cases := []struct {
		input string

		expectedFormat string
		expectedTime   time.Time
	}{
		{
			input:          "Oct 09, 2025 14:50 UTC",
			expectedFormat: "Jan 2, 2006 15:04 MST",
			expectedTime:   time.Date(2025, time.October, 9, 14, 50, 0, 0, time.UTC),
		},
	}
	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			got, format, err := parse.Time(tc.input)
			require.NoError(t, err)

			require.Equal(t, tc.expectedTime, got)
			require.Equal(t, tc.expectedFormat, format)
		})
	}
}
