package main

import (
	"math"
	"strings"
)

var Keys = []string{"dog", "norm", "hello", "format", "subline", "specific", "substring", "television", "intelligent",
	"abbreviation"}

const LGramLength = 5

const Tests = 10

var FreqEn = map[rune]float64{'a': 0.08167, 'b': 0.01492, 'c': 0.02782, 'd': 0.04253, 'e': 0.12702, 'f': 0.0228,
	'g': 0.02015, 'h': 0.06094, 'i': 0.06966, 'j': 0.00153, 'k': 0.00772, 'l': 0.04025, 'm': 0.02406, 'n': 0.06749,
	'o': 0.07507, 'p': 0.01929, 'q': 0.00095, 'r': 0.05987, 's': 0.06327, 't': 0.09056, 'u': 0.02758, 'v': 0.00978,
	'w': 0.0236, 'x': 0.0015, 'y': 0.01974, 'z': 0.00074}

const En = "abcdefghijklmnopqrstuvwxyz"
const Alphabet = int32(26)
const a = rune('a')
const z = rune('z')

func getFrequency(text []rune) map[rune]float64 {
	var freq = make(map[rune]float64)
	for _, c := range text {
		if strings.ContainsRune(En, c) {
			freq[c]++
		}
	}
	for i := a; i <= z; i++ {
		if _, pres := freq[i]; !pres {
			freq[i] = 0
		}
		freq[i] = freq[i] / float64(len(text))
	}
	return freq
}

func getNorm(freq map[rune]float64, freqEn map[rune]float64, shift int) float64 {
	diff := make([]float64, 0, len(freq))
	for i := 0; i < len(freq); i++ {
		j := (i - shift) % len(freq)
		diff = append(diff, math.Pow(freq[a+rune(i)]-freqEn[a+rune(j)], 2))
	}
	return sum(diff)
}

func sum(array []float64) float64 {
	var result float64 = 0
	for _, v := range array {
		result += v
	}
	return result
}

func Gcd(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return Gcd(b, a%b)
	}
}

func getIndexOfMaxElem(slice []int) int {
	pos := 0
	max := 0
	for i, n := range slice {
		if n > max {
			pos = i
			max = n
		}
	}
	return pos
}
