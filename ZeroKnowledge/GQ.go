package zeroknow

import (
	"code-bako/RSA/keygen"
	calc "code-bako/sandbox"
	"fmt"
	"math/rand"
	"time"
)

type pubInfoGQ struct {
	N int
	e int
}

func GQ1() (pk pubInfoGQ, s int, v int) {
	/*
		初期段階
	*/
	N, e, _ := keygen.KeyGen()
	rand.Seed(time.Now().UnixNano())
	for {
		s = rand.Intn(N-1) + 1
		tmp, _, _ := calc.ExtGCD(N, s)
		if tmp == 1 {
			break
		}
	}
	sinv := calc.Inverse(s, N)
	v = calc.FastPower(sinv, e, N)
	pk.N, pk.e = N, e
	return pk, s, v
}

func GQAP() {
	/*
		AB間の認証プロトコル
	*/
	pk, s, v := GQ1()
	fmt.Println("public-info is ", pk.N, pk.e)
	fmt.Println("public-key is ", v, ", secret is ", s)
	x, r := GQSideA1(pk)
	c := GQSideB1(pk)
	y := GQSideA2(pk, s, c, r)
	check := GQSideB2(pk, x, y, v, c)
	fmt.Println(check)
}

func GQSideA1(pk pubInfoGQ) (x int, r int) {
	rand.Seed(time.Now().UnixNano())
	for {
		r = rand.Intn(pk.N-1) + 1
		tmp, _, _ := calc.ExtGCD(pk.N, r)
		if tmp == 1 {
			break
		}
	}
	return calc.FastPower(r, pk.e, pk.N), r
}
func GQSideB1(pk pubInfoGQ) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(pk.e)
}
func GQSideA2(pk pubInfoGQ, s int, c int, r int) int {
	ss := calc.FastPower(s, c, pk.N)
	return (r * ss) % pk.N
}
func GQSideB2(pk pubInfoGQ, x int, y int, v int, c int) bool {
	vv := calc.FastPower(v, c, pk.N)
	yy := calc.FastPower(y, pk.e, pk.N)
	if x == (vv*yy)%pk.N {
		return true
	} else {
		return false
	}
}
