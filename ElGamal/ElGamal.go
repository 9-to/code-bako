package main

import (
	"code-bako/ElGamal/keygen"
	"fmt"
)

func main() {
	ElgamalEx()
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
}

func ElgamalEx() {
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
