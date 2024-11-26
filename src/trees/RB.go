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

func (tree *RBTree) Print() {
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
