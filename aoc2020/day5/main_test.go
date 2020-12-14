package main_test

import (
	"testing"

	main "git.curve.tools/go/playground/adventofcalendar20/day5"
	"github.com/stretchr/testify/assert"
)

func TestGetSeatID(t *testing.T) {
	tests := []struct {
		name         string
		boardingPass []byte
		want         int
	}{
		{
			name:         "FBFBBFFRLR",
			boardingPass: []byte("FBFBBFFRLR"),
			want:         357,
		},
		{
			name:         "BFFFBBFRRR",
			boardingPass: []byte("BFFFBBFRRR"),
			want:         567,
		},
		{
			name:         "FFFBBBFRRR",
			boardingPass: []byte("FFFBBBFRRR"),
			want:         119,
		},
		{
			name:         "BBFFBBFRLL",
			boardingPass: []byte("BBFFBBFRLL"),
			want:         820,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := main.GetSeatID(tt.boardingPass)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestHighestSeatID(t *testing.T) {
	input := []string{
		"FBFBBFFRLR",
		"BFFFBBFRRR",
		"FFFBBBFRRR",
		"BBFFBBFRLL",
	}
	_, got := main.HighestSeatID(input)
	assert.Equal(t, 820, got)
}

func TestFindGap(t *testing.T) {
	input := []int{
		1, 1, 0, 1, 1,
	}

	got := main.FindGap(input)

	assert.Equal(t, 2, got)
}
