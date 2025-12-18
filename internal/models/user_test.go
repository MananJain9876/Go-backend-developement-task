package models

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	dob := time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name string
		now  time.Time
		want int
	}{
		{
			name: "before birthday in year",
			now:  time.Date(2025, 5, 9, 0, 0, 0, 0, time.UTC),
			want: 34,
		},
		{
			name: "on birthday in year",
			now:  time.Date(2025, 5, 10, 0, 0, 0, 0, time.UTC),
			want: 35,
		},
		{
			name: "after birthday in year",
			now:  time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC),
			want: 35,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateAge(dob, tt.now)
			if got != tt.want {
				t.Fatalf("CalculateAge() = %d, want %d", got, tt.want)
			}
		})
	}
}


