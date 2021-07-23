package LRU_cache

import ll "github.com/lee-hen/Algoexpert/medium/22_linked_list"

type lruCache struct {
	maxSize int
	ll *ll.DoublyLinkedList
	nodes map[string]*ll.Node
}


// NewLRUCache(3) -> instantiate an LRUCache of size 3
// InsertKeyValuePair("b", 2)
// InsertKeyValuePair("a", 1)
// InsertKeyValuePair("c", 3)
// GetMostRecentKey()  -> c -> "c" was the most recently inserted key
// GetValueFromKey("a") -> 1
// GetMostRecentKey() -> a -> "a" was the most recently retrieved key
// InsertKeyValuePair("d", 4) -> the cache had 3 entries; the least recently used one is evicted
// GetValueFromKey("b") -> not found  -> "b" was evicted in the previous operation
// InsertKeyValuePair("a", 5) -> "a" already exists in the cache so its value just gets replaced
// GetValueFromKey("a") -> 5

func NewLruCache(size int) *lruCache {
	return &lruCache{
		maxSize: size,
		ll: ll.NewDoublyLinkedList(),
		nodes: make(map[string]*ll.Node),
	}
}

func (cache *lruCache) InsertKeyValuePair(key string, value int) {
	newNode := ll.NewNode(key, value)

	if node, ok := cache.nodes[key]; ok {
		cache.ll.Remove(node)
		delete(cache.nodes, node.Key)
	} else if len(cache.nodes) == cache.maxSize {
		node = cache.ll.Tail
		cache.ll.RemoveTail()
		delete(cache.nodes, node.Key)
	}

	cache.ll.SetHead(newNode)
	cache.nodes[key] = newNode
}

func (cache *lruCache) GetValueFromKey(key string) (int, bool) {
	if node, ok := cache.nodes[key]; ok {
		cache.InsertKeyValuePair(key, node.Value)
		return node.Value, true
	}
	return -1, false
}

func (cache *lruCache) GetMostRecentKey() (string, bool) {
	node := cache.ll.Head
	if node != nil {
		return node.Key, true
	}
	return "", false
}
