package number_of_binary_tree_topologies
// important question

// NumberOfBinaryTreeTopologies
// O(n^2) time | O(n) space
func NumberOfBinaryTreeTopologies(n int) int {
	cache := []int{1}
	for m := 1; m < n+1; m++ {
		numberOfTrees := 0
		for leftTreeSize := 0; leftTreeSize < m; leftTreeSize++ {
			rightTreeSize := m - 1 - leftTreeSize
			numberOfLeftTrees := cache[leftTreeSize]
			numberOfRightTrees := cache[rightTreeSize]
			numberOfTrees += numberOfLeftTrees * numberOfRightTrees
		}
		cache = append(cache, numberOfTrees)
	}
	return cache[n]
}

func numberOfBinaryTreeTopologies(n int) int {
	cache := make(map[int]int)
	return getNumberOfBinaryTreeTopologies(n, cache)
}

func getNumberOfBinaryTreeTopologies(n int, cache map[int]int) int {
	if n == 0 {
		return 1
	}
	if num, ok := cache[n]; ok {
		return num
	}

	var result int
	for x := 0; x <= n-1; x++ {
		//3 0 2 map[]
		//2 0 1 map[]
		//1 0 0 map[]
		//2 1 0 map[1:1 2:1]
		//3 1 1 map[1:1 2:2 3:2]
		//3 2 0 map[1:1 2:2 3:3]
		result += getNumberOfBinaryTreeTopologies(x, cache) * getNumberOfBinaryTreeTopologies(n-1-x, cache)
		cache[n] = result
	}

	return result
}
