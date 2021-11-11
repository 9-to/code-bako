package main

import (
	calc "code-bako/sandbox"
	"fmt"
)

func main() {
	var f calc.EQ
	a := calc.PP{X: 5, Y: 9, O: false}
	b := calc.PP{X: 5, Y: 9, O: false}
	f = calc.EQ{A: 2, B: 1, Prime: 11}
	for i := 0; i < 16; i++ {
		fmt.Println(i, ":", a)
		a = calc.EQAdd(a, b, f, 11)
	}
}
