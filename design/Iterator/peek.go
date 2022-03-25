package Iterator
/* 284. Peeking Iterator
** Design an iterator that supports the peek operation on an existing iterator
** in addition to the hasNext and the next operations.
** Implement the PeekingIterator class:
	PeekingIterator(Iterator<int> nums) Initializes the object with the given integer iterator iterator.
	int next() Returns the next element in the array and moves the pointer to the next element.
	boolean hasNext() Returns true if there are still elements in the array.
	int peek() Returns the next element in the array without moving the pointer.
** Follow up: How would you extend your design to be generic and work with all types, not just integer?
 */
/* 由于迭代器运行一次即变动，所有提前缓存一个，作为下一个
** Think of "looking ahead". You want to cache the next element
 */
/*   Below is the interface for Iterator, which is already defined for you.
 */
    type Iterator struct {

    }

    func (this *Iterator) hasNext() bool {
 		// Returns true if the iteration has more elements.
    	return true
   }

    func (this *Iterator) next() int {
		// Returns the next element in the iteration.
    	return 0
   }


type PeekingIterator struct {
	ahead   interface{}
	p       *Iterator
}

func Constructor(iter *Iterator) *PeekingIterator {
	this := PeekingIterator{}
	if iter.hasNext(){
		this.ahead = iter.next()
	}
	this.p = iter
	return &this
}

func (this *PeekingIterator) hasNext() bool {
	return this.ahead != nil
}

func (this *PeekingIterator) next() int {
	ret := this.ahead.(int)
	this.ahead = nil
	if this.p.hasNext(){
		this.ahead = this.p.next()
	}
	return ret
}

func (this *PeekingIterator) peek() int {
	if this.ahead == nil{
		return -1
	}
	return this.ahead.(int)
}
