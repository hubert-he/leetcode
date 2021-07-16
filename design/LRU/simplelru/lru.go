package simplelru

import (
	"container/list"
	"errors"
)

/*
	EvictCallback is used to get a callback when a cache entry is evicted
 */
type EvictCallback func(key interface{}, value interface{})

// LRU implements a non-thread safe fixed size LRU cache
type LRU struct {
	size		int			// 固定大小，超过此值，则进行evict，删除最老entry
	evictList	*list.List	// 存放cahce entry的地方
	items		map[interface{}]*list.Element	// 映射关系，根据key快速查找entry
	onEvict		EvictCallback	// evict时的回调函数
}

// entry is used to hold a value in the evictList
type entry struct {
	key		interface{}
	value	interface{}
}

func NewLRU(size int, onEvict EvictCallback)(*LRU, error){
	if size <= 0{
		return nil, errors.New("must provide a positive size")
	}
	c := &LRU{
		size: size,
		evictList: list.New(),
		items: make(map[interface{}]*list.Element),
		onEvict: onEvict,
	}
	return c, nil
}
func (c *LRU) Purge(){
	for k, v := range c.items{
		if c.onEvict != nil{
			c.onEvict(k, v.Value.(*entry).value)
		}
		delete(c.items, k)
	}
	c.evictList.Init()
}
// Returns true if an eviction occurred
func (c *LRU) Add(key, value interface{}) (evicted bool){
	// 1. check for existing item
	if ent, ok := c.items[key]; ok{
		c.evictList.MoveToFront(ent)
		ent.Value.(*entry).value = value
		return false
	}
	// 2. add new item
	entry := c.evictList.PushFront(&entry{key, value})
	c.items[key] = entry
	// 3. Verify size on exceeded
	evict := c.evictList.Len() > c.size
	if evict {
		c.removeOldest()
	}
	return evict
}
// Get looks up a key's value from the cache
func (c *LRU) Get(key interface{}) (value interface{}, ok bool){
	if ent, ok := c.items[key]; ok {
		c.evictList.MoveToFront(ent) // update 生命值
		if ent.Value.(*entry) == nil {
			return nil, false
		}
		return ent.Value.(*entry).value, true
	}
	return
}
/*	Contains checks if a key is in the cache, without updating the recent-ness
	or deleting it for being stale.
 */
func (c *LRU) Contains(key interface{})(ok bool){
	_, ok = c.items[key]
	return
}
func (c *LRU) Peek(key interface{})(value interface{}, ok bool){
	var ent *list.Element
	if ent, ok = c.items[key]; ok {
		return ent.Value.(*entry).value, true
	}
	return
}
// Remove removes the provided key from the cache, returning if the key was containe
func (c *LRU) Remove(key interface{})(present bool){
	if ent, ok := c.items[key]; ok {
		c.removeElement(ent)
		return true
	}
	return false
}

func (c *LRU) RemoveOldest()(key, value interface{}, ok bool){
	ent := c.evictList.Back()
	if ent != nil {
		c.removeElement(ent)
		kv := ent.Value.(*entry)
		return kv.key, kv.value, true
	}
	return
}
func (c *LRU) GetOldest() (key, value interface{}, ok bool){
	ent := c.evictList.Back()
	if ent != nil {
		kv := ent.Value.(*entry)
		return kv.key, kv.value, true
	}
	return
}

func (c *LRU) Keys() []interface{}{
	keys := make([]interface{}, 0, len(c.items))
	for ent := c.evictList.Back(); ent != nil; ent = ent.Prev(){
		keys = append(keys, ent.Value.(*entry).key)
	}
	return keys
}

func (c *LRU) Len() int{
	return c.evictList.Len()
}
// Resize changes the cache size.
func (c *LRU) Resize(size int) (evicted int) {
	diff := c.Len() - size
	if diff < 0{ // 扩张
		diff = 0
	}
	// 缩减
	for i := 0; i < diff; i++{
		c.removeOldest()
	}
	c.size = size
	return diff
}

func (c *LRU) removeOldest() {
	ent := c.evictList.Back()
	if ent != nil {
		c.removeElement(ent)
	}
}
func (c *LRU) removeElement(e *list.Element){
	c.evictList.Remove(e) // 从list中删除
	kv := e.Value.(*entry)
	delete(c.items, kv.key) // 删除映射关系
	if c.onEvict != nil {
		c.onEvict(kv.key, kv.value) // 调用evict回调函数
	}
}