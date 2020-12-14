package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	main "git.curve.tools/go/playground/adventofcalendar20/day10"
)

func TestFindJoltageDifference(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name: "small",
			input: []int{
				16,
				10,
				15,
				5,
				1,
				11,
				7,
				19,
				6,
				12,
				4,
			},
			want: 35,
		},
		{
			name: "large",
			input: []int{
				28,
				33,
				18,
				42,
				31,
				14,
				46,
				20,
				48,
				47,
				24,
				23,
				49,
				45,
				19,
				38,
				39,
				11,
				1,
				32,
				25,
				35,
				8,
				17,
				7,
				9,
				4,
				2,
				34,
				10,
				3,
			},
			want: 220,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := main.FindJoltageDifference(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFindAllCombinations(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name: "small",
			input: []int{
				16,
				10,
				15,
				5,
				1,
				11,
				7,
				19,
				6,
				12,
				4,
			},
			want: 8,
		},
		{
			name: "large",
			input: []int{
				28,
				33,
				18,
				42,
				31,
				14,
				46,
				20,
				48,
				47,
				24,
				23,
				49,
				45,
				19,
				38,
				39,
				11,
				1,
				32,
				25,
				35,
				8,
				17,
				7,
				9,
				4,
				2,
				34,
				10,
				3,
			},
			want: 19208,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := main.FindAllCombinations(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
