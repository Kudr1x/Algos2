package trees

import (
	"fmt"
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

func (bst *BSTree) Print() {
	inOrderTraversal(bst.root)
}

func inOrderTraversal(node *bstNode) {
	if node != nil {
		inOrderTraversal(node.left)
		fmt.Print(node.key, " ")
		inOrderTraversal(node.right)
	}
}