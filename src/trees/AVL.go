package trees

import (
	"Algos2/src/util"
	"container/list"
	"fmt"
)

type avlNode struct {
	key    int
	height int
	left   *avlNode
	right  *avlNode
}

type AVLTree struct {
	root *avlNode
}

func (n *avlNode) GetHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *avlNode) GetBalance() int {
	if n == nil {
		return 0
	}
	return n.left.GetHeight() - n.right.GetHeight()
}

func RightRotate(y *avlNode) *avlNode {
	x := y.left
	T2 := x.right

	x.right = y
	y.left = T2

	y.height = util.Max(y.left.GetHeight(), y.right.GetHeight()) + 1
	x.height = util.Max(x.left.GetHeight(), x.right.GetHeight()) + 1

	return x
}

func LeftRotate(x *avlNode) *avlNode {
	y := x.right
	T2 := y.left

	y.left = x
	x.right = T2

	x.height = util.Max(x.left.GetHeight(), x.right.GetHeight()) + 1
	y.height = util.Max(y.left.GetHeight(), y.right.GetHeight()) + 1

	return y
}

func (tree *AVLTree) Insert(key int) {
	tree.root = tree.insert(tree.root, key)
}

func (tree *AVLTree) insert(node *avlNode, key int) *avlNode {
	if node == nil {
		return &avlNode{key: key, height: 1}
	}

	if key < node.key {
		node.left = tree.insert(node.left, key)
	} else if key > node.key {
		node.right = tree.insert(node.right, key)
	} else {
		return node
	}

	node.height = 1 + util.Max(node.left.GetHeight(), node.right.GetHeight())

	balance := node.GetBalance()

	if balance > 1 && key < node.left.key {
		return RightRotate(node)
	}

	if balance < -1 && key > node.right.key {
		return LeftRotate(node)
	}

	if balance > 1 && key > node.left.key {
		node.left = LeftRotate(node.left)
		return RightRotate(node)
	}

	if balance < -1 && key < node.right.key {
		node.right = RightRotate(node.right)
		return LeftRotate(node)
	}

	return node
}

func (tree *AVLTree) Search(key int) bool {
	return tree.search(tree.root, key)
}

func (tree *AVLTree) search(node *avlNode, key int) bool {
	if node == nil {
		return false
	}

	if key < node.key {
		return tree.search(node.left, key)
	} else if key > node.key {
		return tree.search(node.right, key)
	} else {
		return true
	}
}

func (tree *AVLTree) Delete(key int) {
	tree.root = tree.delete(tree.root, key)
}

func (tree *AVLTree) delete(node *avlNode, key int) *avlNode {
	if node == nil {
		return node
	}

	if key < node.key {
		node.left = tree.delete(node.left, key)
	} else if key > node.key {
		node.right = tree.delete(node.right, key)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}

		temp := minValueNode(node.right)
		node.key = temp.key
		node.right = tree.delete(node.right, temp.key)
	}

	node.height = 1 + util.Max(node.left.GetHeight(), node.right.GetHeight())

	balance := node.GetBalance()

	if balance > 1 && node.left.GetBalance() >= 0 {
		return RightRotate(node)
	}

	if balance > 1 && node.left.GetBalance() < 0 {
		node.left = LeftRotate(node.left)
		return RightRotate(node)
	}

	if balance < -1 && node.right.GetBalance() <= 0 {
		return LeftRotate(node)
	}

	if balance < -1 && node.right.GetBalance() > 0 {
		node.right = RightRotate(node.right)
		return LeftRotate(node)
	}

	return node
}

func minValueNode(node *avlNode) *avlNode {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current
}

func (tree *AVLTree) InOrderTraversal() {
	tree.inOrderTraversal(tree.root)
}

func (tree *AVLTree) inOrderTraversal(node *avlNode) {
	if node != nil {
		tree.inOrderTraversal(node.left)
		fmt.Printf("%d ", node.key)
		tree.inOrderTraversal(node.right)
	}
}

func (tree *AVLTree) Height() int {
	return tree.root.GetHeight()
}

func (tree *AVLTree) LevelOrderTraversal() {
	if tree.root == nil {
		return
	}

	queue := list.New()

	queue.PushBack(tree.root)

	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*avlNode)

		fmt.Printf("%d ", node.key)

		if node.left != nil {
			queue.PushBack(node.left)
		}

		if node.right != nil {
			queue.PushBack(node.right)
		}
	}
}

func (tree *AVLTree) PostLevelOrderTraversal() {
	if tree.root == nil {
		return
	}

	queue := list.New()

	stack := list.New()

	queue.PushBack(tree.root)

	for queue.Len() > 0 {

		node := queue.Remove(queue.Front()).(*avlNode)

		stack.PushBack(node)

		if node.right != nil {
			queue.PushBack(node.right)
		}

		if node.left != nil {
			queue.PushBack(node.left)
		}
	}

	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*avlNode)
		fmt.Printf("%d ", node.key)
	}
}

func (tree *AVLTree) PreOrderTraversal() {
	tree.preOrderTraversal(tree.root)
}

func (tree *AVLTree) preOrderTraversal(node *avlNode) {
	if node != nil {
		fmt.Printf("%d ", node.key)        // Обработка корня
		tree.preOrderTraversal(node.left)  // Обход левого поддерева
		tree.preOrderTraversal(node.right) // Обход правого поддерева
	}
}
