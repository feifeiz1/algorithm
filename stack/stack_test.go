package stack

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	st := NewStack[int](10)
	if !st.Empty() {
		t.Fatal("not empty!")
	}
	if st.Cap() != 10 {
		t.Fatal("cap not equal!")
	}
	st.Push(1)
}

func TestStack_Push(t *testing.T) {
	st := NewStack[int](5)
	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.PrintStack()
	fmt.Println()
	if st.Top() != 3 {
		t.Fatal("top is not equal 3!")
	}
	st.Pop()
	st.PrintStack()
	fmt.Println()
	if st.Pop() != 2 {
		t.Fatal("top is not equal 2!")
	}
	st.PrintStack()
}

func TestStack_Push2(t *testing.T) {
	st := NewStack[int](3)
	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.PrintStack()
	fmt.Println()
	fmt.Printf("size:%v,cap:%v\n", st.Len(), st.Cap())
	st.Push(4)
	st.PrintStack()
	fmt.Printf("size:%v,cap:%v\n", st.Len(), st.Cap())

}
