package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	main "git.curve.tools/go/playground/adventofcode/aoc2020/day14"
)

func TestSumBitmask(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name: "1",
			input: []string{
				"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
				"mem[8] = 11",
				"mem[7] = 101",
				"mem[8] = 0",
			},
			want: 165,
		},
		{
			name: "2",
			input: []string{
				"mask = 1101000X0110100X1X1X00001010XX00X0X0",
				"mem[48482] = 6450058",
				"mem[5309] = 2219920",
				"mem[27274] = 43286042",
				"mem[40233] = 504849",
			},
			want: 223778842900,
		},
		{
			name: "3",
			input: []string{
				"mask = 1X1101XX1110X000101100XX001X000100X1",
				"mem[40763] = 53122377",
				"mem[24930] = 244",
				"mem[60497] = 16546",
				"mem[60898] = 654",
			},
			want: 194215942472,
		},
		{
			name: "4",
			input: []string{
				"mask = 1101000X0110100X1X1X00001010XX00X0X0",
				"mem[48482] = 6450058",
				"mem[5309] = 2219920",
				"mem[27274] = 43286042",
				"mem[40233] = 504849",
				"mask = 1X1101XX1110X000101100XX001X000100X1",
				"mem[40763] = 53122377",
				"mem[24930] = 244",
				"mem[60497] = 16546",
				"mem[60898] = 654",
			},
			want: 417994785372,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := main.SumBitmask(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
