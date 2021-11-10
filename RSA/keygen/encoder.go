package keygen

import (
	calc "code-bako/sandbox"
	"fmt"
	"math/rand"
	"os"
)

func EncoderRSA(m, e, N int) int {
	/*
		RSA暗号
	*/
	c := MyModulo(m, e, N)
	fmt.Println("cipher code is ", c)
	return c
}

func EncoderRSAP(m, e, N int) int {
	/*
		RSA暗号
	*/
	//check message
	if m < 0 || N < m {
		os.Exit(0)
	}
	a, _, _ := calc.ExtGCD(m, N)
	if a != 1 {
		os.Exit(1)
	}
	r := rand.Intn(N-1) + 1
	r = 3 //
	c := calc.FastPower(r, e, N*N)
	c = (c * (1 + m*N)) % (N * N)
	fmt.Println("cipher code is ", c)
	return c
}
