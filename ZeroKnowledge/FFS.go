package zeroknow

import (
	calc "code-bako/sandbox"
	"fmt"
	"math/rand"
	"time"
)

func FFS1() (N int, v []int, s []int) {
	/*
		初期段階
	*/
	p := 19 //calc.MakePrime()
	q := 13 //calc.MakePrime()
	N = p * q
	for i := 0; i < N/2; i++ { //ランダムにt個選択する(N/2は変更できる)
		tmp, _, _ := calc.ExtGCD(N, i)
		if tmp == 1 {
			s = append(s, i)
		}
	}
	for _, i := range s {
		ss := calc.Inverse(i, N)
		v = append(v, (ss*ss)%N)
	}
	return N, v, s
}

func FFSAP() {
	/*
		認証段階
	*/
	N, v, s := FFS1()
	fmt.Println("public-info is ", N)
	fmt.Println(v)
	fmt.Println(s)
	x, r := FFSSideA1(N)
	fmt.Println("(x,r) is ", x, r)
	e := FFSSideB1(s)
	fmt.Println(e)
	y := FFSSideA2(N, s, e, r)
	check := FFSSideB2(N, x, y, v, e)
	fmt.Println(check)
}

func FFSSideA1(N int) (int, int) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(N)
	return calc.FastPower(r, 2, N), r
}
func FFSSideB1(s []int) []int {
	rand.Seed(time.Now().UnixNano())
	t := len(s)
	var E []int
	for i := 0; i < t; i++ {
		E = append(E, rand.Intn(2))
	}
	return E
}
func FFSSideA2(N int, s []int, E []int, r int) int {
	ss := 1
	for i, e := range E {
		if e == 1 {
			ss = (ss * s[i]) % N
		}
	}
	return (ss * r) % N
}
func FFSSideB2(N int, x int, y int, v []int, E []int) bool {
	vv := 1
	for i, e := range E {
		if e == 1 {
			vv = (vv * v[i]) % N
		}
	}
	Y := (y * y) % N
	Y = (Y * vv) % N
	if x == Y {
		return true
	} else {
		return false
	}
}
