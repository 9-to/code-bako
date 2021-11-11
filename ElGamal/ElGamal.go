package main

import (
	"code-bako/ElGamal/keygen"
	calc "code-bako/sandbox"
	"fmt"
)

func main() {
	ElGamalEC()
}

func ElGamal() {
	/*
		ElGamal暗号
	*/
}

func ElGamalEC() {
	/*
		楕円Elgamal暗号
	*/
	pk, sk := keygen.KeygenEC()
	var m calc.PP
	fmt.Scanf("%d %d", &m.X, &m.Y)
	m.O = false
	c := keygen.EncoderElGamalEC(pk, m)
	mm := keygen.DecoderElGamalEC(c, pk, sk)
	fmt.Println(mm)

}

func ElGamalEx() {
	/*
		拡張ElGamal暗号
	*/
	pk, sk := keygen.KeyGenEx()
	var m int
	fmt.Println("wrire message m")
	fmt.Scanf("%d", &m)
	c := keygen.EncoderElGamalEx(pk, m)
	keygen.DecoderElGamalEx(c, pk, sk)
}
