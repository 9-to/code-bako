package keygen

import (
	calc "code-bako/sandbox"
	"fmt"
	"os"
)

func KeyGen() (N, e, d int) {
	fmt.Println("write prime numbers... p, q")
	var p, q int
	fmt.Scanf("%d %d", &p, &q)
	if !calc.PrimeCheck(p, 1000) {
		os.Exit(0)
	}
	if !calc.PrimeCheck(q, 1000) {
		os.Exit(0)
	}
	N = p * q
	tmp := 0
	for tmp != 1 {
		fmt.Println("write public-key... e")
		fmt.Scanf("%d", &e)
		tmp, _, _ = calc.ExtGCD((p-1)*(q-1), e)
	}
	_, _, d = calc.ExtGCD((p-1)*(q-1), e)
	d = (((p - 1) * (q - 1)) + d) % ((p - 1) * (q - 1))
	fmt.Println("secret-key d is ", d)
	return N, e, d
}

func MyModulo(m, a, N int) (c int) {
	/*
		mをa乗してNの法をとる
	*/
	tmp := 1
	for i := 0; i < a; i++ {
		tmp = (tmp * m) % N
	}
	c = tmp
	return c
}
