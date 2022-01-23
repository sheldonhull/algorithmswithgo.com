package challenge

func IsValidSubsequenceBetter(array []int, sequence []int) bool {
	// loop through main array
	// look for matching sequence
	var sidx int
	for i := range array {
		if array[i] == sequence[sidx] {
			sidx++
		}
		// if we matched all the items then count as success
		if sidx == len(sequence) {
			return true
		}
	}

	// not all items were matched
	return false
}
