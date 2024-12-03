package main

import (
	"Algos2/src/plot/practice"
	"Algos2/src/plot/theory"
	"Algos2/src/test"
)

func main() {
	practice.DrawPractice()
	theory.DrawTheory()
	startTest()
}

func startTest() {
	test.Avl()
	test.Bs()
	test.Rb()
}
