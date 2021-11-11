package calc

/*
楕円加算
*/
type PP struct {
	X int
	Y int
	O bool //単位元かどうかのフラグ
}

type EQ struct {
	/*
		楕円曲線の標準形
		E; f(x)=x^3+ax+b
	*/
	A int
	B int
}

func EQAdd(a PP, b PP, f EQ, p int) PP {
	/*
		楕円加算
		-1は単位元　O　を表す
	*/
	if a.O == true {
		return b
	} else if b.O == true {
		return a
	}
	if a.X == b.X && 0 == (a.Y+b.Y)%p {
		var out PP
		out.X, out.Y, out.O = -1, -1, true
		return out
	}
	var lambda int
	var out PP
	out.O = false
	if a.X == b.X && a.Y == b.Y {
		lambda = ((3*a.X*a.X + f.A) * Inverse(2*a.Y, p)) % p
	} else {
		lambda = ((p + b.Y - a.Y) * Inverse(p+b.X-a.X, p)) % p
	}
	out.X = (10*p + (lambda*lambda - a.X - b.X)) % p
	out.Y = (10*p + (lambda*(a.X-out.X) - a.Y)) % p
	return out
}
