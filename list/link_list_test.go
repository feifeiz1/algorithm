package list

import (
	"fmt"
	"slices"
	"testing"
)

var (
	l1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
)

func TestList_PushBack(t *testing.T) {
	l := NewList[int]()
	l.PrintList(nil)
	for _, v := range l1 {
		l.PushBack(v)
	}

	var res []int
	l.PrintList(func(i int) {
		res = append(res, i)
	})
	if !slices.Equal(res, l1) {
		t.Fatalf("push back not equal!")
	}
	fmt.Println()
}

func TestList_PushHead(t *testing.T) {
	l := NewList[int]()
	for _, v := range l1 {
		l.PushHead(v)
	}
	var res []int
	l.PrintList(func(i int) {
		res = append(res, i)
	})
	slices.Reverse(l1)
	if !slices.Equal(l1, res) {
		t.Fatal("push head not equal!")
	}
	fmt.Println()
}

func TestList_PopBack(t *testing.T) {
	l := NewList[int]()
	for _, v := range l1 {
		l.PushBack(v)
	}
	l.PrintList(nil)
	l.PopBack()
	var res []int
	l.PrintList(func(i int) {
		res = append(res, i)
	})
	if !slices.Equal(l1[:len(l1)-1], res) {
		t.Fatal("pop back not equal!")
	}
	l.PopBack()
	l.PopBack()
	l.PopBack()
	l.PopBack()
	l.PopBack()
	l.PopBack()
	l.PopBack()
	l.PopBack()

	l.PopBack()
	l.PopBack()
	var res1 []int
	l.PrintList(func(i int) {
		res1 = append(res1, i)
	})
	if !slices.Equal(res1, []int{}) {
		t.Fatal("pop back all not equal!")
	}
}

func TestList_PopHead(t *testing.T) {
	l := NewList[int]()
	for _, v := range l1 {
		l.PushBack(v)
	}
	l.PrintList(nil)
	l.PopHead()
	var res []int
	l.PrintList(func(i int) {
		res = append(res, i)
	})
	if !slices.Equal(l1[1:], res) {
		t.Fatal("pop back not equal!")
	}
	l.PopHead()
	l.PopHead()
	l.PopHead()
	l.PopHead()
	l.PopHead()
	l.PopHead()
	l.PopHead()
	l.PopHead()

	l.PopHead()
	l.PopHead()

	var res1 []int
	l.PrintList(func(i int) {
		res1 = append(res1, i)
	})
	if !slices.Equal(res1, []int{}) {
		t.Fatal("pop head all not equal!")
	}
}

func TestList_Find(t *testing.T) {
	l := NewList[int]()
	for _, v := range l1 {
		l.PushBack(v)
	}
	l.PrintList(nil)
	if !l.Find(5) {
		t.Fatal("list find failed!")
	}
	if l.Find(20) {
		t.Fatal("fatal!!! list find a not exists val!")
	}
}

func TestList_Del(t *testing.T) {
	l := NewList[int]()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(5)
	l.PrintList(nil)
	l.Del(2)
	var res []int
	l.PrintList(func(i int) {
		res = append(res, i)
	})
	if !slices.Equal(res, []int{1, 5}) {
		t.Fatal("list del not equal!")
	}
	l.Del(1)
	var res1 []int
	l.PrintList(func(i int) {
		res1 = append(res1, i)
	})
	if !slices.Equal(res1, []int{5}) {
		t.Fatal("list del1 not equal!")
	}
	l.Del(5)
	var res2 []int
	l.PrintList(func(i int) {
		res2 = append(res2, i)
	})
	if !slices.Equal(res2, []int{}) {
		t.Fatal("list del2 not equal!")
	}
}

func TestList_DelIdx(t *testing.T) {
	l := NewList[int]()
	for _, v := range l1 {
		l.PushBack(v)
	}
	l.PrintList(nil)
	l.DelIdx(0)
	l.PrintList(nil)
	l.DelIdx(3)
	l.PrintList(nil)
	l.DelIdx(6)
	l.PrintList(nil)
	l.DelIdx(5)
	l.PrintList(nil)
	l.DelIdx(4)
	l.PrintList(nil)
	if l.Len() != 4 {
		t.Fatal("del idx len not equal 4!")
	}
	l.DelIdx(0)
	l.DelIdx(0)
	l.DelIdx(0)
	l.DelIdx(0)
	l.DelIdx(0)
	l.PrintList(nil)
	if l.Len() != 0 {
		t.Fatal("del idx len not equal 0!")
	}
	l.PushBack(1)
	l.PrintList(nil)
	if l.Len() != 1 {
		t.Fatal("del idx len not equal 1!")
	}
	l.PushHead(2)
	l.PrintList(nil)

	if l.Del(5) {
		t.Fatal("del a not exists item!!")
	}

	if l.DelIdx(9) {
		t.Fatal("del a not exists idx!!")
	}
}
