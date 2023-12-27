package binary_tree

import "github.com/shao1f/algorithm/stack"

type node struct {
	val   string
	left  *node
	right *node
}

type BinaryTree struct {
	root *node
}

func (bt *BinaryTree) PreOrderTraverse(f func(string)) {
	preOrder(bt.root, f)
}

func preOrder(root *node, f func(string)) {
	if root != nil {
		f(root.val)
		preOrder(root.left, f)
		preOrder(root.right, f)
	}
}

// PreOrderIter 前序遍历 迭代
func (bt *BinaryTree) PreOrderIter(f func(string2 string)) {
	if bt.root == nil {
		return
	}
	// 模拟栈
	st := stack.NewStack[*node](10)
	// 先将根节点入栈
	st.Push(bt.root)
	for !st.Empty() {
		n := st.Pop()
		f(n.val)
		// 先入栈右子树，因为先入栈的后出栈
		if n.right != nil {
			st.Push(n.right)
		}
		if n.left != nil {
			st.Push(n.left)
		}
	}
}

func (bt *BinaryTree) InOrderTraverse(f func(string)) {
	inOrder(bt.root, f)
}

func inOrder(root *node, f func(string)) {
	if root != nil {
		inOrder(root.left, f)
		f(root.val)
		inOrder(root.right, f)
	}
}

// InOrderIter 中序遍历 迭代
func (bt *BinaryTree) InOrderIter(f func(string)) {
	if bt.root == nil {
		return
	}
	st := stack.NewStack[*node](10)
	r := bt.root
	for r != nil || !st.Empty() {
		for r != nil {
			st.Push(r)
			r = r.left
		}
		top := st.Pop()
		f(top.val)
		r = top.right
	}
}

func (bt *BinaryTree) PostOrderTraverse(f func(string)) {
	postOrder(bt.root, f)
}

func postOrder(root *node, f func(string)) {
	if root != nil {
		postOrder(root.left, f)
		postOrder(root.right, f)
		f(root.val)
	}
}

// PostOrderIter 后序遍历，迭代
func (bt *BinaryTree) PostOrderIter(f func(string)) {
	if bt.root == nil {
		return
	}
	st := stack.NewStack[*node](10)
	var output []string
	st.Push(bt.root)
	for !st.Empty() {
		top := st.Pop()
		output = append(output, top.val)
		if top.left != nil {
			st.Push(top.left)
		}
		if top.right != nil {
			st.Push(top.right)
		}
	}
	for i := len(output) - 1; i >= 0; i-- {
		f(output[i])
	}

}

// LevelOrder 层序
func (bt *BinaryTree) LevelOrder(f func(string)) {
	if bt.root == nil {
		return
	}
	st := stack.NewStack[*node](10)
	var res []string
	for !st.Empty() {
		top := st.Pop()
		res = append(res, top.val)
	}
}
