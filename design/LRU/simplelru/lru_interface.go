// hashicorp 实现的lru，基于golang 内置container/list
package simplelru

// LRUCache is the interface for simple LRU cache.
type LRUCache interface {
	/*	Adds a value to the cache
	   	returns true if an eviction occurred and
		updates the "recently used"-ness of the key.
	 */
	Add(key, value interface{}) bool
	/*	Return key's vaule from the cache and
		update the "recently used"-ness of the key
	 */
	Get(key interface{})(value interface{}, ok bool)
	/*	Checks if a key exists in cache Without updating the recent-ness
		return true if key exists
	 */
	Contains(key interface{})(ok bool)
	/*	return key's value Without updating the "recently used"-ness-ness of the key
	 */
	Peek(key interface{})(value interface{}, ok bool)
	/*
		Removes a key from cache
	 */
	Remove(key interface{}) bool
	/*
		Removes the oldest entry from cache
	 */
	RemoveOldest()(interface{}, interface{}, bool)
	/*
		Returns the oldest entry from the cache. #key value isFound
	 */
	GetOldest()(interface{}, interface{}, bool)
	/*
		Returns a slice of the keys in the cache, from oldest to newest
	 */
	Keys() []interface{}
	/*
		Returns the number of items in the cache
	 */
	Len() int
	/*
		Clears all cache entries
	 */
	Purge()
	// Resizes cache, returning number evicted
	Resize(int) int
}