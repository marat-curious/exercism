package cars

import "testing"

func TestCalculateWorkingCarsPerHour(t *testing.T) {
	var inputProductionRate = [5]int{ 0, 1, 3, 5, 7 }
	var inputSuccessRate = [5]float64{ 10, 0, 50, 100, 30 }
	var expected = [5]float64{ 0.0, 0.0, 1.5, 5, 2.1 }
	for i := 0; i < 5; i++ {
		r := CalculateWorkingCarsPerHour(inputProductionRate[i], inputSuccessRate[i])
		if r != expected[i] {
			t.Errorf("CalculateWorkingCarsPerHour(%d, %f) = %f, wants %f", inputProductionRate[i], inputSuccessRate[i], r, expected[i])
		}
	}
}

func TestCalculateWorkingCarsPreMinute(t *testing.T) {
	var inputProductionRate = [6]int{ 0, 100, 2, 459, 2034, 6824 }
	var inputSuccessRate = [6]float64{ 100.0, 40.0, 80.0, 0.0, 20.0, 20.5 }
	var expected = [6]int{ 0, 0, 0, 0, 6, 23 }
	for i := 0; i < 6; i++ {
		r := CalculateWorkingCarsPerMinute(inputProductionRate[i], inputSuccessRate[i])
		if r != expected[i] {
			t.Errorf("CalculateWorkingCarsPerMinute(%d, %f) = %d, wants %d", inputProductionRate[i], inputSuccessRate[i], r, expected[i])
		}
	}
}

func TestCalculateCost(t *testing.T) {
	var inputCarsCount = [5]int{ 0, 37, 21, 123, 9 }
	var expected = [5]uint{ 0, 355000, 200000, 1170000, 90000 }
	for i := 0; i < 5; i++ {
		r := CalculateCost(inputCarsCount[i])
		if r != expected[i] {
			t.Errorf("CalculateCost(%d) = %d, wants %d", inputCarsCount[i], r, expected[i])
		}
	}
}
