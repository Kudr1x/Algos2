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

func (tree *RBTree) Insert(key int) {
	newNode := NewNode(key)
	if tree.root == nil {
		tree.root = newNode
		tree.root.color = BLACK
		return
	}

	tree.insertNode(tree.root, newNode)
	tree.fixInsert(newNode)
}

func (tree *RBTree) insertNode(root, newNode *rbtNode) {
	if newNode.key < root.key {
		if root.left == nil {
			root.left = newNode
			newNode.parent = root
		} else {
			tree.insertNode(root.left, newNode)
		}
	} else {
		if root.right == nil {
			root.right = newNode
			newNode.parent = root
		} else {
			tree.insertNode(root.right, newNode)
		}
	}
}

func (tree *RBTree) fixInsert(newNode *rbtNode) {
	for newNode != tree.root && newNode.parent.color == RED {
		if newNode.parent == newNode.parent.parent.left {
			uncle := newNode.parent.parent.right
			if uncle != nil && uncle.color == RED {
				newNode.parent.color = BLACK
				uncle.color = BLACK
				newNode.parent.parent.color = RED
				newNode = newNode.parent.parent
			} else {
				if newNode == newNode.parent.right {
					newNode = newNode.parent
					tree.rotateLeft(newNode)
				}
				newNode.parent.color = BLACK
				newNode.parent.parent.color = RED
				tree.rotateRight(newNode.parent.parent)
			}
		} else {
			uncle := newNode.parent.parent.left
			if uncle != nil && uncle.color == RED {
				newNode.parent.color = BLACK
				uncle.color = BLACK
				newNode.parent.parent.color = RED
				newNode = newNode.parent.parent
			} else {
				if newNode == newNode.parent.left {
					newNode = newNode.parent
					tree.rotateRight(newNode)
				}
				newNode.parent.color = BLACK
				newNode.parent.parent.color = RED
				tree.rotateLeft(newNode.parent.parent)
			}
		}
	}
	tree.root.color = BLACK
}

func (tree *RBTree) rotateLeft(x *rbtNode) {
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

func (tree *RBTree) rotateRight(x *rbtNode) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		tree.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y
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

func (tree *RBTree) Delete(key int) {
	nodeToDelete := tree.Search(key)
	if nodeToDelete == nil {
		fmt.Println("Node not found")
		return
	}
	tree.deleteNode(nodeToDelete)
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

func (tree *RBTree) fixDelete(x *rbtNode) {
	for x != tree.root && (x == nil || x.color == BLACK) {
		if x == x.parent.left {
			w := x.parent.right
			if w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				tree.rotateLeft(x.parent)
				w = x.parent.right
			}
			if (w.left == nil || w.left.color == BLACK) && (w.right == nil || w.right.color == BLACK) {
				w.color = RED
				x = x.parent
			} else {
				if w.right == nil || w.right.color == BLACK {
					w.left.color = BLACK
					w.color = RED
					tree.rotateRight(w)
					w = x.parent.right
				}
				w.color = x.parent.color
				x.parent.color = BLACK
				if w.right != nil {
					w.right.color = BLACK
				}
				tree.rotateLeft(x.parent)
				x = tree.root
			}
		} else {
			w := x.parent.left
			if w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				tree.rotateRight(x.parent)
				w = x.parent.left
			}
			if (w.left == nil || w.left.color == BLACK) && (w.right == nil || w.right.color == BLACK) {
				w.color = RED
				x = x.parent
			} else {
				if w.left == nil || w.left.color == BLACK {
					w.right.color = BLACK
					w.color = RED
					tree.rotateLeft(w)
					w = x.parent.left
				}
				w.color = x.parent.color
				x.parent.color = BLACK
				if w.left != nil {
					w.left.color = BLACK
				}
				tree.rotateRight(x.parent)
				x = tree.root
			}
		}
	}
	if x != nil {
		x.color = BLACK
	}
}
