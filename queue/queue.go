package queue

import (
	"cmp"
	"github.com/shao1f/algorithm/list"
)

type Queue[T cmp.Ordered] struct {
	data *list.List[T]
}

func NewQueue[T cmp.Ordered]() *Queue[T] {
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

func (q *Queue[T]) LPop() (val T) {
	if q.data == nil {
		return val
	}
	return q.data.PopHead()
}
