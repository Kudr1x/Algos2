package practice

import (
	"Algos2/src/trees"
	"Algos2/src/util"
	"fmt"
	"math/rand"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

func DrawPractice() {
	util.Clear("practiceData")
	drawAvlTree()
	drawRbTree()
	drawBsTree()
	drawGeneralPlot()
}

func drawGeneralPlot() {
	cmd := exec.Command("python", "/home/kudrix/GolandProjects/Algos2/src/plot/practice/drawGeneralPlotPractice.py")

	out, err := cmd.CombinedOutput()
	if err != nil {
		text := string(out)
		fmt.Println(text)
	}
}

func drawAvlTree() {
	rand.Seed(51)
	var data string

	for _, n := range util.ArrN {
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

	util.Write(data, "practiceData")

	cmd := exec.Command("python", "/home/kudrix/GolandProjects/Algos2/src/plot/practice/drawSinglePlotPractice.py", data, "AVL Дерево")

	out, err := cmd.CombinedOutput()

	util.WriteCSV("AVLTree", string(out))

	if err != nil {
		text := string(out)
		fmt.Println(text)
	}
}

func drawBsTree() {
	rand.Seed(51)
	var data string

	for _, n := range util.ArrN {
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

	util.Write(data, "practiceData")

	cmd := exec.Command("python", "/home/kudrix/GolandProjects/Algos2/src/plot/practice/drawSinglePlotPractice.py", data, "BS Дерево")

	out, err := cmd.CombinedOutput()

	util.WriteCSV("BSTree", string(out))

	if err != nil {
		text := string(out)
		fmt.Println(text)
	}
}

func drawRbTree() {
	rand.Seed(51)
	var data string

	for _, n := range util.ArrN {
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

	util.Write(data, "practiceData")

	cmd := exec.Command("python", "/home/kudrix/GolandProjects/Algos2/src/plot/practice/drawSinglePlotPractice.py", data, "RB Дерево")

	out, err := cmd.CombinedOutput()

	util.WriteCSV("RBTree", string(out))

	if err != nil {
		text := string(out)
		fmt.Println(text)
	}
}
