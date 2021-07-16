package lru

import (
	"./simplelru"
	"fmt"
	"sync"
)

const (
	// Default2QRecentRatio is the ratio of the 2Q cache dedicated
	// to recently added entries that have only been accessed once.
	Default2QRecentRatio = 0.25

	// Default2QGhostEntries is the default ratio of ghost
	// entries kept to track entries recently evicted
	Default2QGhostEntries = 0.50
)
type TwoQueueCache struct {
	size		int
	recentSize	int

	recent		simplelru.LRUCache
	frequent	simplelru.LRUCache
	recentEvict	simplelru.LRUCache
	lock		sync.RWMutex
}
func New2Q(size int) (*TwoQueueCache, error){
	return New2QParams(size, Default2QRecentRatio, Default2QGhostEntries)
}
func New2QParams(size int, recentRatio, ghostRatio float64)( *TwoQueueCache, error){
	if size <= 0{
		return nil, fmt.Errorf("invalid size")
	}
	if recentRatio < 0.0 || recentRatio > 1.0{
		return nil, fmt.Errorf("invalid recent ratio")
	}
	if ghostRatio < 0.0 || ghostRatio > 1.0{
		return nil, fmt.Errorf("Invalid ghost ratio")
	}
	// Determine the sub-sizes
	recentSize := int(float64(size) * recentRatio)
	evictSize := int(float64(size) * ghostRatio)
	// allocate LRUs
	recent, err := simplelru.NewLRU(size, nil)
	if err != nil {
		return nil, err
	}
	frequent, err := simplelru.NewLRU(size, nil)
	if err != nil {
		return nil, err
	}
	recentEvict, err := simplelru.NewLRU(evictSize, nil)
	if err != nil {
		return nil, err
	}
	// initialize the cache
	c := &TwoQueueCache{
		size:			size,
		recentSize:		recentSize,
		recent:			recent,
		frequent: 		frequent,
		recentEvict: 	recentEvict,
	}
	return c, nil
}

func (c *TwoQueueCache) Get(key interface{})(value interface{}, ok bool){
	c.lock.Lock()
	defer c.lock.Unlock()
	// Check if this is a frequent value
	if val, ok := c.frequent.Get(key); ok {
		return val, ok
	}
	// If the value is contained in recent, then we promote it to frequent
	if val, ok := c.recent.Peek(key); ok {
		c.recent.Remove(key)
		c.frequent.Add(key, val)
		return val, ok
	}
	// No Hit
	return nil, false
}

func (c *TwoQueueCache) Add(key, value interface{}){
	c.lock.Lock()
	defer c.lock.Unlock()
	// Check if the value is frequently used already, and just update the value
	if c.frequent.Contains(key){
		c.frequent.Add(key, value)
		return
	}
	// Check if the value is recently used, and promote the value into the frequent list
	if c.recent.Contains(key){
		c.recent.Remove(key)
		c.frequent.Add(key, value)
		return
	}
	// If the value was recently evicted, add it to the frequently used list
	if c.recentEvict.Contains(key) {
		c.ensureSpace(true) // ensureSpace is used to ensure we have space in the cache
		c.recentEvict.Remove(key)
		c.frequent.Add(key, value)
		return
	}
	// Add to the recent seen list
	c.ensureSpace(false) // ensureSpace is used to ensure we have space in the cache
	c.recent.Add(key, value)
}
// ensureSpace is used to ensure we have space in the cache
func (c *TwoQueueCache) ensureSpace(recentEvict bool){
	// If we have space, nothing to do
	recentLen := c.recent.Len()
	freqLen := c.frequent.Len()
	if recentLen + freqLen < c.size{
		return
	}
	// If the recent buffer is larger than the target, evict from there
	if recentLen > 0 && (recentLen > c.recentSize || (recentLen == c.recentSize && !recentEvict)) {
		k, _, _ := c.recent.RemoveOldest()
		c.recentEvict.Add(k, nil)
		return
	}
	// Remove from the frequent list otherwise
	c.frequent.RemoveOldest()
}

func (c *TwoQueueCache)Len() int{
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.recent.Len() + c.frequent.Len()
}

func (c *TwoQueueCache) Keys() []interface{}{
	c.lock.RLock()
	defer c.lock.RUnlock()
	k1 := c.frequent.Keys()
	k2 := c.recent.Keys()
	return append(k1,k2...)
}

func (c *TwoQueueCache) Remove(key interface{}){
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.frequent.Remove(key) {
		return
	}
	if c.recent.Remove(key){
		return
	}
	if c.recentEvict.Remove(key){
		return
	}
}

func (c *TwoQueueCache)Purge() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.recent.Purge()
	c.frequent.Purge()
	c.recentEvict.Purge()
}

func (c *TwoQueueCache) Contains(key interface{}) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.frequent.Contains(key) || c.recent.Contains(key)
}

func (c *TwoQueueCache) Peek(key interface{})(value interface{}, ok bool){
	c.lock.RLock()
	defer c.lock.RUnlock()
	if val, ok := c.frequent.Peek(key); ok {
		return val, ok
	}
	return c.recent.Peek(key)
}











