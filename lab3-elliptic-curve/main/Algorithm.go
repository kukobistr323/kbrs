package main

import (
	"crypto/sha512"
	"encoding/binary"
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

func createSignature(message string, n int, g Point, group Group, a Keys) (int, int) {
	r := 0
	s := 0
	k := 0
	hash := hash(n, message)
	for s == 0 {
		for r == 0 {
			k = random(2, n-1)
			kG := scalarMultiply(int64(k), g, group)
			r = mod(kG.x, n)
		}
		s = mod(modInverse(k, n)*(hash+a.nA*r), n)
	}
	return r, s
}

func hash(n int, message string) int {
	byte64Hash := sha512.Sum512([]byte(message))
	byteHash := byte64Hash[:]
	hash := strconv.FormatUint(binary.BigEndian.Uint64(byteHash), 2)[:len(strconv.FormatInt(int64(n), 2))]
	hashInt, _ := strconv.ParseInt(hash, 2, 64)
	return int(hashInt)
}

func checkSignature(message string, r, s, n int, g Point, group Group, a Keys) int {
	hash := hash(n, message)
	w := mod(modInverse(s, n), n)
	u1 := mod(hash*w, n)
	u2 := mod(r*w, n)
	p := add(scalarMultiply(int64(u1), g, group), scalarMultiply(int64(u2), a.pA, group), group)
	return mod(p.x, n)
}
