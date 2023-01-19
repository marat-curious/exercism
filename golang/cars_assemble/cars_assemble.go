package cars

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / float64(100))
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(float64(productionRate) / float64(60) * successRate / float64(100))
}

func CalculateCost(carsCount int) uint {
	var r int = carsCount % 10
	return uint(((carsCount - r) / 10) * 95000 + r * 10000)
}
