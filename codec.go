package main

import (
	calc "code-bako/sandbox"
	"fmt"
)

func main() {
	var a, x, N int
	fmt.Scanf("%d %d %d", &a, &x, &N)
	y, a := calc.SafePrime()
	fmt.Println(y, a)
}
