package trees

import (
	"fmt"
	"sort"
)

type bstNode struct {
	key   int
	left  *bstNode
	right *bstNode
}

type BSTree struct {
	root *bstNode
}

func (bst *BSTree) Insert(key int) {
	bst.root = insertNode(bst.root, key)
}

func insertNode(thisNode *bstNode, key int) *bstNode {
	if thisNode == nil {
		return &bstNode{key: key}
	}
	if key < thisNode.key {
		thisNode.left = insertNode(thisNode.left, key)
	} else if key > thisNode.key {
		thisNode.right = insertNode(thisNode.right, key)
	}
	return thisNode
}

func (bst *BSTree) Search(key int) bool {
	return searchNode(bst.root, key)
}

func searchNode(node *bstNode, key int) bool {
	if node == nil {
		return false
	}
	if key == node.key {
		return true
	} else if key < node.key {
		return searchNode(node.left, key)
	} else {
		return searchNode(node.right, key)
	}
}

func (bst *BSTree) Delete(key int) {
	bst.root = deleteNode(bst.root, key)
}

func deleteNode(node *bstNode, key int) *bstNode {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.left = deleteNode(node.left, key)
	} else if key > node.key {
		node.right = deleteNode(node.right, key)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}

		minNode := findMin(node.right)
		node.key = minNode.key
		node.right = deleteNode(node.right, minNode.key)
	}
	return node
}

func findMin(node *bstNode) *bstNode {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current
}

func (bst *BSTree) Height() int {
	return height(bst.root)
}

func height(node *bstNode) int {
	if node == nil {
		return 0
	}
	leftHeight := height(node.left)
	rightHeight := height(node.right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func (bst *BSTree) InOrderTraversal() {
	inOrderTraversal(bst.root)
}

func inOrderTraversal(node *bstNode) {
	if node != nil {
		inOrderTraversal(node.left)
		fmt.Print(node.key, " ")
		inOrderTraversal(node.right)
	}
}

func (bst *BSTree) LevelOrderTraversal() {
	if bst.root == nil {
		return
	}

	queue := []*bstNode{bst.root}

	for len(queue) > 0 {
		currentNode := queue[0]

		queue = queue[1:]

		fmt.Print(currentNode.key, " ")

		if currentNode.left != nil {
			queue = append(queue, currentNode.left)
		}

		if currentNode.right != nil {
			queue = append(queue, currentNode.right)
		}
	}
}

func (bst *BSTree) PreOrderTraversal() {
	preOrderTraversal(bst.root)
}

func preOrderTraversal(node *bstNode) {
	if node != nil {
		fmt.Print(node.key, " ")
		preOrderTraversal(node.left)
		preOrderTraversal(node.right)
	}
}

func (bst *BSTree) PostOrderTraversal() {
	postOrderTraversal(bst.root)
}

func postOrderTraversal(node *bstNode) {
	if node != nil {
		postOrderTraversal(node.left)
		postOrderTraversal(node.right)
		fmt.Print(node.key, " ")
	}
}

func (bst *BSTree) FillBalanced(keys []int) {
	sort.Ints(keys)
	bst.root = fillBalanced(keys, 0, len(keys)-1)
}

func fillBalanced(keys []int, start, end int) *bstNode {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	node := &bstNode{key: keys[mid]}

	node.left = fillBalanced(keys, start, mid-1)
	node.right = fillBalanced(keys, mid+1, end)

	return node
}
