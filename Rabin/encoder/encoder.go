package encoder

import (
	calc "code-bako/sandbox"
	"fmt"
)

func Encoder(m, N int) int {
	var c int
	c = calc.FastPower(m, 2, N)
	fmt.Println("cipher-code is ... ", c)
	return c
}
