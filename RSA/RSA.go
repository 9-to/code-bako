package main

import (
	"code-bako/RSA/keygen"
	"fmt"
)

func main() {
	N, e, d := keygen.KeyGen()
	var m int
	fmt.Println("write message, m")
	fmt.Scanf("%d", &m)
	c := keygen.EncoderRSAP(m, e, N)
	keygen.DecoderRSAP(c, d, e, N)
}
