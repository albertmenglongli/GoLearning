package stack

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack *Stack) Push(x interface{}) {
	*stack = append(*stack, x)
}
