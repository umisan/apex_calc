package main

import (
	"reflect"
	"testing"
)

func TestCalcKillPoint(t *testing.T) {
	type Input struct {
		kill     int
		leverage int
	}
	cases := []struct {
		input  Input
		output int
	}{
		{input: Input{kill: 0, leverage: 0}, output: 0},
		{input: Input{kill: 0, leverage: 10}, output: 0},
		{input: Input{kill: 1, leverage: 10}, output: 10},
		{input: Input{kill: 5, leverage: 20}, output: 100},
		{input: Input{kill: 6, leverage: 20}, output: 120},
		{input: Input{kill: 7, leverage: 20}, output: 120},
	}
	for _, c := range cases {
		actual := calcKillPoint(c.input.kill, c.input.leverage)
		if actual != c.output {
			t.Errorf("have: %d, want: %d", actual, c.output)
		}
	}
}

func TestMapTableIndexToRank(t *testing.T) {
	cases := []struct {
		input  int
		output int
	}{
		{input: 0, output: 14},
		{input: 5, output: 9},
		{input: 10, output: 4},
	}
	for _, c := range cases {
		actual := mapTableIndexToRank(c.input)
		if actual != c.output {
			t.Errorf("have %d, want %d", actual, c.output)
		}
	}
}

func TestReverseCalculation(t *testing.T) {
	type Input struct {
		needPoint int
		class     int
	}
	cases := []struct {
		input  Input
		output CandList
	}{
		{
			input: Input{
				needPoint: 100,
				class:     1,
			},
			output: CandList{
				Cand{Rank: 5, Kill: 6},
				Cand{Rank: 4, Kill: 5},
				Cand{Rank: 3, Kill: 4},
				Cand{Rank: 2, Kill: 3},
				Cand{Rank: 1, Kill: 1},
			},
		},
		{
			input: Input{
				needPoint: 200,
				class:     3,
			},
			output: CandList{
				Cand{Rank: 1, Kill: 6},
			},
		},
	}
	for _, c := range cases {
		actual := reverseCalculation(c.input.needPoint, c.input.class)
		if !reflect.DeepEqual(actual, c.output) {
			t.Errorf("have %v, want %v", actual, c.output)
		}
	}
}
