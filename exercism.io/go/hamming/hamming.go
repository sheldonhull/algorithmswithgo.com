package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	if a == b {
		return 0, nil
	}
	if len(a) != len(b) {
		return 0, errors.New("Sequence lengths are different")
	}

	diffCount := 0

	for i, x := range a {
		if byte(x) != byte(b[i]) {
			diffCount++
		}
	}

	return diffCount, nil
}
