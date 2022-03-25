package Iterator

/* 341. Flatten Nested List Iterator
** You are given a nested list of integers nestedList.
** Each element is either an integer or a list whose elements may also be integers or other lists.
** Implement an iterator to flatten it.
** Implement the NestedIterator class:
	NestedIterator(List<NestedInteger> nestedList) Initializes the iterator with the nested list nestedList.
	int next() Returns the next integer in the nested list.
	boolean hasNext() Returns true if there are still some integers in the nested list and false otherwise.
** Your code will be tested with the following pseudocode:
		initialize iterator with nestedList
		res = []
		while iterator.hasNext()
			append iterator.next() to the end of res
		return res
If res matches the expected flattened list, then your code will be judged as correct.
 */
type NestedInteger struct {}
func (this NestedInteger) IsInteger() bool {return true }
func (this NestedInteger) GetInteger() int { return 0 }
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
	a	[]*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	a := []*NestedInteger{}
	for i := len(nestedList)-1; i >= 0; i--{
		// [[]] 测试case 不通过的原因：存在 0 值 加 if 判断
		// [[]] 在 nestedList 中nestedList: [*{Value: 0, List: [], IsInt: false}]
		if !nestedList[i].IsInteger() && len(nestedList[i].GetList()) <= 0{
			continue
		}
		a = append(a, nestedList[i])
	}
	this := NestedIterator{a: a}
	return &this
}

func (this *NestedIterator) Next() int {
	q := this.a[len(this.a)-1]
	this.a = this.a[:len(this.a)-1]
	/* 在 HasNext 里 解构 NestedInteger
	for !q.IsInteger(){
		l := q.GetList()
		for i := len(l)-1; i >= 0; i--{
			// [[[[]]],[]] 不通过，原因同上，对NestedInteger 零值的认识
			if !l[i].IsInteger() && len(l[i].GetList()) <= 0{
				continue
			}
			this.a = append(this.a, l[i])
		}
		q = this.a[len(this.a)-1]
		this.a = this.a[:len(this.a)-1]
	}*/
	return q.GetInteger()
}
// 更改记录-1： 题目NestedInteger 结构体 对 零值的定义，需要再 HasNext() 解构
func (this *NestedIterator) HasNext() bool {
	if len(this.a) <= 0{ return false }
	q := this.a[len(this.a)-1]
	this.a = this.a[:len(this.a)-1]
	for !q.IsInteger(){
		l := q.GetList()
		for i := len(l)-1; i >= 0; i--{
			// [[[[]]],[]] 不通过，原因同上，对NestedInteger 零值的认识
			if !l[i].IsInteger() && len(l[i].GetList()) <= 0{
				continue
			}
			this.a = append(this.a, l[i])
		}
		if len(this.a) <= 0{ return false }
		q = this.a[len(this.a)-1]
		this.a = this.a[:len(this.a)-1]
	}
	this.a = append(this.a, q)
	return true
}
