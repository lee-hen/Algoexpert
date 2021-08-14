package LRU_cache
// important question

import ll "github.com/lee-hen/Algoexpert/medium/22_linked_list"

type LRUCache struct {
	index            map[string]*ll.Node
	maxSize          int
	currentSize      int
	listOfMostRecent *ll.DoublyLinkedList
}

func NewLRUCache(size int) *LRUCache {
	lru := &LRUCache{
		index:            map[string]*ll.Node{},
		maxSize:          size,
		currentSize:      0,
		listOfMostRecent: &ll.DoublyLinkedList{},
	}
	if lru.maxSize < 1 {
		lru.maxSize = 1
	}
	return lru
}

// InsertKeyValuePair
// O(1) time | O(1) space
func (cache *LRUCache) InsertKeyValuePair(key string, value int) {
	if _, found := cache.index[key]; !found {
		if cache.currentSize == cache.maxSize {
			cache.evictLeastRecent()
		} else {
			cache.currentSize += 1
		}
		cache.index[key] = &ll.Node{
			Key:   key,
			Value: value,
		}
	} else {
		cache.replaceKey(key, value)
	}
	cache.updateMostRecent(cache.index[key])
}

// GetValueFromKey
// O(1) time | O(1) space
func (cache *LRUCache) GetValueFromKey(key string) (int, bool) {
	if node, found := cache.index[key]; !found {
		return 0, false
	} else {
		cache.updateMostRecent(node)
		return node.Value, true
	}
}

// GetMostRecentKey
// O(1) time | O(1) space
func (cache *LRUCache) GetMostRecentKey() (string, bool) {
	if cache.listOfMostRecent.Head == nil {
		return "", false
	}
	return cache.listOfMostRecent.Head.Key, true
}

func (cache *LRUCache) evictLeastRecent() {
	key := cache.listOfMostRecent.Tail.Key
	cache.listOfMostRecent.RemoveTail()
	delete(cache.index, key)
}

func (cache *LRUCache) updateMostRecent(node *ll.Node) {
	cache.listOfMostRecent.SetHeadTo(node)
}

func (cache *LRUCache) replaceKey(key string, value int) {
	if node, found := cache.index[key]; !found {
		panic("The provided key isn't in the cache!")
	} else {
		node.Value = value
	}
}
