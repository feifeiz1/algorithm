package queue

import "testing"

func TestNewQueue(t *testing.T) {
	q := NewQueue[int]()
	q.LPush(1)
	q.LPush(2)
	q.PrintQueue()
	q.LPop()
	q.PrintQueue()
	q.LPop()
	q.PrintQueue()
	q.LPop()
	q.PrintQueue()
}
