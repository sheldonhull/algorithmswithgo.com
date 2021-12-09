package advent_test

import (
	"testing"

	advent "algo.local/adventofcode/2021/01"
	helpers "algo.local/adventofcode/testhelpers"

	iz "github.com/matryer/is"
)

const (
	adventYear   = 2021
	adventDay    = 1
	puzzleNumber = 1
)

func TestCalcIncreaseEvents(t *testing.T) {
	t.Run("small sample set", func(t *testing.T) {
		is := iz.New(t)
		inputs := []int{
			199, // (N/A - no previous measurement)
			200, // (increased)
			208, // (increased)
			210, // (increased)
			200, // (decreased)
			207, // (increased)
			240, // (increased)
			269, // (increased)
			260, // (decreased)
			263, // (increased)
		}
		got := advent.CalcIncreaseEvents(inputs)
		want := 7
		is.Equal(got, want) // should match want
	})
	t.Run("adventinput", func(t *testing.T) {
		is := iz.New(t)
		inputs := helpers.ReadInputs(t, adventYear, adventDay)
		got := advent.CalcIncreaseEvents(inputs)
		want := 1692
		is.Equal(got, want) // should match
	})
}

func TestCalcIncreaseEventsSliding(t *testing.T) {
	t.Run("small sample set", func(t *testing.T) {
		is := iz.New(t)
		inputs := []int{
			199, // A
			200, // A B
			208, // A B C
			210, //   B C D
			200, // E   C D
			207, // E F   D
			240, // E F G
			269, //   F G H
			260, //     G H
			263, //       H
		}
		got := advent.CalcIncreaseEventsSliding(inputs)
		want := 5
		is.Equal(got, want) // event count should match want
	})
	t.Run("adventinput", func(t *testing.T) {
		is := iz.New(t)
		inputs := helpers.ReadInputs(t, adventYear, adventDay)
		got := advent.CalcIncreaseEventsSliding(inputs)
		want := 1724
		is.Equal(got, want) // should match
	})
}
