package advent_test

import (
	"testing"

	advent "algo/adventofcode/2021/01-01"
	helpers "algo/adventofcode/testhelpers"

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
