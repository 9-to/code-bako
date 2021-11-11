package keygen

import (
	calc "code-bako/sandbox"
	"fmt"
)

func DecoderElGamal() {
	var c1, c2, x, p int
	fmt.Println("write p")
	fmt.Scanf("%d", &p)
	fmt.Println("write cipher-code c1, c2")
	fmt.Scanf("%d %d", &c1, &c2)
	fmt.Println("write secret-key x")
	fmt.Scanf("%d", &x)
	m := 1
	for i := 0; i < (p - 1 - x); i++ {
		m = (m * c1) % p
	}
	m = (m * c2) % p
	fmt.Println("message is ", m)
}

func DecoderElGamalEx(c cipherC, pk pkEx, sk int) (m int) {
	cUnder := calc.Inverse(calc.FastPower(c.c1, sk, pk.p), pk.p)
	m = (c.c2 * cUnder) % pk.p
	m = (pk.p + m) % pk.p
	fmt.Println("Estimated-message is ", m)
	return m
}

func DecoderElGamalEC(c CipherEC, pk PkEC, sk int) (m calc.PP) {
	xC1 := calc.EQTimes(c.C1, sk, pk.F)
	m.X, m.Y = (pk.F.Prime+c.C2.X-xC1.X)%pk.F.Prime, (pk.F.Prime+c.C2.Y-xC1.Y)%pk.F.Prime
	return m
}
