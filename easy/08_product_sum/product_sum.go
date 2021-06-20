package product_sum

type SpecialArray []interface{}

// ProductSum
// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.
func ProductSum(array []interface{}) int {
	return helper(array, 1)
}

func helper(array []interface{}, depth int) int {
	for i, el := range array {
		switch v := el.(type) {
		case int:
			return v*depth + helper(array[i+1:], depth)
		case SpecialArray:
			return helper(array[i+1:], depth) + depth*helper(v, depth+1)
		}
	}
	return 0
}
