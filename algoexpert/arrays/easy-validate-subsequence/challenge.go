package challenge

import "fmt"

func IsValidSubsequence(array []int, sequence []int) bool {
	fmt.Printf("sequence (len %d):%+v\tarray (len %d): %+v\n", len(sequence), sequence, len(array), array)

	if len(sequence) > len(array) {
		fmt.Println("SKIP: len(sequence) > len(array)")
		return false
	}

	matchedSequenceIndex := 0
	for i := 0; i < len(array); i++ {
		if matchedSequenceIndex >= len(sequence) {
			fmt.Printf("DONE: reached end of sequence evaluation\n")
			break
		}
		if array[i] == sequence[matchedSequenceIndex] {
			fmt.Printf("\tarray[%d] %d == sequence[%d] %d\tmatchedSequenceIndex: %d\n", i, array[i], matchedSequenceIndex, sequence[matchedSequenceIndex], matchedSequenceIndex)
			matchedSequenceIndex++
			continue
		}
	}
	if matchedSequenceIndex < len(sequence) {
		fmt.Printf("FAIL: matchedSequenceIndex %d < len(sequence)-1 %d\n", matchedSequenceIndex, len(sequence)-1)
		return false
	}
	if matchedSequenceIndex == len(sequence) {
		fmt.Printf("PASS: matchedSequenceIndex %d == len(sequence)-1 %d\n", matchedSequenceIndex, len(sequence)-1)
		return true
	}
	fmt.Printf("UNKNOWN CONDITION: matchedSequenceIndex %d // len(sequence)-1 %d\n", matchedSequenceIndex, len(sequence)-1)
	return false
}
