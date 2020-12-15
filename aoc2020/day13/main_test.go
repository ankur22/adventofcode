package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	main "git.curve.tools/go/playground/adventofcode/aoc2020/day13"
)

func TestCalcBusIDAndWait(t *testing.T) {
	input := []string{
		"939",
		"7,13,x,x,59,x,31,19",
	}

	want := 295

	got := main.CalcBusIDAndWait(input)
	assert.Equal(t, want, got)
}

func TestCalcEarliestConsecutiveTime(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name: "7,13,15",
			input: []string{
				"939",
				"7,13,15",
			},
			want: 1078,
		},
		{
			name: "7,13,x,x,59,x,31,19",
			input: []string{
				"939",
				"7,13,x,x,59,x,31,19",
			},
			want: 1068781,
		},
		{
			name: "17,x,13,19",
			input: []string{
				"939",
				"17,x,13,19",
			},
			want: 3417,
		},
		{
			name: "67,7,59,61",
			input: []string{
				"939",
				"67,7,59,61",
			},
			want: 754018,
		},
		{
			name: "67,x,7,59,61",
			input: []string{
				"939",
				"67,x,7,59,61",
			},
			want: 779210,
		},
		{
			name: "67,7,x,59,61",
			input: []string{
				"939",
				"67,7,x,59,61",
			},
			want: 1261476,
		},
		{
			name: "1789,37,47,1889",
			input: []string{
				"939",
				"1789,37,47,1889",
			},
			want: 1202161486,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := main.CalcEarliestConsecutiveTime(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
