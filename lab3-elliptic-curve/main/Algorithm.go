package main

import (
	"crypto/sha512"
	"fmt"
	"strconv"
)

type Keys struct {
	nA int
	pA Point
	k  Point
}

func findN(n int) int {
	dividers := primeDividers(n)
	if len(dividers) == 0 {
		return n
	} else {
		return dividers[len(dividers)-1]
	}
}

func findG(n int, points []Point, group Group) Point {
	g := Point{0, 0}
	for g.x == 0 && g.y == 0 {
		h := len(points) / n
		i := random(1, len(points))
		g = scalarMultiply(int64(h), points[i], group)
	}
	return g
}

func swapKeys(n int, g Point, group Group) (Keys, Keys) {
	a := Keys{}
	b := Keys{}
	//a.nA = random(1,n)
	//b.nA = random(1,n)
	a.nA = 9
	b.nA = 6

	a.pA = scalarMultiply(int64(a.nA), g, group)
	b.pA = scalarMultiply(int64(b.nA), g, group)

	a.k = scalarMultiply(int64(a.nA), b.pA, group)
	b.k = scalarMultiply(int64(b.nA), a.pA, group)

	return a, b
}

func createSignature(n int, g Point, group Group, a Keys) int {
	r := 0
	k := 0
	hash := hash(n, Message)
	fmt.Println(hash)
	for r == 0 {
		k = random(2, n-1)
		kG := scalarMultiply(int64(k), g, group)
		r = mod(kG.x, n)
	}
	//s := mod(modInverse(k, n)*(hash(n, Message)+a.nA*r), n)
	return 1
}

func hash(n int, message string) int {
	byte64Hash := sha512.Sum512([]byte(message))
	byteHash := byte64Hash[:]
	hash, _ := strconv.Atoi(string(byteHash[:len(strconv.FormatInt(int64(n), 2))]))
	return hash
}
