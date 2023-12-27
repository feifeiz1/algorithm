package binary_tree

import (
	"fmt"
	"slices"
	"testing"
)

func genBTree() *BinaryTree {
	bt := &BinaryTree{}
	root := &node{
		val: "0",
	}
	node1 := &node{
		val: "1",
	}
	node2 := &node{
		val: "2",
	}
	node3 := &node{
		val: "3",
	}
	node4 := &node{
		val: "4",
	}
	node5 := &node{
		val: "5",
	}
	node6 := &node{
		val: "6",
	}
	node7 := &node{
		val: "7",
	}
	node8 := &node{
		val: "8",
	}
	node9 := &node{
		val: "9",
	}
	node10 := &node{
		val: "10",
	}
	node11 := &node{
		val: "11",
	}
	root.left, root.right = node1, node2
	node1.left, node1.right = node3, node4
	node2.left, node2.right = node5, node6
	node3.left, node3.right = node7, node8
	node4.left, node4.right = node9, node10
	node5.left, node5.right = node11, nil

	bt.root = root
	return bt
}

func TestBinaryTree_PreOrderTraverse(t *testing.T) {
	bt := genBTree()
	bt.PreOrderTraverse(func(s string) {
		fmt.Printf("%v ", s)
	})
}

func TestBinaryTree_InOrderTraverse(t *testing.T) {
	bt := genBTree()
	bt.InOrderTraverse(func(s string) {
		fmt.Printf("%v ", s)
	})
}

func TestBinaryTree_PostOrderTraverse(t *testing.T) {
	bt := genBTree()
	bt.PostOrderTraverse(func(s string) {
		fmt.Printf("%v ", s)
	})
}

func TestBinaryTree_PreOrderIter(t *testing.T) {
	bt := genBTree()
	var res []string
	bt.PreOrderTraverse(func(s string) {
		fmt.Printf("%v ", s)
		res = append(res, s)
	})
	fmt.Println()
	var res1 []string
	bt.PreOrderIter(func(s string) {
		fmt.Printf("%v ", s)
		res1 = append(res1, s)
	})
	if !slices.Equal(res, res1) {
		t.Fatal("preorder traverse is not equal iter")
	}
}

func TestBinaryTree_InOrderIter(t *testing.T) {
	bt := genBTree()
	var res []string
	bt.InOrderTraverse(func(s string) {
		fmt.Printf("%v ", s)
		res = append(res, s)
	})
	fmt.Println()
	var res1 []string
	bt.InOrderIter(func(s string) {
		fmt.Printf("%v ", s)
		res1 = append(res1, s)
	})
	if !slices.Equal(res, res1) {
		t.Fatal("inorder traverse is not equal iter")
	}
}

func TestBinaryTree_PostOrderIter(t *testing.T) {
	bt := genBTree()
	var res []string
	bt.PostOrderTraverse(func(s string) {
		fmt.Printf("%v ", s)
		res = append(res, s)
	})
	fmt.Println()
	var res1 []string
	bt.PostOrderIter(func(s string) {
		fmt.Printf("%v ", s)
		res1 = append(res1, s)
	})
	if !slices.Equal(res, res1) {
		t.Fatal("postorder traverse is not equal iter")
	}
}
