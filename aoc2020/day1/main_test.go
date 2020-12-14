package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	main "git.curve.tools/go/playground/adventofcode/aoc2020/day1"
)

func TestMultiplyFind2020(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{values: []int{1721, 979, 366, 1456, 299, 675}},
			//want: 514579,
			want: 241861950,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := main.MultiplyFind2020(tt.args.values)
			assert.Equal(t, tt.want, got)
		})
	}
}
