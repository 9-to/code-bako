package encoder

import (
	calc "code-bako/sandbox"
	"fmt"
	"os"
)

func Encoder(m, N int) int {
	var c int
	c = calc.FastPower(m, 2, N)
	fmt.Println("cipher-code is ... ", c)
	return c
}

func Encoder2(m, N int) int {
	/*
		制限Rabin暗号
	*/
	if 0 <= m || N/2 <= m {
		os.Exit(0)
	}
	if calc.Legendre(m, N) != 1 {
		os.Exit(1)
	}
	var c int
	c = calc.FastPower(m, 2, N)
	fmt.Println("cipher-code is ... ", c)
	return c
}
