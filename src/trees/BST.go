package trees

import (
	"fmt"
)

type bstNode struct {
	Value int
	Left  *bstNode
	Right *bstNode
}

type BST struct {
	Root *bstNode
}

func (bst *BST) Insert(value int) {
	bst.Root = insertNode(bst.Root, value)
}

func insertNode(thisNode *bstNode, value int) *bstNode {
	if thisNode == nil {
		return &bstNode{Value: value}
	}
	if value < thisNode.Value {
		thisNode.Left = insertNode(thisNode.Left, value)
	} else if value > thisNode.Value {
		thisNode.Right = insertNode(thisNode.Right, value)
	}
	return thisNode
}

func (bst *BST) Search(value int) bool {
	return searchNode(bst.Root, value)
}

func searchNode(node *bstNode, value int) bool {
	if node == nil {
		return false
	}
	if value == node.Value {
		return true
	} else if value < node.Value {
		return searchNode(node.Left, value)
	} else {
		return searchNode(node.Right, value)
	}
}

func (bst *BST) Delete(value int) {
	bst.Root = deleteNode(bst.Root, value)
}

func deleteNode(node *bstNode, value int) *bstNode {
	if node == nil {
		return nil
	}
	if value < node.Value {
		node.Left = deleteNode(node.Left, value)
	} else if value > node.Value {
		node.Right = deleteNode(node.Right, value)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		minNode := findMin(node.Right)
		node.Value = minNode.Value
		node.Right = deleteNode(node.Right, minNode.Value)
	}
	return node
}

func findMin(node *bstNode) *bstNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func (bst *BST) InOrderTraversal() {
	inOrderTraversal(bst.Root)
}

func inOrderTraversal(node *bstNode) {
	if node != nil {
		inOrderTraversal(node.Left)
		fmt.Print(node.Value, " ")
		inOrderTraversal(node.Right)
	}
}
