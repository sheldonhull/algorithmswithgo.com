package advent_test

import (
	"testing"

	advent "algo/adventofcode/2021/01-02"
	helpers "algo/adventofcode/testhelpers"

	iz "github.com/matryer/is"
)

const (
	adventYear   = 2021
	adventDay    = 1
	puzzleNumber = 2
)

func TestGetCalcEntries(t *testing.T) {
	targetSum := 2020
	t.Run("small sample set", func(t *testing.T) {
		is := iz.New(t)

		inputs := []int{
			1721,
			979,
			366,
			299,
			675,
			1456,
		}
		got := advent.GetCalcEntries(inputs, targetSum)
		want := 514579
		is.Equal(got, want) // should match
	})
	t.Run("adventinput", func(t *testing.T) {
		is := iz.New(t)
		inputs := helpers.ReadInputs(t, adventYear, adventDay)
		got := advent.GetCalcEntries(inputs, targetSum)
		want := 0
		is.Equal(got, want) // should match
	})
}
