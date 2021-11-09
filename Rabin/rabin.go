package main

import (
	"code-bako/Rabin/decoder"
	"code-bako/Rabin/encoder"
	"code-bako/Rabin/keygen"
	"fmt"
)

func main() {
	p, q, N := keygen.KeyGen()
	var m int
	fmt.Println("write message... m")
	fmt.Scanf("%d %d", &m)
	c := encoder.Encoder(m, N)
	mm := decoder.Decoder(c, p, q)
	fmt.Println("estimaded message is ", mm)
}
