package stack

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

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

// 155. Min Stack
type MinStack struct {
	elems 		[]interface{}
	minElem		[]interface{}
}

func Constructor() MinStack{
	// Go 特性： 若是slice 则无需设置初始化；若是map，则必须MinStack{elems:map[int]int{}} 显式初始化
	return MinStack{}
}

func (st *MinStack)Push(val interface{}) {
	// 易错点2
	min := st.GetMin()
	if val.(int) < min.(int){
		min = val
	}
	st.elems = append([]interface{}{val}, st.elems...)
	st.minElem = append([]interface{}{min}, st.minElem...)
}

func (st *MinStack)Pop() interface{}{
	st.minElem = st.minElem[1:]
	ret := st.elems[0]
	st.elems = st.elems[1:]
	return  ret
}
func (st *MinStack)Top() interface{}{
	return st.elems[0]
}

func (st *MinStack) GetMin() interface{} {
	if len(st.minElem) <= 0{ // 易错点1
		return math.MaxInt32
	}
	return st.minElem[0]
}

// 150. Evaluate Reverse Polish Notation
func EvalRPN(tokens []string) int {
	handler := map[string]func(int,int)int{
		"+": add,
		"-": minus,
		"*": mutiply,
		"/": divid,
	}
	st := []int{}
	ans := 0
	for _,token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			ans = handler[token](st[1], st[0])
			st = append([]int{ans}, st[2:]...)
		}else if ok, _ := regexp.Match("[0-9]+", []byte(token)); ok {
			val, err := strconv.Atoi(token)
			if err != nil {
				return 0
			}
			st = append([]int{val}, st...)
		}else {
			return 0
		}
	}
	if len(st) != 1{
		return 0
	}
	return st[0]
}
func add(fst, snd int) int {
	fmt.Println(fst, snd, "+", fst + snd)
	return fst + snd
}
func divid(fst, snd int) int{
	if snd == 0{
		return 0
	}
	fmt.Println(fst, snd, "/", fst / snd)
	return fst / snd
}
func minus(fst, snd int) int{
	return fst - snd
}
func mutiply(fst, snd int)int {
	fmt.Println(fst, snd, "*")
	return fst * snd
}