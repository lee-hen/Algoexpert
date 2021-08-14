package valid_starting_city
// important question

func ValidStartingCity(distances []int, fuel []int, mpg int) int {
	indexOfStartingCityCandidate := 0
	milesRemainingAtStartingCityCandidate := 0

	milesRemaining := 0
	for cityIdx := 1; cityIdx < len(distances); cityIdx++ {

		distanceFromPreviousCity := distances[cityIdx-1]
		fuelFromPreviousCity := fuel[cityIdx-1]

		milesRemaining += fuelFromPreviousCity*mpg - distanceFromPreviousCity
		if milesRemaining < milesRemainingAtStartingCityCandidate {
			milesRemainingAtStartingCityCandidate = milesRemaining
			indexOfStartingCityCandidate = cityIdx
		}
	}

	return indexOfStartingCityCandidate
}

// mpg 10
// 10 20  10  0   30
// 5, 25, 15, 10, 15

// 0-1 5
// 1-2 25
// 2-3 15
// 3-4 10
// 4-0 15

func validStartingCity(distances []int, fuel []int, mpg int) int {
	var startCity, remain int

	for i := 0; i < len(fuel); i++ {
		if fuel[i] * mpg < distances[i] {
			continue
		}

		remain = 0
		startCity = i
		nextCity := startCity
		for {
			if remain = remain + fuel[nextCity] * mpg - distances[nextCity]; remain >= 0 {
				nextCity = (nextCity+1) % len(fuel)

				if nextCity == startCity {
					break
				}
			} else {
				break
			}
		}

		if nextCity == startCity {
			break
		}
	}

	return startCity
}
