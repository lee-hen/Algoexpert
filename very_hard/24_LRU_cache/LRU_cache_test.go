package LRU_cache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	lruCache := NewLRUCache(3)
	lruCache.InsertKeyValuePair("b", 2)
	lruCache.InsertKeyValuePair("a", 1)
	lruCache.InsertKeyValuePair("c", 3)
	key, _ := lruCache.GetMostRecentKey()
	require.True(t, key == "c")
	value, _ := lruCache.GetValueFromKey("a")
	require.True(t, value == 1)
	key, _ = lruCache.GetMostRecentKey()
	require.True(t, key == "a")
	lruCache.InsertKeyValuePair("d", 4)
	_, found := lruCache.GetValueFromKey("b")
	require.True(t, !found)
	lruCache.InsertKeyValuePair("a", 5)
	value, _ = lruCache.GetValueFromKey("a")
	require.True(t, value == 5)
}