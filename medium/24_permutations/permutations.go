package permutations
// important question

func GetPermutations(array []int) [][]int {
	permutations := make([][]int, 0)
	permutationsHelper(0, array, &permutations)
	return permutations
}

func permutationsHelper(i int, array []int, permutations *[][]int) {
	if i == len(array)-1 {
		newPerm := make([]int, len(array))
		copy(newPerm, array)
		*permutations = append(*permutations, newPerm)
		return
	}

	for j := i; j < len(array); j++ {
		// [1, 2, 3] i is first index
		swap(array, i, j)
		// [1, 3, 2] when i is came back to the second index this time j is increased by prev permutationsHelper(i+1 (j is third index)
		// [2, 1, 3] after swap back to [1,2,3] j is increased by prev permutationsHelper(i+1 (j is second index) and i will be the first index
		permutationsHelper(i+1, array, permutations) // just like stack
		// [1, 2, 3] second index swap back
		// [1, 2, 3] first index swap back
		swap(array, i, j) // -> [2, 1, 3]
	}
}

func swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

func getPermutations2(array []int) [][]int {
	permutations := make([][]int, 0)
	permutationsHelper2(array, []int{}, &permutations)
	return permutations
}

func permutationsHelper2(array []int, currentPermutation []int, permutations *[][]int) {
	if len(array) == 0 && len(currentPermutation) != 0 {
		*permutations = append(*permutations, currentPermutation)
		return
	}

	for i := range array {
		// Make a new array that not exist i
		newArray := make([]int, i)
		copy(newArray, array[:i])
		newArray = append(newArray, array[i+1:]...)

		// make a new permutation with that i
		newPermutation := make([]int, len(currentPermutation))
		copy(newPermutation, currentPermutation)
		newPermutation = append(newPermutation, array[i])

		permutationsHelper2(newArray, newPermutation, permutations)
	}
}

// GetPermutations
// nextPermutation is difficult
func getPermutations(array []int) (perms [][]int) {
	if len(array) == 0 {
		return [][]int{}
	}

	perms = append(perms, array)
	newPerm := make([]int, len(array))
	copy(newPerm, array)

	return append(perms, getPermutations(nextPermutation(newPerm))...)
}

// ex.
// [1,2,4,3]
// k = 1 v-> 2
// l = 2 v-> 3
// [1 3 4 2]
// Reverse k+1 up to final position
// [1 3 2 4]
// According the wikipedia
func nextPermutation(arr []int) []int{

	//Find the largest index k such that a[k] < a[k + 1]. If no such index exists, the permutation is the last permutation.
	var k = -1
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			k = i
		}
	}

	if k == -1 {
		return nil
	}

	//Find the largest index l greater than k such that a[k] < a[l].
	var l = -1
	for i := 0; i < len(arr); i++ {
		if arr[k] < arr[i] {
			l = i
		}
	}

	//Swap the value of a[k] with that of a[l].
	if l != -1 {
		arr[k], arr[l] = arr[l], arr[k]
	}

	//Reverse the sequence from a[k + 1] up to and including the final element a[n].
	ReverseSlice(arr[k+1:])

	return arr
}

func ReverseSlice(slice []int) {
	start := 0
	end := len(slice) -1

	for start < end {
		slice[start], slice[end] = slice[end], slice[start]
		start += 1
		end -= 1
	}
}
