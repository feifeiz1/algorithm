package binary_tree

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
	var st []*node
	// 先将根节点入栈
	st = append(st, bt.root)
	for len(st) != 0 {
		n := st[len(st)-1]
		st = st[:len(st)-1]
		f(n.val)
		// 先入栈右子树，因为先入栈的后出栈
		if n.right != nil {
			st = append(st, n.right)
		}
		if n.left != nil {
			st = append(st, n.left)
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
	var st []*node
	r := bt.root
	for r != nil || len(st) != 0 {
		for r != nil {
			st = append(st, r)
			r = r.left
		}
		top := st[len(st)-1]
		st = st[:len(st)-1]
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
	var st []*node
	var output []string
	st = append(st, bt.root)
	for len(st) != 0 {
		top := st[len(st)-1]
		st = st[:len(st)-1]
		output = append(output, top.val)
		if top.left != nil {
			st = append(st, top.left)
		}
		if top.right != nil {
			st = append(st, top.right)
		}
	}
	for i := len(output) - 1; i >= 0; i-- {
		f(output[i])
	}

}
