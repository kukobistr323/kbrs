package main

import (
	"fmt"
	util "kbrs/lab1-vigenere/main"
	"math"
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
	ab := generateAB(group.m)
	fmt.Printf("Length of (a,b) is %v\n", len(ab))
	//fmt.Println(ab)
	points := generatePoints(group)
	fmt.Printf("Length of points is %v\n", len(points))
	fmt.Println(points)
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

func movInverse(b int, m int) int {
	g := util.Gcd(b, m)
	if g != 1 {
		return -1
	} else {
		return
	}
}

func modDivide(a int, b int, m int) {
	a = a % m
inv:
}

//func scalarMultiply(n int, p Point) Point {
//
//}
//
//func add(p1 Point, p2 Point, g Group) Point {
//	var l
//	if p1 == p2 {
//
//	} else {
//		l = (p2.y + (-p1.y)%g.m) % g.m + (p2.x+(-p1.x)%g.m)%g.m
//	}
//}

//func scalarMultiply(n int64, p Point) Point {
//	for _, bit := range reverse(strconv.FormatInt(n, 2)) {
//		if bit == 1 {
//
//		}
//	}
//}
//
//func reverse(s string) string {
//	runes := []rune(s)
//	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
//		runes[i], runes[j] = runes[j], runes[i]
//	}
//	return string(runes)
//}
