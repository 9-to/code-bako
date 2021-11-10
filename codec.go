package main

import (
	calc "code-bako/sandbox"
	"fmt"
)

func main() {
	var a, x, N int
	fmt.Scanf("%d %d %d", &a, &x, &N)
	y := calc.Sqrt(a, x)
	fmt.Println(y)
}
