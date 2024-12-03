package trees

import "fmt"

const (
	RED   = true
	BLACK = false
)

type rbtNode struct {
	color  bool
	key    int
	left   *rbtNode
	right  *rbtNode
	parent *rbtNode
}

type RBTree struct {
	root *rbtNode
}

func NewNode(key int) *rbtNode {
	return &rbtNode{color: RED, key: key, left: nil, right: nil, parent: nil}
}

func (tree *RBTree) Search(key int) *rbtNode {
	return tree.searchNode(tree.root, key)
}

func (tree *RBTree) inOrderTraversal(node *rbtNode) {
	if node != nil {
		tree.inOrderTraversal(node.left)
		fmt.Printf("%d (%v) ", node.key, node.color)
		tree.inOrderTraversal(node.right)
	}
}

func (tree *RBTree) InOrderTraversal() {
	tree.inOrderTraversal(tree.root)
}

func (tree *RBTree) searchNode(node *rbtNode, key int) *rbtNode {
	if node == nil || node.key == key {
		return node
	}
	if key < node.key {
		return tree.searchNode(node.left, key)
	}
	return tree.searchNode(node.right, key)
}

func (tree *RBTree) Height() int {
	return tree.height(tree.root)
}

func (tree *RBTree) height(node *rbtNode) int {
	if node == nil {
		return 0
	}

	leftHeight := tree.height(node.left)
	rightHeight := tree.height(node.right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func (tree *RBTree) transplant(u, v *rbtNode) {
	if u.parent == nil {
		tree.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

func (tree *RBTree) minimum(node *rbtNode) *rbtNode {
	for node.left != nil {
		node = node.left
	}
	return node
}

func (tree *RBTree) LevelOrderTraversal() {
	if tree.root == nil {
		return
	}

	queue := []*rbtNode{tree.root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Printf("%d (%v) ", node.key, node.color)

		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
}

func (tree *RBTree) PreOrderTraversal() {
	tree.preOrderTraversal(tree.root)
}

func (tree *RBTree) preOrderTraversal(node *rbtNode) {
	if node != nil {
		fmt.Printf("%d (%v) ", node.key, node.color)
		tree.preOrderTraversal(node.left)
		tree.preOrderTraversal(node.right)
	}
}

func (tree *RBTree) PostOrderTraversal() {
	tree.postOrderTraversal(tree.root)
}

func (tree *RBTree) postOrderTraversal(node *rbtNode) {
	if node != nil {
		tree.postOrderTraversal(node.left)
		tree.postOrderTraversal(node.right)
		fmt.Printf("%d (%v) ", node.key, node.color)
	}
}

func (tree *RBTree) Insert(key int) {
	newNode := NewNode(key)
	tree.root = tree.insertNode(tree.root, newNode)
	tree.fixInsert(newNode)
}

func (tree *RBTree) insertNode(root, newNode *rbtNode) *rbtNode {
	if root == nil {
		return newNode
	}

	if newNode.key < root.key {
		root.left = tree.insertNode(root.left, newNode)
		root.left.parent = root
	} else {
		root.right = tree.insertNode(root.right, newNode)
		root.right.parent = root
	}

	return root
}

func (tree *RBTree) fixInsert(node *rbtNode) {
	for node.parent != nil && node.parent.color == RED {
		if node.parent == node.parent.parent.left {
			uncle := node.parent.parent.right
			if uncle != nil && uncle.color == RED {
				// Case 1: Uncle is red
				node.parent.color = BLACK
				uncle.color = BLACK
				node.parent.parent.color = RED
				node = node.parent.parent
			} else {
				if node == node.parent.right {
					// Case 2: Node is right child
					node = node.parent
					tree.leftRotate(node)
				}
				// Case 3: Node is left child
				node.parent.color = BLACK
				node.parent.parent.color = RED
				tree.rightRotate(node.parent.parent)
			}
		} else {
			uncle := node.parent.parent.left
			if uncle != nil && uncle.color == RED {
				// Case 1: Uncle is red
				node.parent.color = BLACK
				uncle.color = BLACK
				node.parent.parent.color = RED
				node = node.parent.parent
			} else {
				if node == node.parent.left {
					// Case 2: Node is left child
					node = node.parent
					tree.rightRotate(node)
				}
				// Case 3: Node is right child
				node.parent.color = BLACK
				node.parent.parent.color = RED
				tree.leftRotate(node.parent.parent)
			}
		}
	}
	tree.root.color = BLACK
}

func (tree *RBTree) leftRotate(x *rbtNode) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		tree.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (tree *RBTree) rightRotate(y *rbtNode) {
	x := y.left
	y.left = x.right
	if x.right != nil {
		x.right.parent = y
	}
	x.parent = y.parent
	if y.parent == nil {
		tree.root = x
	} else if y == y.parent.right {
		y.parent.right = x
	} else {
		y.parent.left = x
	}
	x.right = y
	y.parent = x
}

func (tree *RBTree) Delete(key int) {
	node := tree.Search(key)
	if node != nil {
		tree.deleteNode(node)
	}
}

func (tree *RBTree) deleteNode(z *rbtNode) {
	y := z
	yOriginalColor := y.color
	var x *rbtNode

	if z.left == nil {
		x = z.right
		tree.transplant(z, z.right)
	} else if z.right == nil {
		x = z.left
		tree.transplant(z, z.left)
	} else {
		y = tree.minimum(z.right)
		yOriginalColor = y.color
		x = y.right
		if y.parent == z {
			x.parent = y
		} else {
			tree.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		tree.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}

	if yOriginalColor == BLACK {
		tree.fixDelete(x)
	}
}

func (tree *RBTree) fixDelete(x *rbtNode) {
	for x != tree.root && (x == nil || x.color == BLACK) {
		if x == nil {
			break
		}
		if x == x.parent.left {
			w := x.parent.right
			if w != nil && w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				tree.leftRotate(x.parent)
				w = x.parent.right
			}
			if w == nil || (w.left == nil || w.left.color == BLACK) && (w.right == nil || w.right.color == BLACK) {
				if w != nil {
					w.color = RED
				}
				x = x.parent
			} else {
				if w.right == nil || w.right.color == BLACK {
					if w.left != nil {
						w.left.color = BLACK
					}
					tree.rightRotate(w)
					w = x.parent.right
				}
				if w != nil {
					w.color = x.parent.color
				}
				x.parent.color = BLACK
				if w != nil && w.right != nil {
					w.right.color = BLACK
				}
				tree.leftRotate(x.parent)
				x = tree.root
			}
		} else {
			w := x.parent.left
			if w != nil && w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				tree.rightRotate(x.parent)
				w = x.parent.left
			}
			if w == nil || (w.right == nil || w.right.color == BLACK) && (w.left == nil || w.left.color == BLACK) {
				if w != nil {
					w.color = RED
				}
				x = x.parent
			} else {
				if w.left == nil || w.left.color == BLACK {
					tree.leftRotate(w)
					w = x.parent.left
				}
				if w != nil {
					w.color = x.parent.color
				}
				x.parent.color = BLACK
				if w != nil && w.left != nil {
					w.left.color = BLACK
				}
				tree.rightRotate(x.parent)
				x = tree.root
			}
		}
	}
	if x != nil {
		x.color = BLACK
	}
}
