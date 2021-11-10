package keygen

import (
	calc "code-bako/sandbox"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type cipherC struct {
	c1 int
	c2 int
}

func main() {
	var p, g, y int
	p, g, y = KeyGenerate()
	fmt.Println("write r")
	var r, m int
	fmt.Scanf("%d", &r)
	if r < 0 || p-1 <= r {
		fmt.Println("rの数が不適正です")
		os.Exit(4)
	}
	fmt.Println("wrire message m")
	fmt.Scanf("%d", &m)
	if m < 0 || p <= m {
		fmt.Println("mが不適正です")
		os.Exit(5)
	}
	c1 := genC1(p, g, r)
	c2 := genC2(m, p, y, r)
	fmt.Println(c1, c2)
}

func EncoderElGamalEx(pk pkEx, m int) (c cipherC) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(pk.q)
	fmt.Println("r is", r)
	c.c1 = calc.FastPower(pk.a, r, pk.p)
	c.c2 = (calc.FastPower(pk.y, r, pk.p) * m) % pk.p
	fmt.Println("cipher-code is", c)
	return c
}

func genC1(p, g, r int) (c1 int) {
	c1 = 1
	for i := 0; i < r; i++ {
		c1 = (g * c1) % p
	}
	return c1
}

func genC2(m, p, y, r int) (c2 int) {
	c2 = 1
	for i := 0; i < r; i++ {
		c2 = (y * c2) % p
	}
	return (c2 * m) % p
}
