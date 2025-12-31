package timex_test

import (
	"testing"
	"time"
	_ "time/tzdata"

	"github.com/ndx-technologies/timex"
)

func TestIsTimeZoneEqual(t *testing.T) {
	// For locations with DST, equality depends on the time.
	// Test with a time in Northern Hemisphere summer.
	summerTime := time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC)
	// Test with a time in Northern Hemisphere winter.
	winterTime := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	ny, _ := time.LoadLocation("America/New_York")     // EST: UTC-5, EDT: UTC-4
	chicago, _ := time.LoadLocation("America/Chicago") // CST: UTC-6, CDT: UTC-5
	paris, _ := time.LoadLocation("Europe/Paris")      // CET: UTC+1, CEST: UTC+2
	fixedMinus5 := time.FixedZone("UTC-5", -5*3600)
	fixedMinus4 := time.FixedZone("UTC-4", -4*3600)

	t.Run("equal", func(t *testing.T) {
		tests := []struct {
			name string
			a, b *time.Location
			t    time.Time
		}{
			{"nil locations", nil, nil, summerTime},
			{"same location object", ny, ny, summerTime},
			{"different objects same location", ny, ny, summerTime},
			{"different locations same offset summer", ny, fixedMinus4, summerTime},
			{"different locations same offset winter", ny, fixedMinus5, winterTime},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				if !timex.IsTimeZoneEqual(tc.a, tc.b, tc.t) {
					t.Error(tc.a, tc.b, tc.t)
				}
			})
		}
	})

	t.Run("not equal", func(t *testing.T) {
		tests := []struct {
			name string
			a, b *time.Location
			t    time.Time
		}{
			{"one nil location", ny, nil, summerTime},
			{"different locations different offset summer", ny, fixedMinus5, summerTime},
			{"different locations different offset winter", ny, fixedMinus4, winterTime},
			{"different locations different offset", ny, chicago, summerTime},
			{"different locations different offset", ny, paris, summerTime},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				if timex.IsTimeZoneEqual(tc.a, tc.b, tc.t) {
					t.Error(tc.a, tc.b, tc.t)
				}
			})
		}
	})
}
