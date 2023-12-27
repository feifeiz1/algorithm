package binary_search_tree

import (
	"cmp"
	"fmt"
	"sync"
)

type node[K cmp.Ordered, V fmt.Stringer] struct {
	key   K
	val   V
	left  *node[K, V]
	right *node[K, V]
}

type BinarySearchTree[K cmp.Ordered, V fmt.Stringer] struct {
	sync.RWMutex
	root *node[K, V]
}

func (bst *BinarySearchTree[K, V]) PrintTree() {
	bst.RLock()
	defer bst.RUnlock()
	fmt.Println("=================================================")
	stringify(bst.root, 0)
	fmt.Println("=================================================")

}

func stringify[K cmp.Ordered, V fmt.Stringer](n *node[K, V], level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---["
		level++
		stringify(n.right, level)
		fmt.Printf(format+"%v\n", n.val.String())
		stringify(n.left, level)
	}
}

func (bst *BinarySearchTree[K, V]) Insert(k K, v V) {
	bst.Lock()
	defer bst.Unlock()
	n := &node[K, V]{
		key: k,
		val: v,
	}
	if bst.root == nil {
		bst.root = n
	} else {
		insert(bst.root, n)
	}
}

func insert[K cmp.Ordered, V fmt.Stringer](root *node[K, V], n *node[K, V]) {
	if n.key < root.key {
		if root.left == nil {
			root.left = n
		} else {
			insert(root.left, n)
		}
	} else if n.key > root.key {
		if root.right == nil {
			root.right = n
		} else {
			insert(root.right, n)
		}
	}
}

func (bst *BinarySearchTree[K, V]) PreOrderTraverse(f func(V)) {
	bst.RLock()
	defer bst.RUnlock()

	preOrderTraverse(bst.root, f)
}

func preOrderTraverse[K cmp.Ordered, V fmt.Stringer](root *node[K, V], f func(V)) {
	if root != nil {
		f(root.val)
		preOrderTraverse(root.left, f)
		preOrderTraverse(root.right, f)
	}
}

func (bst *BinarySearchTree[K, V]) InOrderTraverse(f func(V)) {
	bst.RLock()
	defer bst.RUnlock()
	inOrderTraverse(bst.root, f)
}

func inOrderTraverse[K cmp.Ordered, V fmt.Stringer](root *node[K, V], f func(V)) {
	if root != nil {
		inOrderTraverse(root.left, f)
		f(root.val)
		inOrderTraverse(root.right, f)
	}
}

func (bst *BinarySearchTree[K, V]) PostOrderTraverse(f func(V)) {
	bst.RLock()
	defer bst.RUnlock()
	postOrderTraverse(bst.root, f)
}

func postOrderTraverse[K cmp.Ordered, V fmt.Stringer](root *node[K, V], f func(V)) {
	if root != nil {
		postOrderTraverse(root.left, f)
		postOrderTraverse(root.right, f)
		f(root.val)
	}
}

func (bst *BinarySearchTree[K, V]) Min() (val V) {
	bst.RLock()
	defer bst.RUnlock()
	if bst.root == nil {
		return
	}
	cur := bst.root
	for cur.left != nil {
		cur = cur.left
	}
	return cur.val
}

func (bst *BinarySearchTree[K, V]) Max() (val V) {
	bst.RLock()
	defer bst.RUnlock()
	if bst.root == nil {
		return
	}
	cur := bst.root
	for cur.right != nil {
		cur = cur.right
	}
	return cur.val
}

func (bst *BinarySearchTree[K, V]) Search(k K) bool {
	bst.RLock()
	defer bst.RUnlock()
	return search(bst.root, k)
}

func search[K cmp.Ordered, V fmt.Stringer](root *node[K, V], k K) bool {
	if root == nil {
		return false
	}
	if k < root.key {
		return search(root.left, k)
	}
	if k > root.key {
		return search(root.right, k)
	}
	return true
}

func (bst *BinarySearchTree[K, V]) Remove(k K) {
	bst.Lock()
	defer bst.Unlock()
	remove(bst.root, k)
}

func remove[K cmp.Ordered, V fmt.Stringer](root *node[K, V], k K) *node[K, V] {
	if root == nil {
		return nil
	}

	if k < root.key {
		root.left = remove(root.left, k)
		return root
	}
	if k > root.key {
		root.right = remove(root.right, k)
		return root
	}
	// 找到了要删除的节点
	if root.left == nil {
		root = root.right
		return root
	}
	if root.right == nil {
		root = root.left
		return root
	}
	rightSmallNode := root.right
	for rightSmallNode.left != nil {
		rightSmallNode = rightSmallNode.left
	}
	root.key, root.val = rightSmallNode.key, rightSmallNode.val
	root.right = remove(root.right, root.key)
	return root
}
