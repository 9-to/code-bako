package zeroknow

import (
	calc "code-bako/sandbox"
	"fmt"
	"math/rand"
	"time"
)

type pubInfoS struct {
	p int
	q int
	g int
}

func Schnorr1() (pi pubInfoS, v int, s int) {
	p, q, g := calc.MakeQ() //センタの公開情報
	fmt.Println("public-info is ", p, q, g)
	rand.Seed(time.Now().UnixNano())
	s = rand.Intn(q)
	ginv := calc.Inverse(g, p)
	v = calc.FastPower(ginv, s, p)
	fmt.Println("pk = ", v, ", secret = ", s)
	pi.p, pi.q, pi.g = p, q, g
	return pi, v, s
}

func SchnorrAP() {
	pi, v, s := Schnorr1()
	x, r := SSideA1(pi)
	c := SSideB1(pi)
	y := SSideA2(pi, r, s, c)
	check := SSideB2(pi, x, y, v, c)
	fmt.Println(check)
}

func SSideA1(pi pubInfoS) (x, r int) {
	/*
		証明者Aはrを選んでxを計算してBに送る
	*/
	rand.Seed(time.Now().UnixNano())
	r = rand.Intn(pi.q)
	x = calc.FastPower(pi.g, r, pi.p)
	return x, r
}

func SSideB1(pi pubInfoS) int {
	/*
		検証者Bはcを選んでAに送る
	*/
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(pi.q)
}

func SSideA2(pi pubInfoS, r int, s int, c int) int {
	/*
		Aは受け取ったcからyを計算する
	*/
	return (r + s*c) % pi.p
}

func SSideB2(pi pubInfoS, x int, y int, v int, c int) bool {
	/*
		Bは、x = g^yv^c mod p を検証する
	*/
	gg := calc.FastPower(pi.g, y, pi.p)
	vv := calc.FastPower(v, c, pi.p)
	if (gg*vv)%pi.p == x {
		return true
	} else {
		return false
	}
}
