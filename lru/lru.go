package lru

import "fmt"

/***

请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。


示例：

输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
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

type node struct {
	k    int
	val  any
	prev *node
	next *node
}

type LRUCache struct {
	head, tail *node
	cap, size  int
	cache      map[int]*node
}

func newNode(k int, v any) *node {
	return &node{
		k:   k,
		val: v,
	}
}

func NewLRUCache(cap int) *LRUCache {
	h, t := newNode(0, 0), newNode(0, 0)
	h.next = t
	t.prev = h
	return &LRUCache{
		head:  h,
		tail:  t,
		cap:   cap,
		size:  0,
		cache: map[int]*node{},
	}
}

func (lc *LRUCache) print() {
	fmt.Printf("size:%v,cap:%v\n", lc.size, lc.cap)
	for cur := lc.head; cur != nil; cur = cur.next {
		if cur == lc.tail {
			fmt.Printf("%v", cur.val)

		} else {
			fmt.Printf("%v->", cur.val)
		}
	}
	fmt.Println()
}

func (lc *LRUCache) Get(k int) (v any) {
	tmp, ok := lc.cache[k]
	if !ok {
		return -1
	}
	//fmt.Println("tmp,k:", tmp.k, "tmp.v:", tmp.val)
	lc.removeNode(tmp)
	lc.moveToHead(tmp)
	return tmp.val
}

func (lc *LRUCache) Put(k int, v any) {
	tmp, ok := lc.cache[k]
	if !ok {
		// 不存在
		tmpNode := newNode(k, v)
		if lc.size >= lc.cap {
			// 淘汰
			delNode := lc.tail.prev
			lc.removeNode(delNode)
			lc.size--
			delete(lc.cache, delNode.k)
			// 插入新的
			lc.moveToHead(tmpNode)
			lc.cache[k] = tmpNode
			lc.size++
		} else {
			lc.moveToHead(tmpNode)
			lc.cache[k] = tmpNode
			lc.size++
		}
	} else {
		// 已经存在
		tmp.val = v
		lc.removeNode(tmp)
		lc.moveToHead(tmp)
		lc.cache[k] = tmp
	}
}

func (lc *LRUCache) moveToHead(n *node) {
	n.next = lc.head.next
	lc.head.next.prev = n
	n.prev = lc.head
	lc.head.next = n
}

func (lc *LRUCache) moveToTail(n *node) {
	lc.removeNode(n)
	lc.tail.prev.next = n
	n.prev = lc.tail.prev
	n.next = lc.tail
	lc.tail.prev = n
}

func (lc *LRUCache) removeNode(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = nil
	n.prev = nil
}

// 这泛型太坑爹了，用起来过于难受，不用了

//type node[K comparable, V any] struct {
//	k    K
//	v    V
//	prev *node[K, V]
//	next *node[K, V]
//}
//
//type linkedList[K comparable, V any] struct {
//	head *node[K, V]
//	tail *node[K, V]
//}
//
//func newLinkedList[K comparable, V any]() *linkedList[K, V] {
//	return &linkedList[K, V]{
//		head: nil,
//		tail: nil,
//	}
//}
//
//func (l *linkedList[K, V]) pushHead(k K, v V) (n *node[K, V]) {
//	n = &node[K, V]{
//		k: k,
//		v: v,
//	}
//	if l.head == nil {
//		l.head, l.tail = n, n
//		return
//	}
//	n.next = l.head
//	l.head.prev = n
//	l.head = n
//	return
//}
//
//func (l *linkedList[K, V]) pushBack(k K, v V) (n *node[K, V]) {
//	n = &node[K, V]{
//		k: k,
//		v: v,
//	}
//	if l.tail == nil {
//		l.head, l.tail = n, n
//		return
//	}
//	l.tail.next = n
//	n.prev = l.tail
//	l.tail = n
//	return
//}
//
//// LRUCache 缓存，最近最少未使用
//type LRUCache[K comparable, V any] struct {
//	sync.RWMutex
//	data *linkedList[K, V]
//	mp   map[K]*node[K, V]
//	cap  int
//	size int
//}
//
//func NewLRUCache[K comparable, V any](cap int) *LRUCache[K, V] {
//	return &LRUCache[K, V]{
//		data: newLinkedList[K, V](),
//		mp:   make(map[K]*node[K, V]),
//		cap:  cap,
//		size: 0,
//	}
//}
//
//func (lc *LRUCache[K, V]) Get(k K) (val V) {
//	lc.Lock()
//	defer lc.Unlock()
//	data, ok := lc.mp[k]
//	if !ok {
//		return
//	}
//	val = data.v
//	// 将此元素放置在列表头
//	lc.data.Del(data)
//	lc.data.In
//}
//
//func (lc *LRUCache[T]) Put(k, v T) {
//	lc.Lock()
//	defer lc.Unlock()
//	// 已经达到容量，需要将最久未使用的删除
//	if lc.size >= lc.cap {
//		lc.data.PopBack()
//		lc.size--
//		return
//	}
//	// 未达到容量，直接头插
//	lc.data.PushHead(v)
//	lc.size++
//	return
//}
