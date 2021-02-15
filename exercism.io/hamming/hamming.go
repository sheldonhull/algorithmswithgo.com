package hamming

import (
	"errors"
	"strings"
)

func Distance(a, b string) (int, error) {

	if a == b {
		return 0, nil
	}
	if len(a) != len(b) {
		return 0, errors.New("Sequence lengths are different")
	}

	diffCount := 0
	aString := strings.Split(a, "")
	bString := strings.Split(b, "")

	for i, x := range aString {
		if x != bString[i] {
			diffCount++
		}
	}

	return diffCount, nil
}
