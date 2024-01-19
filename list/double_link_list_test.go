package list

import (
	"slices"
	"testing"
)

func TestDLinkList_Push(t *testing.T) {
	dl := NewDLinkList[int]()
	dl.PrintList(nil)
	dl.PushBack(1)
	dl.PushBack(2)
	dl.PushHead(3)
	dl.PrintList(nil)
}

func TestDLinkList_Pop(t *testing.T) {
	dl := NewDLinkList[int]()
	var emptyRes []int
	dl.PrintList(func(i int) {
		emptyRes = append(emptyRes, i)
	})
	if !slices.Equal(emptyRes, []int{}) {
		t.Fatal("not empty dLinklist")
	}
	dl.PushHead(1)
	dl.PushHead(2)
	dl.PushHead(3)
	dl.PushHead(4)
	dl.PushBack(5)
	var pushRes []int
	dl.PrintList(func(i int) {
		pushRes = append(pushRes, i)
	})
	if !slices.Equal(pushRes, []int{4, 3, 2, 1, 5}) {
		t.Fatal("push res not equal!")
	}

	dl.PopHead()
	var popRes1 []int
	dl.PrintList(func(i int) {
		popRes1 = append(popRes1, i)
	})
	if !slices.Equal(popRes1, []int{3, 2, 1, 5}) {
		t.Fatal("pop res1 not equal!")
	}

	dl.PopBack()
	var popRes2 []int
	dl.PrintList(func(i int) {
		popRes2 = append(popRes2, i)
	})
	if !slices.Equal(popRes2, []int{3, 2, 1}) {
		t.Fatal("pop res1 not equal!")
	}
	dl.PopHead()
	var popRes3 []int
	dl.PrintList(func(i int) {
		popRes3 = append(popRes3, i)
	})
	if !slices.Equal(popRes3, []int{2, 1}) {
		t.Fatal("pop res1 not equal!")
	}

	dl.PopBack()
	dl.PopBack()
	dl.PopHead()
	dl.PopHead()
	dl.PopBack()
	var popEmpty []int
	dl.PrintList(func(i int) {
		popEmpty = append(popEmpty, i)
	})
	if !slices.Equal(popEmpty, []int{}) {
		t.Fatal("pop empty not equal!")
	}
}

func TestDLinkList_Del(t *testing.T) {
	dl := NewDLinkList[int]()

	dl.PushHead(1)
	dl.PushHead(2)
	dl.PushHead(3)
	dl.PushHead(4)
	dl.PushBack(5)

	dl.PrintList(nil)
	if dl.Len() != 5 {
		t.Fatal("dl len != 5")
	}

	if !dl.Del(3) {
		t.Fatal("del res not true!!!")
	}
	if dl.Len() != 4 {
		t.Fatal("before del,dl len != 4")
	}

	var delRes []int
	dl.PrintList(func(i int) {
		delRes = append(delRes, i)
	})
	if !slices.Equal(delRes, []int{4, 2, 1, 5}) {
		t.Fatal("del 3 error,!")
	}

	if dl.Del(3) {
		t.Fatal("del a not exists val err!!")
	}

	dl.Del(1)
	dl.Del(2)
	dl.Del(4)
	dl.Del(5)
	dl.Del(6)
	dl.PrintList(nil)
}
