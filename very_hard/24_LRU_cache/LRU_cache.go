package LRU_cache

type LRUCache struct {
	maxSize int
	// Add fields here.
}

func NewLRUCache(size int) *LRUCache {
	return nil
}

func (cache *LRUCache) InsertKeyValuePair(key string, value int) {
}


func (cache *LRUCache) GetValueFromKey(key string) (int, bool) {
	return -1, false
}

func (cache *LRUCache) GetMostRecentKey() (string, bool) {
	return "", false
}
