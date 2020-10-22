package main

import "fmt"

type Group struct {
	a int
	b int
	m int
}

type Point struct {
	x int
	y int
}

const Message = "Finally It's working!!!"

func main() {
	group := Group{
		a: 22,
		b: 14,
		m: 47,
	}
	points := generatePoints(group)
	n := findN(len(points))
	fmt.Printf("n=%v\n", n)
	g := findG(n, points, group)
	fmt.Printf("G=%v\n", g)
	g = Point{11, 6}
	a, b := swapKeys(n, g, group)
	fmt.Printf("Keys for user A:%v\nKeys for user B:%v\n", a, b)
	createSignature(n, g, group, a)
}

func generateAB(m int) []Point {
	ab := make([]Point, 0, 50)
	for a := 0; a < m; a++ {
		for b := 0; b < m; b++ {
			if mod(4*a*a*a+27*b*b, m) != 0 {
				ab = append(ab, Point{a, b})
			}
		}
	}
	return ab
}

func generatePoints(g Group) []Point {
	points := make([]Point, 0, 20)
	points = append(points, Point{0, 0})
	for x := 0; x < g.m; x++ {
		for _, y := range modSquareRoot(((x*(x*x%g.m)%g.m+(g.a*x)%g.m)%g.m+g.b)%g.m, g.m) {
			points = append(points, Point{x, y})
		}
	}
	return points
}
