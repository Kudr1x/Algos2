package plot

import (
	"Algos2/src/trees"
	"fmt"
	"math/rand"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

func DrawPlots() {
	drawAvlTree()
	drawRbTree()
	drawBsTree()
}

func drawAvlTree() {
	rand.Seed(51)
	var data string

	for n := 10; n <= 1000000; n *= 10 {
		var arrXY []string
		tree := trees.AVLTree{}
		keys := rand.Perm(n)

		sort.Ints(keys)
		for _, key := range keys {
			tree.Insert(key)
		}

		arrXY = append(arrXY, strconv.Itoa(n))
		arrXY = append(arrXY, strconv.Itoa(tree.Height()))

		dataXY := strings.Join(arrXY, ",")
		data += dataXY + ";"
	}

	data = data[:len(data)-1]

	cmd := exec.Command("python", "/home/kudrix/GolandProjects/Algos2/src/plot/draw.py", data, "AVL Дерево")

	out, err := cmd.CombinedOutput()
	if err != nil {
		text := string(out)
		fmt.Println(text)
	}
}

func drawBsTree() {
	rand.Seed(51)
	var data string

	for n := 10; n <= 1000000; n *= 10 {
		var arrXY []string
		tree := trees.BSTree{}
		keys := rand.Perm(n)

		for _, key := range keys {
			tree.Insert(key)
		}

		arrXY = append(arrXY, strconv.Itoa(n))
		arrXY = append(arrXY, strconv.Itoa(tree.Height()))

		dataXY := strings.Join(arrXY, ",")
		data += dataXY + ";"
	}

	data = data[:len(data)-1]

	cmd := exec.Command("python", "/home/kudrix/GolandProjects/Algos2/src/plot/draw.py", data, "BS Дерево")

	out, err := cmd.CombinedOutput()
	if err != nil {
		text := string(out)
		fmt.Println(text)
	}
}

func drawRbTree() {
	rand.Seed(51)
	var data string

	for n := 10; n <= 1000000; n *= 10 {
		var arrXY []string
		tree := trees.RBTree{}
		keys := rand.Perm(n)

		sort.Ints(keys)
		for _, key := range keys {
			tree.Insert(key)
		}

		arrXY = append(arrXY, strconv.Itoa(n))
		arrXY = append(arrXY, strconv.Itoa(tree.Height()))

		dataXY := strings.Join(arrXY, ",")
		data += dataXY + ";"
	}

	data = data[:len(data)-1]

	cmd := exec.Command("python", "/home/kudrix/GolandProjects/Algos2/src/plot/draw.py", data, "RB Дерево")

	out, err := cmd.CombinedOutput()
	if err != nil {
		text := string(out)
		fmt.Println(text)
	}
}
