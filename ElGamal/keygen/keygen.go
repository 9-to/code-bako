package keygen

import (
	"fmt"
	"os"
)

func KeyGenerate() (p int, g int, y int) {
	/*
		鍵生成
		p,g,yを返す
	*/
	fmt.Println("write p, g")
	//var p int
	//var g int
	fmt.Scanf("%d %d", &p, &g)
	CheckPrime(p)
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
	return p, g, y
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
