package main

import (
	"code-bako/RSA/keygen"
	"fmt"
)

func main() {
	N, e, d := keygen.KeyGen()
	var m int
	fmt.Println("write message, m")
	fmt.Scanf("%d", &m)
	c := keygen.MyModulo(m, e, N)
	fmt.Println("cipher code is ", c)
	mm := keygen.MyModulo(c, d, N)
	fmt.Println("estimated message is ", mm)

}
