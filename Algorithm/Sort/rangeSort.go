package Sort
import "github.com/emirpasic/gods/trees/redblacktree"
/*
** You are implementing a program to use as your calendar.
** We can add a new event if adding the event will not cause a double booking.
** A double booking happens when two events have some non-empty intersection(i.e., some moment is common to both events.).
** The event can be represented as a pair of integers start and end that represents a booking on the half-open interval [start, end),
** the range of real numbers x such that start <= x < end.
** Implement the MyCalendar class:
	MyCalendar() Initializes the calendar object.
	boolean book(int start, int end) Returns true
		if the event can be added to the calendar successfully without causing a double booking.
		Otherwise, return false and do not add the event to the calendar.
 */
/* 解题的常见问题
** 1. 区间的边界是否属于重叠，即[1,10] 与 [10,20] 共有 10 是否属于重叠
** 2. 插排时，插入的是最大的元素，for 循环后 记得要插入值
*/
type MyCalendar struct {
	data    [][2]int
}

func Constructor() MyCalendar {
	this := MyCalendar{}
	return this
}

func (this *MyCalendar) Book(start int, end int) bool {
	if len(this.data) == 0{
		this.data = append(this.data, [2]int{start, end})
		return true
	}
	isOverlap := func(x, y [2]int)bool{
		// 边界重叠的条件
		if (x[0] <= y[0] && x[1] > y[0]) ||
			(y[0] <= x[0] && x[0] < y[1]){
			return true
		}
		return false
	}
	//fmt.Println(this.data)
	for i := range this.data{
		if isOverlap(this.data[i], [2]int{start, end}){
			//fmt.Println(start, end)
			return false
		}
		if end <= this.data[i][0]{
			this.data = append(this.data[:i], append([][2]int{[2]int{start, end}}, this.data[i:]...)...)
			return true
		}
	}
	// 2. 插排问题
	this.data = append(this.data, [2]int{start, end})
	return true
}
/* 方法二 平衡树
**
 */
type Calendar struct {
	t *redblacktree.Tree
}

func ConstructorCalendar() Calendar {
	// NewWithIntComparator instantiates a red-black tree with the IntComparator, i.e. keys are of type int.
	t := redblacktree.NewWithIntComparator()
	//Put inserts node into the tree. Key should adhere to the comparator's type assertion, otherwise method panics.
	t.Put(-1, -1) // 哨兵
	return Calendar{t}
}
/* overLap 的2中可能
**       |_______|     待插入值，情况-1
**  |________|		   已在队列中值， 情况-1

**       |________|			待插入值，情况-2
**			  |_________|   已在队列中值， 情况-2
 */
// 平衡树中，使用区间的start 作为 key， end 作为 value
func (c *Calendar) Book(start, end int) bool {
	//Floor node is defined as the largest node that is smaller than or equal to the given node.
	floor, _ := c.t.Floor(start)
	if floor.Value.(int) > start { //情况-1 [start,end) 左侧区间的右端点超过了 start
		return false
	}
	// IteratorAt 迭代器，从floor的后继元素开始迭代后继
	if it := c.t.IteratorAt(floor); it.Next() && it.Key().(int) < end {
		// 情况-2 [start,end) 右侧区间的左端点小于 end
		return false
	}
	c.t.Put(start, end) // 可以插入区间 [start,end)
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */