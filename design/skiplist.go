package design

import (
	"math"
	"math/rand"
)

/* 1206 Design SkipList
Design a Skiplist without using any built-in libraries.
A skiplist is a data structure that takes O(log(n)) time to add, erase and search.
Comparing with treap(树堆) and red-black tree which has the same function and performance,
the code length of Skiplist can be comparatively short and the idea behind Skiplists is just simple linked lists.
For example, we have a Skiplist containing [30,40,50,60,70,90] and we want to add 80 and 45 into it.
The Skiplist works this way:
You can see there are many layers in the Skiplist. Each layer is a sorted linked list.
With the help of the top layers, add, erase and search can be faster than O(n).
It can be proven that the average time complexity for each operation is O(log(n)) and space complexity is O(n).
See more about Skiplist: https://en.wikipedia.org/wiki/Skip_list
Implement the Skiplist class:
	Skiplist() Initializes the object of the skiplist.
	bool search(int target) Returns true if the integer target exists in the Skiplist or false otherwise.
	void add(int num) Inserts the value num into the SkipList.
	bool erase(int num) Removes the value num from the Skiplist and returns true.
		If num does not exist in the Skiplist, do nothing and return false.
		If there exist multiple num values, removing any one of them is fine.
Note that duplicates may exist in the Skiplist, your code needs to handle this situation.
样例:

Skiplist skiplist = new Skiplist();

skiplist.add(1);
skiplist.add(2);
skiplist.add(3);
skiplist.search(0);   // 返回 false
skiplist.add(4);
skiplist.search(1);   // 返回 true
skiplist.erase(0);    // 返回 false，0 不在跳表中
skiplist.erase(1);    // 返回 true
skiplist.search(1);   // 返回 false，1 已被擦除
约束条件:

0 <= num, target <= 20000
最多调用 50000 次 search, add, 以及 erase操作。
不同跳表的实现，有的用数组方式表示上下层的关系
还有 只定义right和down两个方向的链表
对于跳表以及跳表的同类竞争产品：红黑树，
为啥Redis的有序集合(zset) 使用跳表呢？
因为跳表除了查找插入维护和红黑树有着差不多的效率，
它是个链表，能确定范围区间，而区间问题在树上可能就没那么方便查询啦。
JDK中跳跃表ConcurrentSkipListSet(key 可以相同)
和ConcurrentSkipListMap(key 唯一)。
*/
const MAX_LEVEL = 32

type SkipNode struct {
	key 	int
	value	interface{}
	right	*SkipNode
	down	*SkipNode
}
type Skiplist struct {
	head	*SkipNode
	level	int
	random	int
}


func ConstructorSkipList() Skiplist {
	head := SkipNode{key: math.MaxInt32}
	return Skiplist{head: &head, level: 0, random: 0}
}


func (this *Skiplist) Search(target int) bool {
	cur := this.head
	for cur != nil {
		if target < cur.key{
			return false
		}
		if cur.key == target{
			return true
		}
		if cur.right == nil || cur.right.key > target{
			cur = cur.down
		}else {
			cur = cur.right
		}
	}
	return false
}

/* 插入需要考虑是否插入索引，插入几层等问题
  由于需要插入删除所以 肯定无法维护一个完全理想的索引结构，因为耗费代价太高
  但我们使用随机化的方法去判断是否向上层插入索引。即产生一个[0,1]的随机数，如果小于0.5就向上插入索引
  插入完毕后再次使用随机数判断是否向上插入索引。运气好这个值可能有多层索引， 不好只能插入最底层（无直接索引）
  索引不能无限制增加，因为高度决定查询效率，如果超过最大高度则不再继续添加索引。
  1. 找待插入的左节点
  2. 插入最底层，注意处理是链尾情况
  3. 考虑上一层是否插入索引。
     3.1 判断当前层级，不超过最高则继续
     3.2 设置一个0.5的概率向上插入一层索引（理想状态时每2个向上建立一个索引节点）
  4. 如何找到上层节点： 借助查询过程中记录下降的节点，曾经下降的节点倒序就是需要插入的节点
  5. 如果该层是目前的最高层索引，需要继续向上建立索引应该怎么办
     跳表的head需要改变了，新建一个ListNode节点作为新的head，将它的down指向老head，
	 将这个head节点加入栈中(也就是这个节点作为下次后面要插入的节点)，
 */
func (this *Skiplist) Add(num int)  {
	node := SkipNode{key: num}
	st := []*SkipNode{}
	cur := this.head
	for cur != nil { // 借助查询过程中记录下降的节点，曾经下降的节点倒序就是需要插入的节点,利用栈实现
		if cur.right == nil {
			cur = cur.down
		}else if cur.right.key == num{
			cur.right.value = node.value
		}else if cur.right.key > num {
			st = append([]*SkipNode{cur}, st...)
			cur = cur.down
		}else{
			cur = cur.right
		}
	}
	level := 1 // 当前层数，从第一层添加
	var downNode *SkipNode
	for len(st) != 0{
		cur = st[0]
		st = st[1:]
		newSkipNode := &SkipNode{key: node.key, value: node.value}
		//1. 开始处理垂直方向
		newSkipNode.down = downNode // 初始为nil
		downNode = newSkipNode  // down 方向的后继
		// 2. 开始处理水平方向, 插入在两者之间
		if cur.right == nil { // 右侧为nil 说明插入位置在末尾
			cur.right = newSkipNode
		}else{
			newSkipNode.right = cur.right
			cur.right = newSkipNode
		}
		// 3. 决定是否需要往上创建索引
		if level > MAX_LEVEL{
			break
		}
		randomNum := rand.Float64() // [0,1] 随机数
		if randomNum > 0.5{
			break
		}
		level++
		if level > this.level{//比当前最大高度要高但是依然在允许范围内 需要改变head节点
			this.level = level
			newHeadNode := &SkipNode{key: math.MinInt32}
			newHeadNode.down = this.head
			this.head = newHeadNode
			st = append([]*SkipNode{this.head}, st...)
		}
	}
}

/* (1)删除当前节点和这个节点的前后节点都有关系，需要拿到前一个节点
	不直接判断和操作节点，先找到待删除节点的左侧节点
   (2)删除当前层节点之后，下一层该key的节点也要删除，一直删除到最底层为nil 结束
    删除操作： 每层索引如果有 删除就可以了
 */
func (this *Skiplist) Erase(num int) bool {
	ret := false
	cur := this.head
	for cur != nil{
		if cur.right == nil {
			cur = cur.down
		}else if cur.right.key == num { // cur 为待删除前一个节点
			ret = true
			cur.right = cur.right.right // 删除右侧节点
			cur = cur.down // 向下继续删除
		}else if cur.right.key > num {
			cur = cur.down
		}else{
			cur = cur.right
		}
	}
	return ret
}


/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */