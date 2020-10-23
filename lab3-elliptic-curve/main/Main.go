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
		a: 17,
		b: 8,
		m: 61,
	}
	points := generatePoints(group)
	fmt.Printf("size of points array: %v\n", len(points))
	n := findN(len(points))
	fmt.Printf("n=%v\n", n)
	g := findG(n, points, group)
	fmt.Printf("G=%v\n", g)
	a, b := swapKeys(n, g, group)
	fmt.Printf("Keys for user A:%v\nKeys for user B:%v\n", a, b)
	r, s := createSignature(Message, n, g, group, a)
	fmt.Printf("Signature: (r=%v,s=%v)\n", r, s)
	r1 := checkSignature(Message, r, s, n, g, group, a)
	wrongR := checkSignature(Message, 5, 12, n, g, group, a)
	fmt.Printf("r*=%v\n", r1)
	fmt.Printf("wrong r*=%v\n", wrongR)
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
		for _, y := range modSquareRoot(mod(x*x*x+g.a*x+g.b, g.m), g.m) {
			points = append(points, Point{x, y})
		}
	}
	return points
}
