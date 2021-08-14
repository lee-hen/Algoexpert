package detect_arbitrage
// important question

import "math"

// DetectArbitrage
// O(n^3) time | O(n^2) space - where n is the number of currencies
func DetectArbitrage(exchangeRates [][]float64) bool {
	// To use exchange rates as edge weights, we must be able to add them.
	// Since log(a*b) = log(a) + log(b), we can convert all rates to
	// -log10(rate) to use them as edge weights.
	logExchangeRates := convertToLogMatrix(exchangeRates)

	// A negative weight cycle indicates an arbitrage.
	return foundNegativeWeightCycle(logExchangeRates, 0)
}

// Runs the Bellman–Ford Algorithm to detect any negative-weight cycles.
// Bellman–Ford runs in O(|V|*|E|) time, where |V| and |E| are the number of vertices and edges respectively.
func foundNegativeWeightCycle(graph [][]float64, start int) bool {
	distancesFromStart := make([]float64, len(graph))
	for i := range distancesFromStart {
		distancesFromStart[i] = math.MaxFloat64
	}
	distancesFromStart[start] = 0.0

	for unused := 0; unused < len(graph)-1; unused++ {
		// If no update occurs, that means there's no negative cycle.
		if !relaxEdgesAndUpdateDistances(graph, distancesFromStart) {
			return false
		}
	}

	// if A*B >= 1 (log(a)+log(b)>=0) this will not be needed.
	return relaxEdgesAndUpdateDistances(graph, distancesFromStart)
}

// Returns `true` if any distance was updated.
func relaxEdgesAndUpdateDistances(graph [][]float64, distances []float64) bool {
	var updated = false
	for sourceIdx := range graph {
		edges := graph[sourceIdx]
		for destinationIdx := range edges {
			edgeWeight := edges[destinationIdx]
			newDistanceToDestination := distances[sourceIdx] + edgeWeight
			if newDistanceToDestination < distances[destinationIdx] {
				updated = true
				distances[destinationIdx] = newDistanceToDestination
			}
		}
	}

	return updated
}

func convertToLogMatrix(matrix [][]float64) [][]float64 {
	newMatrix := make([][]float64, len(matrix))
	for row, rates := range matrix {
		for _, rate := range rates {
			newMatrix[row] = append(newMatrix[row], -math.Log10(rate))
		}
	}

	return newMatrix
}



//[1, 0.5, 0.25, 2, 4],
//[2, 1, 0.5, 4, 8],
//[4, 2, 1, 8, 16],
//[0.5, 0.25, 0.0125, 1, 2],
//[0.25, 0.0125, 0.00625, 0.5, 1]

// [1 2 3 4]
// [0 2 3 4]
// [0 1 3 4]
// [0 1 2 4]
// [0 1 2 3]


//       0:USD 1:CAD  2:GBP
//0:USD [  1.0, 1.27, 0.718]
//1:CAD [ 0.74,  1.0,  0.56]
//2:GBP [ 1.39, 1.77,   1.0]

// [1 2]
// [0 2]
// [0 1]


//{1.0, 0.8631, 0.5903},
//{1.1586, 1.0, 0.6849},
//{1.6939, 1.46, 1.0},

// i think this is not correct but it works
func detectArbitrage(exchangeRates [][]float64) bool {
	for idx := range exchangeRates {
		if traverseExchangeRates(idx, idx, exchangeRates[idx][idx], exchangeRates) {
			return true
		}
	}

	return false
}

func traverseExchangeRates(idx, currencyUnit int, currentExchange float64, exchangeRates [][]float64) bool {
	for i := 0; i < len(exchangeRates[idx]); i++ {
		if i == idx && i == currencyUnit {
			continue
		}

		if currentExchange * exchangeRates[idx][i] * exchangeRates[i][currencyUnit] > exchangeRates[currencyUnit][currencyUnit] {
			return true
		}

		if i > idx && traverseExchangeRates(i, currencyUnit, currentExchange * exchangeRates[idx][i], exchangeRates) {
			return true
		}
	}

	return false
}
