package keygen

import (
	"fmt"
	"os"
)

func KeyGen() (N, e, d int) {
	fmt.Println("write prime numbers... p, q")
	var p, q int
	fmt.Scanf("%d %d", &p, &q)
	CheckPrime(p)
	CheckPrime(q)
	N = p * q
	tmp := 0
	for tmp != 1 {
		fmt.Println("write public-key... e")
		fmt.Scanf("%d", &e)
		tmp, _, _ = ExtGCD((p-1)*(q-1), e)
	}
	_, _, d = ExtGCD((p-1)*(q-1), e)
	d = (((p - 1) * (q - 1)) + d) % ((p - 1) * (q - 1))
	fmt.Println("secret-key d is ", d)
	return N, e, d
}

func CheckPrime(p int) int {
	/*
		素数を判定する
	*/
	if p%2 == 0 { //とりあえず2で割り切れる
		fmt.Println("pが素数ではありません")
		os.Exit(1)
	}
	return 0
}

func Inverse(p, N int) int {
	/*
		pの逆元を出力する
	*/
	_, a, _ := ExtGCD(p, N)
	return a
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
