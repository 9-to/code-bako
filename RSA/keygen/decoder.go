package keygen

import (
	calc "code-bako/sandbox"
	"fmt"
)

func DecoderRSA(c, d, N int) int {
	mm := MyModulo(c, d, N)
	fmt.Println("estimated message is ", mm)
	return mm
}

func DecoderRSAP(c, d, e, N int) int {
	/*
		/ rを計算する
		/ C = r^e(1+mN) mod N^2
	*/
	r := calc.FastPower(c, d, N)
	re := calc.FastPower(r, e, N*N)
	var mm int
	for i := 1; i < N; i++ {
		gcd, _, _ := calc.ExtGCD(i, N)
		if gcd == 1 {
			tmp := (re * (1 + i*N)) % (N * N)
			if tmp == c {
				mm = i
				break
			}
		}
	}
	fmt.Println("estimated message is ", mm)
	return mm
}
