package main

import (
	"math/big"
	"math/rand"
	"time"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func primeDividers(n int) []int {

	dividers := make([]int, 0)
	for i := 2; i < n; i++ {
		if n%i == 0 && big.NewInt(int64(i)).ProbablyPrime(0) {
			dividers = append(dividers, i)
		}
	}
	return dividers
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
