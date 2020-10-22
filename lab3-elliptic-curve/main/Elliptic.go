package main

import "strconv"

func scalarMultiply(n int64, p Point, g Group) Point {
	var result Point
	addend := p
	for _, bit := range reverse(strconv.FormatInt(n, 2)) {
		if bit == '1' {
			result = add(result, addend, g)
		}
		addend = add(addend, addend, g)
	}
	return result
}

func add(p1 Point, p2 Point, g Group) Point {

	if (p1.x == 0 && p1.y == 0) || (p2.x == 0 && p2.y == 0) {
		if !(p1.x == 0 && p1.y == 0) {
			return p1
		} else {
			return p2
		}
	} else {
		if p1.x == p2.x && mod(p1.y+p2.y, g.m) == 0 {
			return Point{0, 0}
		}
	}

	var l int
	if p1 == p2 {
		l = mod((3*p1.x*p1.x+g.a)*modInverse(2*p1.y%g.m, g.m), g.m)
	} else {
		l = mod((p1.y-p2.y)*modInverse(p1.x-p2.x, g.m), g.m)
	}
	x3 := mod(l*l-p1.x-p2.x, g.m)
	y3 := mod(p1.y+l*(x3-p1.x), g.m)
	return Point{x3, mod(-y3, g.m)}
}
