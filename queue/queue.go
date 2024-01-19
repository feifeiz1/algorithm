package queue

import "github.com/feifeiz1/algorithm/list"

type Queue[T comparable] struct {
	data *list.List[T]
}

func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{
		data: list.NewList[T](),
	}
}

func (q *Queue[T]) PrintQueue() {
	q.data.PrintList(nil)
}

func (q *Queue[T]) LPush(x T) {
	if q.data == nil {
		q.data = list.NewList[T]()
	}
	q.data.PushHead(x)
}

func (q *Queue[T]) RPop() (val T) {
	if q.data == nil {
		return val
	}
	return q.data.PopBack()
}

func (q *Queue[T]) Empty() bool {
	if q.data == nil {
		return true
	}
	return q.data.Len() == 0
}
