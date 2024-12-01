package theory

import (
	"Algos2/src/util"
	"fmt"
	"math"
	"os/exec"
	"strconv"
	"strings"
)

func DrawTheory() {
	util.Clear("theoryData")
	drawAvlThree()
	drawRbThree()
	drawBsThree()
	drawGeneralPlot()
}

func drawGeneralPlot() {
	cmd := exec.Command("python", "/home/kudrix/GolandProjects/Algos2/src/plot/theory/drawGeneralPlotTheory.py")

	out, err := cmd.CombinedOutput()
	if err != nil {
		text := string(out)
		fmt.Println(text)
	}
}

func drawAvlThree() {
	var data string

	for _, n := range util.ArrN {
		var arrXY []string

		arrXY = append(arrXY, strconv.Itoa(n))
		arrXY = append(arrXY, strconv.Itoa(int(1.440*math.Log2(float64(n)+1.065)-0.328)))

		dataXY := strings.Join(arrXY, ",")
		data += dataXY + ";"
	}

	data = data[:len(data)-1]

	util.Write(data, "theoryData")

	cmd := exec.Command("python", "/home/kudrix/GolandProjects/Algos2/src/plot/theory/drawSinglePlotTheory.py", data, "AVL Дерево")

	out, err := cmd.CombinedOutput()
	if err != nil {
		text := string(out)
		fmt.Println(text)
	}
}

func drawBsThree() {
	var data string

	for _, n := range util.ArrN {
		var arrXY []string

		arrXY = append(arrXY, strconv.Itoa(n))
		arrXY = append(arrXY, strconv.Itoa(int(4.311*math.Log2(float64(n)))))

		dataXY := strings.Join(arrXY, ",")
		data += dataXY + ";"
	}

	data = data[:len(data)-1]

	util.Write(data, "theoryData")

	cmd := exec.Command("python", "/home/kudrix/GolandProjects/Algos2/src/plot/theory/drawSinglePlotTheory.py", data, "BS Дерево")

	out, err := cmd.CombinedOutput()
	if err != nil {
		text := string(out)
		fmt.Println(text)
	}
}

func drawRbThree() {
	var data string

	for _, n := range util.ArrN {
		var arrXY []string

		arrXY = append(arrXY, strconv.Itoa(n))
		arrXY = append(arrXY, strconv.Itoa(int(2*math.Log2(float64(n+1)))))

		dataXY := strings.Join(arrXY, ",")
		data += dataXY + ";"
	}

	data = data[:len(data)-1]

	util.Write(data, "theoryData")

	cmd := exec.Command("python", "/home/kudrix/GolandProjects/Algos2/src/plot/theory/drawSinglePlotTheory.py", data, "RB Дерево")

	out, err := cmd.CombinedOutput()
	if err != nil {
		text := string(out)
		fmt.Println(text)
	}
}
