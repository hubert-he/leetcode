package Iterator
/* 251. Flatten 2D Vector
** Design an iterator to flatten a 2D vector. It should support the next and hasNext operations.
** Implement the Vector2D class:
		Vector2D(int[][] vec) initializes the object with the 2D vector vec.
		next() returns the next element from the 2D vector and moves the pointer one step forward.
			You may assume that all the calls to next are valid.
		hasNext() returns true if there are still some elements in the vector, and false otherwise.
 */
// 核心点在 advance 上， 考虑vector-向量的列长度不等，特别注意空向量列的出现
// 特别case
// ["Vector2D",    "hasNext","next","hasNext"]
// [[[[-1],[],[]]],[],    [],        []]
// ["Vector2D", "hasNext","next","hasNext"]
// [[[[],[3]]], [],       [],     []]
type Vector2D struct {
	i,j     int
	row     int
	vec     [][]int
}


func ConstructorVertorIter(vec [][]int) Vector2D {
	this := Vector2D{row: len(vec), vec:vec}
	return this
}
// 优化-1 增加一个函数，来处理进入下一个的逻辑
func (this *Vector2D)  advance(){
	//for this.i < this.row && len(this.vec[this.i]) <= 0{
	for this.i < this.row && this.j == len(this.vec[this.i]){
		this.i++
		this.j = 0 // 增加
	}
}

func (this *Vector2D) Next() int {
	// advance 优化， hasNext 函数处理下一个
	if !this.HasNext(){ panic("nil access ")}
	/* 抽象成advance
	for len(this.vec[this.i]) <= 0{
		this.i++
	} */
	ret := this.vec[this.i][this.j]
	this.j++
	/* 抽象成advance
	if this.j >= len(this.vec[this.i]){
		this.i++
		this.j = 0
	}
	 */
	return ret
}


func (this *Vector2D) HasNext() bool {
	/* 通过advance 抽象
	for this.i < this.row && len(this.vec[this.i]) <= 0{
		this.i++
	}
	return this.i < this.row-1 ||( this.i == this.row-1 && this.j < len(this.vec[this.i]))
	 */
	this.advance()
	return this.i < this.row
}


/**
 * Your Vector2D object will be instantiated and called as such:
 * obj := Constructor(vec);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */