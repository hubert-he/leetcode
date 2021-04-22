package stack

func ReverseInPlace(stack *[]int) {
	if len(*stack) == 0 || len(*stack) == 1{
		return
	}
	var addTop func(int)
	addTop = func(topValue int){
		if len(*stack) == 0{
			*stack = append(*stack, topValue)
			return
		}
		top := (*stack)[0]
		(*stack) = (*stack)[1:]

	}
	var dfs func() int
	dfs = func() int {
		top := (*stack)[0]
		if len(*stack) == 1{
			(*stack)=(*stack)[1:]
			return top
		}
		bottom := dfs()
		(*stack) = append([]int{bottom}, (*stack)...)
		return bottom
	}
	dfs()
	return
}
