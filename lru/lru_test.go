package lru

import (
	"fmt"
	"testing"
)

func TestNewLRUCache(t *testing.T) {
	lc := NewLRUCache(2)

	lc.print()

	fmt.Println(lc.Get(1))

	lc.Put(1, 1)
	lc.print()
	fmt.Println(lc.Get(1))

	lc.Put(2, 2)
	lc.print()
	fmt.Println(lc.Get(1))
	fmt.Println(lc.Get(2))

	lc.Put(3, 3)
	lc.print()
	fmt.Println(lc.Get(1))
	fmt.Println(lc.Get(2))
	fmt.Println(lc.Get(3))

}

/**
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4
*/

func TestNewLRUCacheGetPut(t *testing.T) {
	lc := NewLRUCache(2)
	lc.print()

	lc.Put(1, 1)
	lc.Put(2, 2)
	lc.print()

	fmt.Println(lc.Get(1))

	lc.Put(3, 3)
	lc.print()
	fmt.Println(lc.Get(2))

	lc.Put(4, 4)
	lc.print()

	fmt.Println(lc.Get(1))
	fmt.Println(lc.Get(3))
	fmt.Println(lc.Get(4))
}
