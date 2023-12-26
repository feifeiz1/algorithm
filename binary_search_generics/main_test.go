package main

import (
	"slices"
	"strconv"
	"testing"
)

type MyVal int

func (mv MyVal) String() string {
	return strconv.Itoa(int(mv))
}

var preorderList []string
var inorderList []string
var postorderList []string

func fillMyValBst(bst *BinarySearchTree[int, MyVal]) {
	bst.Insert(8, MyVal(8))
	bst.Insert(4, MyVal(4))
	bst.Insert(10, MyVal(10))
	bst.Insert(2, MyVal(2))
	bst.Insert(6, MyVal(6))
	bst.Insert(1, MyVal(1))
	bst.Insert(3, MyVal(3))
	bst.Insert(5, MyVal(5))
	bst.Insert(7, MyVal(7))
	bst.Insert(9, MyVal(9))
	bst.Insert(11, MyVal(11))

	preorderList = []string{"8", "4", "2", "1", "3", "6", "5", "7", "10", "9", "11"}
	inorderList = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}
	postorderList = []string{"1", "3", "2", "5", "7", "6", "4", "9", "11", "10", "8"}
}

func TestBinarySearchTree_Insert(t *testing.T) {
	bst := &BinarySearchTree[int, MyVal]{}
	fillMyValBst(bst)
	bst.PrintTree()

	bst.Insert(12, MyVal(12))
	bst.PrintTree()
}

func TestBinarySearchTree_PreOrderTraverse(t *testing.T) {
	bst := &BinarySearchTree[int, MyVal]{}
	fillMyValBst(bst)
	var res []string
	bst.PreOrderTraverse(func(val MyVal) {
		res = append(res, val.String())
	})
	if !slices.Equal(res, preorderList) {
		t.Fatal("pre order failed")
	}
}

func TestBinarySearchTree_InOrderTraverse(t *testing.T) {
	bst := &BinarySearchTree[int, MyVal]{}
	fillMyValBst(bst)
	var res []string
	bst.InOrderTraverse(func(val MyVal) {
		res = append(res, val.String())
	})
	if !slices.Equal(res, inorderList) {
		t.Fatal("in order failed")
	}
}

func TestBinarySearchTree_PostOrderTraverse(t *testing.T) {
	bst := &BinarySearchTree[int, MyVal]{}
	fillMyValBst(bst)
	var res []string
	bst.PostOrderTraverse(func(val MyVal) {
		res = append(res, val.String())
	})
	if !slices.Equal(res, postorderList) {
		t.Fatal("post order failed")
	}
}

func TestBinarySearchTree_Search(t *testing.T) {
	bst := &BinarySearchTree[int, MyVal]{}
	fillMyValBst(bst)
	if !bst.Search(10) {
		t.Fatal("search 10 failed")
	}
	if bst.Search(50) {
		t.Fatal("search not exits failed")
	}
}

func TestBinarySearchTree_Max(t *testing.T) {
	bst := &BinarySearchTree[int, MyVal]{}
	fillMyValBst(bst)
	if bst.Max() != MyVal(11) {
		t.Fatal("bst max failed")
	}
}

func TestBinarySearchTree_Min(t *testing.T) {
	bst := &BinarySearchTree[int, MyVal]{}
	fillMyValBst(bst)
	if bst.Min() != MyVal(1) {
		t.Fatal("bst min failed")
	}
}

func TestBinarySearchTree_Remove(t *testing.T) {
	bst := &BinarySearchTree[int, MyVal]{}
	fillMyValBst(bst)
	bst.PrintTree()

	bst.Remove(10)
	bst.PrintTree()

	bst.Remove(4)
	bst.PrintTree()
}
