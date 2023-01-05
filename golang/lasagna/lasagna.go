package lasagna

const LayerTime = 2
const OvenTime = 40

func RemainingOvenTime(actualMinutesInOven int) int {
	if actualMinutesInOven < 0 || actualMinutesInOven > OvenTime {
		return 0
	}
	return OvenTime - actualMinutesInOven
}

func PreparationTime(numberOfLayers int) int {
	return LayerTime * numberOfLayers
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}
