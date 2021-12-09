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

func CalcIncreaseEventsSliding(inputs []int) (total int) {
	const windowSize = 3

	for idx := range inputs {
		var currentWindowSum int
		var nextWindowSum int
		// if we can't compare due to not enough space left in array
		// to give a n sliding window, then return 0
		if idx == len(inputs) {
			continue // TODO: don't think this is needed?
		}
		// if we can't compare to an equal window since near the end of the window, then exit comparison
		if idx+windowSize >= len(inputs) {
			continue
		}
		// calc current window based on n windowSize
		for i := 0; i < windowSize; i++ {
			currentWindowSum += inputs[idx+i]
		}

		// calc next Window based on N size
		for i := 0; i < windowSize; i++ {
			nextWindowSum += inputs[(idx+1)+i]
		}

		// If the next window > current window, this is an increase event and we should add
		if nextWindowSum > currentWindowSum {
			total++
		}
	}

	return total
}
