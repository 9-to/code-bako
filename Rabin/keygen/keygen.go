package keygen

import (
	calc "code-bako/sandbox"
	"fmt"
	"os"
)

func KeyGen() (p, q, N int) {
	fmt.Println("write prime numbers... p, q")
	fmt.Scanf("%d %d", &p, &q)
	if !calc.PrimeCheck(p, 1000) {
		os.Exit(0)
	} else if !calc.PrimeCheck(q, 1000) {
		os.Exit(0)
	}
	N = p * q
	return p, q, N
}
