package advent

func CalcIncreaseEvents(inputs []int) (total int) {
	for i, item := range inputs {
		// if current value is great than prior we'll count as a increase event
		if i > 0 && item > inputs[(i-1)] {
			total++
		}
	}
	return total
}
