package plot

import (
	"Algos2/src/trees"
	"Algos2/src/util"
	"fmt"
	"math"
	"math/rand"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

func DrawPractice() {
	util.Clear("avlData")
	util.Clear("rbData")
	drawTree("avl", trees.AVLTree{}, avlHeightCalculations)
	drawTree("rb", trees.RBTree{}, rbHeightCalculations)
}

func drawTree(treeType string, tree interface{}, heightCalculations []func(int) int) {
	rand.Seed(51)
	var data string

	for _, n := range util.ArrN {
		var arrXY []string
		keys := rand.Perm(n)
		sort.Ints(keys)

		switch t := tree.(type) {
		case trees.AVLTree:
			for _, key := range keys {
				t.Insert(key)
			}
			arrXY = append(arrXY, strconv.Itoa(n), strconv.Itoa(t.Height()))
		case trees.RBTree:
			for _, key := range keys {
				t.Insert(key)
			}
			arrXY = append(arrXY, strconv.Itoa(n), strconv.Itoa(t.Height()))
		}

		dataXY := strings.Join(arrXY, ",")
		data += dataXY + ";"
	}

	data = data[:len(data)-1]
	util.Write(data, treeType+"Data")
	data = ""

	for _, calc := range heightCalculations {
		for _, n := range util.ArrN {
			var arrXY []string
			arrXY = append(arrXY, strconv.Itoa(n), strconv.Itoa(calc(n)))
			dataXY := strings.Join(arrXY, ",")
			data += dataXY + ";"
		}
		data = data[:len(data)-1]
		util.Write(data, treeType+"Data")
		data = ""
	}

	cmd := exec.Command("python", "/home/kudrix/GolandProjects/Algos2/src/plot/drawGraphics.py", treeType)
	out, err := cmd.CombinedOutput()
	util.WriteCSV(strings.ToUpper(treeType)+"Tree", string(out))

	if err != nil {
		fmt.Println(string(out))
	}
}

var avlHeightCalculations = []func(int) int{
	func(n int) int { return int(math.Ceil(1.440*math.Log2(float64(n)+1.065) - 0.328)) },
	func(n int) int { return int(math.Ceil(math.Log2(float64(n)))) },
}

var rbHeightCalculations = []func(int) int{
	func(n int) int { return int(math.Ceil(2 * math.Log2(float64(n)+1))) },
	func(n int) int { return int(math.Ceil(math.Log2(float64(n)))) },
}
