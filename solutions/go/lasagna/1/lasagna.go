package lasagna

const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
    var remainingTime int
    remainingTime = OvenTime - actualMinutesInOven
    return remainingTime
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
    var prepTime int
    prepTime = numberOfLayers * 2
    return prepTime
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
    prepTime := PreparationTime(numberOfLayers)
    elapsedTime := actualMinutesInOven + prepTime
    return elapsedTime
}
