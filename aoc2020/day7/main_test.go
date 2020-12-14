package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	main "git.curve.tools/go/playground/adventofcode/aoc2020/day7"
)

func TestFindNumBags(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name: "first",
			input: []string{
				"light red bags contain 1 bright white bag, 2 muted yellow bags.",
				"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
				"bright white bags contain 1 shiny gold bag.",
				"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
				"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
				"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
				"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
				"faded blue bags contain no other bags.",
				"dotted black bags contain no other bags.",
			},
			want: 32,
		},
		{
			name: "second",
			input: []string{
				"shiny gold bags contain 2 dark red bags.",
				"dark red bags contain 2 dark orange bags.",
				"dark orange bags contain 2 dark yellow bags.",
				"dark yellow bags contain 2 dark green bags.",
				"dark green bags contain 2 dark blue bags.",
				"dark blue bags contain 2 dark violet bags.",
				"dark violet bags contain no other bags.",
			},
			want: 126,
		},
	}

	bag := "shiny gold"

	for _, tt := range tests {
		rules := main.ParseRules(tt.input)
		got := main.FindNumBags(bag, rules)
		assert.Equal(t, tt.want, got)
	}
}

func TestFindNumRuleCombination(t *testing.T) {
	input := []string{
		"light red bags contain 1 bright white bag, 2 muted yellow bags.",
		"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
		"bright white bags contain 1 shiny gold bag.",
		"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
		"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
		"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
		"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
		"faded blue bags contain no other bags.",
		"dotted black bags contain no other bags.",
	}

	bag := "shiny gold"

	rules := main.ParseRules(input)
	got := main.FindNumRuleCombination(bag, rules)
	assert.Equal(t, 4, got)
}

func TestParseRules(t *testing.T) {
	input := []string{
		"light red bags contain 1 bright white bag, 2 muted yellow bags.",
		"dotted black bags contain no other bags.",
		"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
		"bright white bags contain 1 shiny gold bag.",
	}

	want := map[string]*main.Rule{
		"light red": {
			BagColour: "light red",
			Counts: map[string]int{
				"bright white": 1,
				"muted yellow": 2,
			},
			Children: map[string]*main.Rule{
				"bright white": {
					BagColour: "bright white",
					Counts: map[string]int{
						"shiny gold": 1,
					},
					Children: map[string]*main.Rule{
						"shiny gold": {BagColour: "shiny gold", Counts: make(map[string]int), Children: make(map[string]*main.Rule)},
					},
				},
				"muted yellow": {BagColour: "muted yellow", Counts: make(map[string]int), Children: make(map[string]*main.Rule)},
			},
		},
		"dotted black": {
			BagColour: "dotted black", Counts: make(map[string]int), Children: make(map[string]*main.Rule),
		},
		"dark orange": {
			BagColour: "dark orange",
			Counts: map[string]int{
				"bright white": 3,
				"muted yellow": 4,
			},
			Children: map[string]*main.Rule{
				"bright white": {
					BagColour: "bright white",
					Counts: map[string]int{
						"shiny gold": 1,
					},
					Children: map[string]*main.Rule{
						"shiny gold": {BagColour: "shiny gold", Counts: make(map[string]int), Children: make(map[string]*main.Rule)},
					},
				},
				"muted yellow": {BagColour: "muted yellow", Counts: make(map[string]int), Children: make(map[string]*main.Rule)},
			},
		},
		"bright white": {
			BagColour: "bright white",
			Counts: map[string]int{
				"shiny gold": 1,
			},
			Children: map[string]*main.Rule{
				"shiny gold": {BagColour: "shiny gold", Counts: make(map[string]int), Children: make(map[string]*main.Rule)},
			},
		},
		"muted yellow": {BagColour: "muted yellow", Counts: make(map[string]int), Children: make(map[string]*main.Rule)},
		"shiny gold":   {BagColour: "shiny gold", Counts: make(map[string]int), Children: make(map[string]*main.Rule)},
	}

	got := main.ParseRules(input)
	assert.Equal(t, want, got)
}
