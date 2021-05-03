package array_of_products

func ArrayOfProducts(array []int) []int {
	products := make([]int, len(array))
	for i := range array {
		products[i] = 1
	}

	leftRunningProduct := 1
	for i := range array {
		products[i] = leftRunningProduct
		leftRunningProduct *= array[i]
	}

	rightRunningProduct := 1
	for i := len(array) - 1; i >= 0; i-- {
		products[i] *= rightRunningProduct
		rightRunningProduct *= array[i]
	}
	return products
}

func ArrayOfProducts1(array []int) []int {
	products := make([]int, len(array))
	leftProducts := make([]int, len(array))
	rightProducts := make([]int, len(array))

	for i := range array {
		products[i] = 1
		leftProducts[i] = 1
		rightProducts[i] = 1
	}

	leftRunningProduct := 1
	for i := range array {
		leftProducts[i] = leftRunningProduct
		leftRunningProduct *= array[i]
	}

	rightRunningProduct := 1
	for i := len(array) - 1; i >= 0; i-- {
		rightProducts[i] = rightRunningProduct
		rightRunningProduct *= array[i]
	}

	for i := range array {
		products[i] = leftProducts[i] * rightProducts[i]
	}
	return products
}


// Better space but does not have better time
func arrayOfProducts1(array []int) []int {
	products := make([]int, len(array))
	for i := range array {
		runningProduct := 1
		for j := range array {
			if i != j {
				runningProduct *= array[j]
			}
		}
		products[i] = runningProduct
	}
	return products
}

// 	input := []int{5, 1, 4, 2}
//	expected := []int{8, 40, 10, 20}

// my solution not very well
func arrayOfProducts(arr []int) (result []int) {
	for idx := range arr {
		mewArr := make([]int, idx)
		copy(mewArr, arr[:idx])
		mewArr = append(mewArr, arr[idx+1:]...)
		var product = 1
		for _, el := range mewArr {
			product *= el
		}
		result = append(result, product)
	}
	return
}

