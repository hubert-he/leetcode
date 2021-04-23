package stack

func ReverseInPlace(stack *[]int) {
	if len(*stack) == 0 || len(*stack) == 1{
		return
	}
	var addTop func(int)
	addTop = func(bottom int){
		top := (*stack)[0]
		(*stack)=(*stack)[1:]
		if len(*stack) == 0{
			(*stack) = append((*stack), top, bottom)
			return
		}
		addTop(bottom)
		(*stack) = append([]int{top}, (*stack)...) // push
	}
	var dfs func() int
	dfs = func() int{
		top := (*stack)[0]
		(*stack)=(*stack)[1:]
		if len(*stack) == 0{
			(*stack) = append(*stack, top)
			return top
		}
		bottom := dfs()
		addTop(top)
		return bottom
	}
	dfs()
	return
}

type Stack struct{
	array []interface{}
}

func New() Stack {
	return Stack{}
}
func Construct(a []interface{}) Stack{
	return Stack{array: a}
}

func (st *Stack) Push(value interface{}){
	st.array = append([]interface{}{value}, st.array...)
}

func (st *Stack) Pop() (value interface{}) {
	if st.Empty(){
		return nil
	}
	value = st.array[0]
	st.array = st.array[1:]
	return
}

func (st Stack) Top() (value interface{}) {
	if st.Empty(){
		return
	}
	value = st.array[0]
	return
}

func (st Stack) Empty()bool{
	if len(st.array) > 0{
		return false
	}else {
		return true
	}
}

func (st *Stack) Clear(){
	st.array = []interface{}{}
}

func(st *Stack) Reverse() {
	var dfs func()
	var top2bottom func(interface{})
	top2bottom = func(value interface{}){
		top := st.Pop()
		if st.Empty(){
			st.Push(value)
			st.Push(top)
			return
		}
		top2bottom(value)
		st.Push(top)
	}
	dfs = func(){
		top := st.Pop()
		if st.Empty(){
			st.Push(top)
			return
		}
		dfs()
		top2bottom(top)
	}
	dfs()
}
