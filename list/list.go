package list

import (
	"cmp"
	"fmt"
	"sync"
)

type node[T comparable] struct {
	val  T
	next *node[T]
}

type List[T comparable] struct {
	sync.RWMutex
	root *node[T]
	tail *node[T]
	size int
}

func NewList[T cmp.Ordered]() *List[T] {
	return &List[T]{
		root: nil,
		tail: nil,
		size: 0,
	}
}

func (l *List[T]) PrintList(f func(T)) {
	l.RLock()
	defer l.RUnlock()
	fmt.Printf("---------------------size:%v-------------------------------\n", l.size)
	for cur := l.root; cur != nil; cur = cur.next {
		if f != nil {
			f(cur.val)
		}
		if cur == l.tail {
			fmt.Printf("%v\n", cur.val)
		} else {
			fmt.Printf("%v->", cur.val)
		}
	}
	fmt.Println("-----------------------------------------------------------")
}

func (l *List[T]) PushBack(x T) {
	l.Lock()
	defer l.Unlock()
	n := &node[T]{
		val:  x,
		next: nil,
	}
	if l.root == nil {
		l.root = n
		l.tail = n
		l.size++
		return
	}
	l.tail.next = n
	l.tail = n
	l.size++
}

func (l *List[T]) PushHead(x T) {
	l.Lock()
	defer l.Unlock()
	n := &node[T]{
		val:  x,
		next: nil,
	}
	if l.root == nil {
		l.root = n
		l.tail = n
		l.size++
		return
	}
	n.next = l.root
	l.root = n
	l.size++
}

func (l *List[T]) PopBack() (val T) {
	l.Lock()
	defer l.Unlock()
	if l.tail == nil {
		return
	}
	val = l.tail.val
	if l.root == l.tail {
		l.root, l.tail = nil, nil
		l.size--
		return
	}
	cur := l.root
	for ; cur.next != l.tail; cur = cur.next {
	}
	cur.next = nil
	l.tail = cur
	l.size--
	return
}

func (l *List[T]) PopHead() (val T) {
	l.Lock()
	defer l.Unlock()
	if l.root == nil {
		return
	}
	val = l.root.val
	if l.root == l.tail {
		l.root, l.tail = nil, nil
		l.size--
		return
	}
	l.root = l.root.next
	l.size--
	return
}

func (l *List[T]) Find(x T) bool {
	l.RLock()
	defer l.RUnlock()
	if l.root == nil {
		return false
	}
	for cur := l.root; cur != nil; cur = cur.next {
		if cur.val == x {
			return true
		}
	}
	return false
}

func (l *List[T]) Del(x T) bool {
	l.Lock()
	defer l.Unlock()
	if l.root == nil {
		return false
	}

	dummy := &node[T]{}
	dummy.next = l.root
	for cur := dummy; cur != nil; cur = cur.next {
		if cur.next != nil && cur.next.val == x {
			cn := cur.next
			if cn == l.tail {
				l.tail = cur
			}
			cur.next = cn.next
			l.root = dummy.next
			cn = nil
			l.size--
			return true
		}
	}
	return false
}

func (l *List[T]) Len() int {
	l.RLock()
	defer l.RUnlock()
	return l.size
}

func (l *List[T]) DelIdx(idx int) bool {
	l.Lock()
	if idx >= l.size {
		l.Unlock()
		return false
	}
	cur := l.root
	for i := 0; i < idx; i++ {
		cur = cur.next
	}
	l.Unlock()
	l.Del(cur.val)
	return false
}
