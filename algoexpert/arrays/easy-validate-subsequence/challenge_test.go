package challenge_test

import (
	"testing"

	challenge "algo.local/algoexpert/arrays/easy-validate-subsequence"
	iz "github.com/matryer/is"
)

func TestIsValidSubsequence(t *testing.T) {
	is := iz.New(t)
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}

	got := challenge.IsValidSubsequence(array, sequence)
	want := true

	is.Equal(got, want) // Should be true
}

func TestIsValidSubsequenceBetter(t *testing.T) {
	is := iz.New(t)
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}

	got := challenge.IsValidSubsequenceBetter(array, sequence)
	want := true

	is.Equal(got, want) // Should be true
}
