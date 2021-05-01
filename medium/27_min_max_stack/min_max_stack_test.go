package min_max_stack

import "testing"

func testMinMaxPeek(t *testing.T, min, max, peek int, stack *MinMaxStack) {
	if min != stack.GetMin() {
		t.Fail()
	}
	if max != stack.GetMax() {
		t.Fail()
	}
	if peek != stack.Peek() {
		t.Fail()
	}
}

func assertEqual(t *testing.T, a, b int) {
	if a != b {
		t.Fail()
	}
}

func TestCase1(t *testing.T) {
	stack := NewMinMaxStack()
	stack.Push(5)
	testMinMaxPeek(t, 5, 5, 5, stack)
	stack.Push(7)
	testMinMaxPeek(t, 5, 7, 7, stack)
	stack.Push(2)
	testMinMaxPeek(t, 2, 7, 2, stack)
	assertEqual(t, stack.Pop(), 2)
	assertEqual(t, stack.Pop(), 7)
	testMinMaxPeek(t, 5, 5, 5, stack)
}
