package queue

import "testing"

func TestNewQueue(t *testing.T) {
	q := NewQueue[int]()
	q.LPush(1)
	q.LPush(2)
	q.PrintQueue()
	q.RPop()
	q.PrintQueue()
	q.RPop()
	q.PrintQueue()
	q.RPop()
	q.PrintQueue()
	q.LPush(3)
	q.LPush(4)
	q.LPush(5)
	q.LPush(6)
	q.PrintQueue()
	q.RPop()
	q.PrintQueue()
	q.RPop()
	q.PrintQueue()
	q.RPop()
	q.PrintQueue()
	q.RPop()
	q.PrintQueue()
	q.RPop()
	q.PrintQueue()

}
