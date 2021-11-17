package cars

// productionRate returns the number of cars that can be created relative to the speed of 1 per hour.
const productionRate = 221

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	switch {
	case speed == 0:
		return 0.0
	case speed >= 1 && speed < 5:
		return 1.0
	case speed >= 5 && speed < 9:
		return 0.9
	case speed >= 9 && speed <= 10:
		return 0.77
	default:
		return 0.0
	}
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	return (float64(speed) * productionRate) * SuccessRate(speed)
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	return int(CalculateProductionRatePerHour(speed) / 60)
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	p := CalculateProductionRatePerHour(speed)
	if p > limit {
		return limit
	}
	return p
}
