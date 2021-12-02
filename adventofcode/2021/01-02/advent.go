package advent

import (
	"log"
)

func GetCalcEntries(inputs []int, targetSum int) (total int) {
	for i, item := range inputs {
		for inneri, testNum := range inputs {
			if i == inneri {
				continue
			}
			if testNum+item == targetSum {
				total = testNum * item
				log.Printf("testNum %v + item %v == targetSum %v so returning %v\n", testNum, item, targetSum, total)

				return total
			}
		}
	}
	return total
}
