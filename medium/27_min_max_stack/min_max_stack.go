package min_max_stack

type MinMaxStack struct {
	stack       []int
	minMaxStack []entry
}

type entry struct {
	min int
	max int
}

func NewMinMaxStack() *MinMaxStack {
	return &MinMaxStack{}
}

func (stack *MinMaxStack) Peek() int {
	return stack.stack[len(stack.stack)-1]
}

func (stack *MinMaxStack) Pop() int {
	stack.minMaxStack = stack.minMaxStack[:len(stack.minMaxStack)-1]
	out := stack.stack[len(stack.stack)-1]
	stack.stack = stack.stack[:len(stack.stack)-1]
	return out
}

func (stack *MinMaxStack) Push(number int) {
	newMinMax := entry{min: number, max: number}
	if len(stack.minMaxStack) > 0 {
		lastMinMax := stack.minMaxStack[len(stack.minMaxStack)-1]
		newMinMax.min = min(lastMinMax.min, number)
		newMinMax.max = max(lastMinMax.max, number)
	}
	stack.minMaxStack = append(stack.minMaxStack, newMinMax)
	stack.stack = append(stack.stack, number)
}

func (stack *MinMaxStack) GetMin() int {
	return stack.minMaxStack[len(stack.minMaxStack)-1].min
}

func (stack *MinMaxStack) GetMax() int {
	return stack.minMaxStack[len(stack.minMaxStack)-1].max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//
//import "math"
//
//type MinMaxStack struct {
//	prevMinValue, prevMaxValue []int
//	minValue, maxValue int
//	bucket []int
//}
//
//func NewMinMaxStack() *MinMaxStack {
//	return &MinMaxStack{
//		minValue: math.MaxInt32,
//		maxValue: math.MinInt32,
//	}
//}
//
//func (stack *MinMaxStack) Push(number int) {
//	stack.bucket = append(stack.bucket, number)
//
//	if number >= stack.maxValue {
//		stack.prevMaxValue = append(stack.prevMaxValue, stack.maxValue)
//		stack.maxValue = number
//	}
//
//	if number <= stack.minValue {
//		stack.prevMinValue = append(stack.prevMinValue, stack.minValue)
//		stack.minValue = number
//	}
//}
//
//func (stack *MinMaxStack) Pop() int {
//	poped := stack.bucket[len(stack.bucket)-1]
//	stack.bucket = stack.bucket[:len(stack.bucket)-1]
//
//	if len(stack.bucket) == 0 {
//		*stack = MinMaxStack{
//			prevMinValue: []int{},
//			prevMaxValue: []int{},
//			minValue: math.MaxInt32,
//			maxValue: math.MinInt32,
//		}
//
//		return poped
//	}
//
//
//	if poped == stack.GetMin() {
//		stack.minValue = stack.prevMinValue[len(stack.prevMinValue)-1]
//		stack.prevMinValue = stack.prevMinValue[:len(stack.prevMinValue)-1]
//	}
//
//	if poped == stack.GetMax() {
//		stack.maxValue = stack.prevMaxValue[len(stack.prevMaxValue)-1]
//		stack.prevMaxValue = stack.prevMaxValue[:len(stack.prevMaxValue)-1]
//	}
//
//	return poped
//}
//
//func (stack *MinMaxStack) GetMin() int {
//	return stack.minValue
//}
//
//func (stack *MinMaxStack) GetMax() int {
//	return stack.maxValue
//}
//
//func (stack *MinMaxStack) Peek() int {
//	return stack.bucket[len(stack.bucket)-1]
//}
