package list

import (
	"fmt"
	"sync"
)

type dnode[T comparable] struct {
	val  T
	prev *dnode[T]
	next *dnode[T]
}

// DLinkList is a double link list
type DLinkList[T comparable] struct {
	sync.RWMutex
	head *dnode[T]
	tail *dnode[T]
	size int
}

// NewDLinkList returns a new double link list
func NewDLinkList[T comparable]() *DLinkList[T] {
	return &DLinkList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (dl *DLinkList[T]) PrintList(f func(T)) {
	dl.RLock()
	defer dl.RUnlock()
	fmt.Printf("---------------------size:%v-------------------------------\n", dl.size)
	for cur := dl.head; cur != nil; cur = cur.next {
		if f != nil {
			f(cur.val)
		}
		if cur == dl.tail {
			fmt.Printf("%v\n", cur.val)
		} else {
			fmt.Printf("%v<->", cur.val)
		}
	}
	fmt.Println("-----------------------------------------------------------")
}

func (dl *DLinkList[T]) PushBack(val T) {
	dl.Lock()
	defer dl.Unlock()
	n := &dnode[T]{
		val: val,
	}
	if dl.tail == nil {
		dl.head, dl.tail = n, n
		dl.size++
		return
	}
	n.prev = dl.tail
	dl.tail.next = n
	dl.tail = n
	dl.size++
}

func (dl *DLinkList[T]) PushHead(val T) {
	dl.Lock()
	defer dl.Unlock()
	n := &dnode[T]{
		val: val,
	}
	if dl.head == nil {
		dl.head, dl.tail = n, n
		dl.size++
		return
	}
	n.next = dl.head
	dl.head.prev = n
	dl.head = n
	dl.size++
}

func (dl *DLinkList[T]) PopHead() (val T) {
	dl.Lock()
	defer dl.Unlock()

	if dl.head == nil {
		return
	}

	return dl.popHead()
}

func (dl *DLinkList[T]) popHead() (val T) {
	val = dl.head.val
	if dl.head == dl.tail {
		dl.head, dl.tail = nil, nil
		dl.size = 0
		return
	}
	tmp := dl.head
	dl.head = tmp.next
	tmp.next = nil
	dl.head.prev = nil
	dl.size--

	tmp = nil
	return
}

func (dl *DLinkList[T]) PopBack() (val T) {
	dl.Lock()
	defer dl.Unlock()
	if dl.tail == nil {
		return
	}
	return dl.popBack()
}

func (dl *DLinkList[T]) popBack() (val T) {
	val = dl.tail.val
	if dl.head == dl.tail {
		dl.head, dl.tail = nil, nil
		dl.size = 0
		return
	}
	tmp := dl.tail
	dl.tail = tmp.prev
	dl.tail.next = nil
	tmp.prev = nil
	dl.size--
	tmp = nil
	return
}

func (dl *DLinkList[T]) Find(val T) bool {
	dl.RLock()
	defer dl.RUnlock()
	if dl.head == nil {
		return false
	}

	return findNode(dl.head, val) != nil
}

func findNode[T comparable](head *dnode[T], val T) *dnode[T] {
	for cur := head; cur != nil; cur = cur.next {
		if cur.val == val {
			return cur
		}
	}
	return nil
}

func (dl *DLinkList[T]) Del(val T) bool {
	dl.Lock()
	defer dl.Unlock()
	if dl.head == nil {
		return false
	}
	n := findNode(dl.head, val)
	return dl.delNode(n)
}

func (dl *DLinkList[T]) delNode(n *dnode[T]) bool {
	if n == nil {
		return false
	}
	if n == dl.head {
		return dl.popHead() == n.val
	}
	if n == dl.tail {
		return dl.popBack() == n.val
	}
	n.prev.next = n.next
	n.next.prev = n.prev
	n.prev = nil
	n.next = nil
	dl.size--
	return true
}

func (dl *DLinkList[T]) Len() int {
	dl.RLock()
	defer dl.RUnlock()
	return dl.size
}

func (dl *DLinkList[T]) DelIdx(idx int) bool {
	dl.RLock()
	defer dl.RUnlock()
	if dl.head == nil {
		return false
	}
	if dl.size < idx {
		return false
	}
	cur := dl.head
	for i := 0; i < idx; i++ {
		cur = cur.next
	}
	return dl.delNode(cur)
}
