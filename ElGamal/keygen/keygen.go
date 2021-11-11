package keygen

import (
	calc "code-bako/sandbox"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type pkEx struct {
	/*
		拡張ElGamal暗号の公開鍵
	*/
	p int
	q int
	a int
	y int
}

type PkEC struct {
	/*
		楕円ElGamal暗号の公開鍵
	*/
	F  calc.EQ
	Pg calc.PP
	Py calc.PP
}

func KeyGenerate() (p int, g int, y int) {
	/*
		鍵生成
		p,g,yを返す
	*/
	fmt.Println("write p, g")
	//var p int
	//var g int
	fmt.Scanf("%d %d", &p, &g)
	if !calc.PrimeCheck(p, 1000) {
		os.Exit(0)
	}
	if g <= 0 || p <= g {
		fmt.Println("gの数が不適正です")
		os.Exit(2)
	}
	fmt.Println("write secret-key x")
	var x int
	fmt.Scanf("%d", &x)
	if x < 0 || p-1 <= x {
		fmt.Println("xの数が不適正です")
		os.Exit(3)
	}
	y = CalcPublicKeyY(p, g, x)
	fmt.Println("public-key y is ", y)
	return p, g, y
}

func KeyGenEx() (pk pkEx, x int) {
	/*
		拡張ElGamal暗号の鍵生成
	*/
	pk.p, pk.q, _ = calc.SafePrime()
	rand.Seed(time.Now().UnixNano())
	flag := false
	var a int
	for flag == false {
		a = rand.Intn(pk.p-1) + 1
		if calc.FastPower(a, pk.q, pk.p) == 1 {
			flag = true
		}
	}
	x = rand.Intn(pk.q)
	pk.y = calc.FastPower(a, x, pk.p)
	pk.a = a
	fmt.Println("public-key is (", pk, ")")
	fmt.Println("secret-key is ", x)
	return pk, x
}

func KeygenEC() (pk PkEC, sk int) {
	/*
		楕円ElGamal暗号の鍵生成
		///
		p:大きな素数
		f:GF(p)上の楕円曲線E(Fp)
		Pg:E(Fp)有理点=生成元
		q:Pgの位数
	*/
	var p, a, b int
	rand.Seed(time.Now().UnixNano())
	fmt.Println("write p")
	fmt.Scanf("%d", &p)
	if !calc.PrimeCheck(p, 1000) {
		os.Exit(0)
	}
	fmt.Println("write elliptic curve... f= x^3+ax+b\n (a,b) is ?")
	fmt.Scanf("%d %d", &a, &b)
	pk.F.A, pk.F.B, pk.F.Prime = a, b, p
	fmt.Println("Generator Pg is ... (x,y)")
	fmt.Scanf("%d %d", &pk.Pg.X, &pk.Pg.Y)
	pk.Pg.O = false
	fmt.Println("secret-key is ...")
	fmt.Scanf("%d", &sk)
	pk.Py = calc.EQTimes(pk.Pg, sk, pk.F)
	return pk, sk
}

func CalcPublicKeyY(p, g, x int) int {
	/*
		公開鍵yを計算する
	*/
	y := 1
	for i := 0; i < x; i++ {
		y = (g * y) % p
	}
	return y
}
