package laptop_rentals

import "sort"

// LaptopRentals
// O(nlog(n)) time | O(n) space - where n is the number of times
func LaptopRentals(times [][]int) int {
	if len(times) == 0 {
		return 0
	}

	var usedLaptops int
	startTimes, endTimes := make([]int,0), make([]int,0)

	for _, time := range times {
		startTimes = append(startTimes, time[0])
		endTimes = append(endTimes, time[1])
	}

	sort.Ints(startTimes)
	sort.Ints(endTimes)

	var startIterator, endIterator int
	for startIterator < len(times) {
		if startTimes[startIterator] >= endTimes[endIterator] {
			usedLaptops -= 1
			endIterator += 1
		}

		usedLaptops += 1
		startIterator += 1
	}

	return usedLaptops
}

// LaptopRentals1 heap implementation
// O(nlog(n)) time | O(n) space - where n is the number of times
func LaptopRentals1(times [][]int) int {
	if len(times) == 0 {
		return 0
	}
	sort.Slice(times, func(i, j int) bool {
		return times[i][0] < times[j][0]
	})

	heap := NewHeap(MinEndHeapFunc)
	heap.Insert(times[0])

	for idx := 1; idx < len(times); idx++ {
		currentInterval := times[idx]
		if heap.Peek()[1] <= currentInterval[0] {
			heap.Remove()
		}
		heap.Insert(currentInterval)
	}
	return heap.Length()
}

var MinStartHeapFunc = func(a, b []int) bool { return a[0] < b[0] }
var MinEndHeapFunc = func(a, b []int) bool { return a[1] < b[1] }

// laptopRentals
// my solution heap implementation
func laptopRentals(times [][]int) int {
	minStartHeap := NewHeap(MinStartHeapFunc)

	for _, interval := range times {
		minStartHeap.Insert(interval)
	}

	minEndHeap := NewHeap(MinEndHeapFunc)
	for !minStartHeap.IsEmpty() {
		interval := minStartHeap.Remove()

		if !minEndHeap.IsEmpty() && interval[0] >= minEndHeap.Peek()[1] {
			minEndHeap.Update(interval)
		} else {
			minEndHeap.Insert(interval)
		}
	}

	return minEndHeap.Length()
}
