package main

import (
	"fmt"
	"slices"
	"sync"
)

/*
	二叉搜索树
		1. 二叉搜索树的定义
			二叉搜索树是一种特殊的二叉树，它的每个节点的值都大于其左子树的任意节点的值，小于其右子树的任意节点的值。
*/

type node struct {
	key   int
	val   string
	left  *node
	right *node
}

type BinarySearchTree struct {
	root *node
	lock sync.RWMutex
}

func (bst *BinarySearchTree) Insert(k int, v string) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	n := &node{
		key:   k,
		val:   v,
		left:  nil,
		right: nil,
	}
	if bst.root == nil {
		bst.root = n
	} else {
		insertNode(bst.root, n)
	}
}

func insertNode(node, newNode *node) {
	if newNode.key < node.key {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

func (bst *BinarySearchTree) InorderTraverse(f func(string)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	inorderTraverse(bst.root, f)
}

func inorderTraverse(node *node, f func(string)) {
	if node != nil {
		inorderTraverse(node.left, f)
		f(node.val)
		inorderTraverse(node.right, f)
	}
}

func (bst *BinarySearchTree) PreOrderTraverse(f func(string)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	preOrderTraverse(bst.root, f)
}

func preOrderTraverse(node *node, f func(string)) {
	if node != nil {
		f(node.val)
		preOrderTraverse(node.left, f)
		preOrderTraverse(node.right, f)
	}
}

func (bst *BinarySearchTree) PostOrderTraverse(f func(string)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	postOrderTraverse(bst.root, f)
}

func postOrderTraverse(node *node, f func(string)) {
	if node != nil {
		postOrderTraverse(node.left, f)
		postOrderTraverse(node.right, f)
		f(node.val)
	}
}

func (bst *BinarySearchTree) Min() (val string) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	n := bst.root
	if n == nil {
		return
	}
	for {
		if n.left == nil {
			val = n.val
			break
		}
		n = n.left
	}
	return
}

func (bst *BinarySearchTree) Max() (val string) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	n := bst.root
	if n == nil {
		return
	}
	for {
		if n.right == nil {
			val = n.val
			break
		}
		n = n.right
	}
	return
}

func (bst *BinarySearchTree) Search(key int) bool {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	return search(bst.root, key)
}

func search(node *node, key int) bool {
	if node == nil {
		return false
	}
	if key < node.key {
		return search(node.left, key)
	}
	if key > node.key {
		return search(node.right, key)
	}
	return true
}

func (bst *BinarySearchTree) Remove(key int) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	remove(bst.root, key)
}

func remove(node *node, key int) *node {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.left = remove(node.left, key)
		return node
	}
	if key > node.key {
		node.right = remove(node.right, key)
		return node
	}
	// 找到了要删除的节点

	// 没有左右子树
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}
	// 只有右子树
	if node.left == nil {
		node = node.right
		return node
	}
	// 只有左子树
	if node.right == nil {
		node = node.left
		return node
	}
	// 左右子树都有，找到右子树最小的，和当前节点交换
	rightLeftSmallNode := node.right
	for {
		if rightLeftSmallNode != nil && rightLeftSmallNode.left != nil {
			rightLeftSmallNode = rightLeftSmallNode.left
		} else {
			break
		}
	}
	node.key, node.val = rightLeftSmallNode.key, rightLeftSmallNode.val
	node.right = remove(node.right, node.key)
	return node
}

func (bst *BinarySearchTree) String() {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	fmt.Println("------------------------------------------------")
	stringify(bst.root, 0)
	fmt.Println("------------------------------------------------")
}

// internal recursive function to print a tree
func stringify(n *node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.right, level)
		fmt.Printf(format+"%d\n", n.key)
		stringify(n.left, level)

	}
}
func fillBst(bst *BinarySearchTree) {
	bst.Insert(8, "8")
	bst.Insert(4, "4")
	bst.Insert(10, "10")
	bst.Insert(2, "2")
	bst.Insert(6, "6")
	bst.Insert(1, "1")
	bst.Insert(3, "3")
	bst.Insert(5, "5")
	bst.Insert(7, "7")
	bst.Insert(9, "9")
}

func main() {
	b := &BinarySearchTree{}
	fillBst(b)
	b.String()

	b.Insert(11, "11")
	b.String()

	var preorderResult []string
	b.PreOrderTraverse(func(s string) {
		preorderResult = append(preorderResult, s)
	})
	fmt.Println(preorderResult)
	preorderList := []string{"8", "4", "2", "1", "3", "6", "5", "7", "10", "9", "11"}
	fmt.Printf("pre res:%v\n", slices.Equal(preorderList, preorderResult))

	var inorderResult []string
	b.InorderTraverse(func(s string) {
		inorderResult = append(inorderResult, s)
	})
	fmt.Println(inorderResult)
	inorderList := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}
	fmt.Printf("in res:%v\n", slices.Equal(inorderList, inorderResult))

	var postorderResult []string
	b.PostOrderTraverse(func(s string) {
		postorderResult = append(postorderResult, s)
	})
	fmt.Println(postorderResult)
	postorderList := []string{"1", "3", "2", "5", "7", "6", "4", "9", "11", "10", "8"}
	fmt.Printf("post res:%v\n", slices.Equal(postorderList, postorderResult))

	fmt.Println(b.Min())
	fmt.Println(b.Max())

	fmt.Println(b.Search(6))
	fmt.Println(b.Search(10))
	fmt.Println(b.Search(50))

	b.Remove(7)
	b.String()

	b.Remove(4)
	b.String()

	b.Remove(8)
	b.String()
}
