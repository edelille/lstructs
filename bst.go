package lstructs

type BSTNode[T comparable] struct {
	Val   T
	Left  *BSTNode[T]
	Right *BSTNode[T]
}

type BST[T comparable] struct {
	Length int
	Head   *BSTNode[T]
	Less   LessFn[T]
}

func NewBST[T comparable](less LessFn[T]) *BST[T] {
	return &BST[T]{0, nil, less}
}

func (b *BST[T]) Insert(x T) {
	newNode := &BSTNode[T]{x, nil, nil}
	if b.Head == nil {
		b.Head = newNode
	} else {
		bstInsert(b.Head, newNode, b.Less)
	}
	b.Length++
}

func (b *BST[T]) Search(x T) bool {
	return bstSearch(b.Head, x, b.Less)
}

func (b *BST[T]) Remove(x T) {
	if bstSearch(b.Head, x, b.Less) {
		bstRemove(b.Head, x, b.Less)
		b.Length--
	}
}

func (b *BST[T]) Rebalance() {
	inorder := bstInorderTraversal(b.Head)
	b.Head = bstRebalance(inorder)
}

func (b *BST[T]) Size() int {
	return b.Length
}

func (b *BST[T]) InorderTraversal() []T {
	return bstInorderTraversal(b.Head)
}

func (b *BST[T]) PreorderTraversal() []T {
	return bstPreorderTraversal(b.Head)
}

func (b *BST[T]) PostorderTraversal() []T {
	return bstPostorderTraversal(b.Head)
}

func bstInsert[T comparable](b, c *BSTNode[T], less LessFn[T]) {
	if less(c.Val, b.Val) {
		if b.Left == nil {
			b.Left = c
		} else {
			bstInsert(b.Left, c, less)
		}
	} else {
		if b.Right == nil {
			b.Right = c
		} else {
			bstInsert(b.Right, c, less)
		}
	}
}

func bstSearch[T comparable](root *BSTNode[T], x T, less LessFn[T]) bool {
	if root == nil {
		return false
	}

	if root.Val == x {
		return true
	}

	if less(x, root.Val) {
		return bstSearch(root.Left, x, less)
	} else {
		return bstSearch(root.Right, x, less)
	}
}

// The boolean determines if an item was successfully deleted or not
func bstRemove[T comparable](root *BSTNode[T], x T, less LessFn[T]) *BSTNode[T] {
	// base case
	if root == nil {
		return root
	}

	// Preserves the root, and recurses deeper into the tree to find and delete
	if less(x, root.Val) {
		root.Left = bstRemove(root.Left, x, less)
		return root
	} else if less(root.Val, x) {
		root.Right = bstRemove(root.Right, x, less)
		return root
	}

	// if we reach this point then, root.val == x, delete it
	if root.Left == nil {
		tmp := root.Right
		root = nil
		return tmp
	} else if root.Right == nil {
		tmp := root.Left
		root = nil
		return tmp
	} else { // both childs exist
		parent := root
		succ := root.Right
		for succ.Left != nil {
			parent = succ
			succ = succ.Left
		}

		if parent != root {
			parent.Left = succ.Right
		} else {
			parent.Right = succ.Right
		}

		root.Val = succ.Val
		succ = nil
		return root
	}
}

func bstRebalance[T comparable](arr []T) *BSTNode[T] {
	if len(arr) == 0 {
		return nil
	}

	if len(arr) == 1 {
		return &BSTNode[T]{arr[0], nil, nil}
	} else if len(arr) == 2 {
		return &BSTNode[T]{arr[1], &BSTNode[T]{arr[0], nil, nil}, nil}
	}

	mi := len(arr)/2 + 1

	return &BSTNode[T]{
		arr[mi],
		bstRebalance(arr[:mi]),
		bstRebalance(arr[mi+1:]),
	}
}

func bstInorderTraversal[T comparable](x *BSTNode[T]) []T {
	if x == nil {
		return []T{}
	}

	return append(append(bstInorderTraversal[T](x.Left), x.Val), bstInorderTraversal[T](x.Right)...)
}

func bstPreorderTraversal[T comparable](x *BSTNode[T]) []T {
	if x == nil {
		return []T{}
	}

	return append(append([]T{x.Val}, bstPreorderTraversal[T](x.Left)...), bstPreorderTraversal[T](x.Right)...)
}

func bstPostorderTraversal[T comparable](x *BSTNode[T]) []T {
	if x == nil {
		return []T{}
	}

	return append(append(bstPostorderTraversal[T](x.Left), bstPostorderTraversal[T](x.Right)...), x.Val)
}
