package myaes

import (
	"fmt"
	"os"
)

type keylength struct {
	blen int //bit長
	wlen int //ワード長
	r    int //ラウンド数
}

const BROCK_LEN int = 128 //ブロック長

func KeyLen(n int) (kl keylength) {
	if n == 128 {
		kl.blen, kl.wlen, kl.r = n, 4, 10
	} else if n == 196 {
		kl.blen, kl.wlen, kl.r = n, 6, 12
	} else if n == 256 {
		kl.blen, kl.wlen, kl.r = n, 8, 14
	} else {
		fmt.Println(n, " is incrrect bit size.")
		os.Exit(0)
	}
	return kl
}
