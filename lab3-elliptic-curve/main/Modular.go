package main

func mod(a int, m int) int {
	if a >= 0 {
		return a % m
	} else {
		return m + a%m
	}
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
	res := gcdExtended(b, m)
	if res[0] != 1 {
		panic("PIZDEC")
	} else {
		return res[1] % m
	}
}

func gcdExtended(p, q int) []int {
	if q == 0 {
		return []int{p, 1, 0}
	}
	gcd := gcdExtended(q, mod(p, q))
	return []int{gcd[0], gcd[2], gcd[1] - (p/q)*gcd[2]}
}
