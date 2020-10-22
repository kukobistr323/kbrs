package main

import (
	"fmt"
)

type Group struct {
	a int
	b int
	m int
}

type Point struct {
	x int
	y int
}

func main() {
	group := Group{
		a: 9,
		b: 17,
		m: 23,
	}
	//ab := generateAB(group.m)
	//fmt.Printf("Length of (a,b) is %v\n", len(ab))
	//fmt.Println(ab)
	//points := generatePoints(group)
	//fmt.Printf("Length of points is %v\n", len(points))
	//fmt.Println(points)

	fmt.Println(add(Point{16, 5}, Point{16, 5}, group))
}

func generateAB(m int) []Point {
	ab := make([]Point, 0, 50)
	for a := 0; a < m; a++ {
		for b := 0; b < m; b++ {
			if (27*(a*(a*a%m)%m)%m+4*(b*b%m)%m)%m != 0 {
				ab = append(ab, Point{a, b})
			}
		}
	}
	return ab
}

func generatePoints(g Group) []Point {
	points := make([]Point, 0, 20)
	for x := 0; x < g.m; x++ {
		for _, y := range modSquareRoot(((x*(x*x%g.m)%g.m+(g.a*x)%g.m)%g.m+g.b)%g.m, g.m) {
			points = append(points, Point{x, y})
		}
	}
	return points
}

func modSquareRoot(n int, p int) []int {
	squares := make([]int, 0)
	n = n % p
	for i := 0; i < p; i++ {
		if (i*i)%p == n {
			squares = append(squares, i)
		}
	}
	return squares
}

func modInverse(b int, m int) int {
	var x, y int
	g := gcdExtended(b, m, &x, &y)
	if g != 1 {
		return -1
	}

	return (x%m + m) % m
}

func modDivide(a int, b int, m int) int {
	a = a % m
	inv := modInverse(b, m)
	if inv == -1 {
		return -1
	} else {
		return (inv * a) % m
	}
}

func gcdExtended(a, b int, x, y *int) int {
	if a == 0 {
		*x = 0
		*y = 1
		return b
	}
	var x1, y1 int
	gcd := gcdExtended(b%a, a, &x1, &y1)
	*x = y1 - (b/a)*x1
	*y = x1
	return gcd
}

//func scalarMultiply(n int, p Point) Point {
//
//}

func add(p1 Point, p2 Point, g Group) Point {
	var l int
	if p1 == p2 {
		l = (3*(p1.x*p1.x%g.m)%g.m + g.a) % g.m * modInverse(2*p1.y%g.m, g.m) % g.m
	} else {
		l = ((p2.y + (-p1.y)%g.m) % g.m * modInverse((p2.x+(-p1.x)%g.m)%g.m, g.m)) % g.m
	}
	x3 := ((l*l%g.m-p1.x%g.m+g.m)%g.m - p2.x%g.m + g.m) % g.m
	y3 := l * (((p1.x-x3+g.m)%g.m - p1.y + g.m) % g.m) % g.m
	return Point{x3, y3}
}
