package test

import (
	"Algos2/src/trees"
	"fmt"
)

func Avl() {
	fmt.Println("AVL - Дерево")
	tree := &trees.AVLTree{}

	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(40)
	tree.Insert(50)
	tree.Insert(25)

	fmt.Println("Симметричный обход")
	tree.InOrderTraversal()
	fmt.Println("\nОбход в высоту")
	tree.LevelOrderTraversal()
	fmt.Println("\nОбратный обход")
	tree.PostLevelOrderTraversal()
	fmt.Println("\nПрямой обход")
	tree.PreOrderTraversal()
	fmt.Println("\nУдаление элемента и симметричный обход")
	tree.Delete(10)
	tree.InOrderTraversal()
}

func Bs() {
	fmt.Println("\n\nBs - Дерево")
	tree := &trees.BSTree{}

	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(40)
	tree.Insert(50)
	tree.Insert(25)

	fmt.Println("Симметричный обход")
	tree.InOrderTraversal()
	fmt.Println("\nОбход в высоту")
	tree.LevelOrderTraversal()
	fmt.Println("\nОбратный обход")
	tree.PostOrderTraversal()
	fmt.Println("\nПрямой обход")
	tree.PreOrderTraversal()
	fmt.Println("\nУдаление элемента и симметричный обход")
	tree.Delete(10)
	tree.InOrderTraversal()
}

func Rb() {
	fmt.Println("\n\nRb - Дерево")
	tree := &trees.RBTree{}

	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(40)
	tree.Insert(50)
	tree.Insert(25)

	fmt.Println("Симметричный обход")
	tree.InOrderTraversal()
	fmt.Println("\nОбход в высоту")
	tree.LevelOrderTraversal()
	fmt.Println("\nОбратный обход")
	tree.PostOrderTraversal()
	fmt.Println("\nПрямой обход")
	tree.PreOrderTraversal()
	fmt.Println("\nУдаление элемента и симметричный обход")
	tree.Delete(10)
	tree.InOrderTraversal()
}
