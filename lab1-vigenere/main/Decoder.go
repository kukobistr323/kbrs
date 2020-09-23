package main

import (
	"fmt"
	"strings"
)

func getKeySize(text string, lGramLength int) int {
	repeat := make([]int, 0, len(text))
	size := len(text) - lGramLength + 1
	for i := 0; i < size; i++ {
		first := text[i : i+lGramLength]
		for j := i + 1; j < size; j++ {
			second := text[j : j+lGramLength]
			if first == second {
				repeat = append(repeat, j-i)
			}
		}
	}
	nods := make([]int, len(text))
	for i := 0; i < len(repeat); i++ {
		for j := i + 1; j < len(repeat); j++ {
			nods[gcd(repeat[i], repeat[j])]++
		}
	}
	return getIndexOfMaxElem(nods)
}

func decrypt(text string, keySize int) string {

	caesar := make([][]rune, keySize, keySize)

	for i, c := range text {
		caesar[i%keySize] = append(caesar[i%keySize], c)
	}

	decryptedCaesar := make([][]rune, keySize, keySize)

	for i, str := range caesar {
		decryptedCaesar[i%keySize] = decryptCaesar(str)
	}

	decrypted := make([]rune, 0, len(text))

	for i := 0; i < len(text); i++ {
		decrypted = append(decrypted, decryptedCaesar[i%keySize][i/keySize])
	}
	fmt.Println(string(decrypted))
	return ""
}

func decryptCaesar(text []rune) []rune {
	shift := getKeyLetter(text)
	decrypted := make([]rune, 0, len(text))

	for _, c := range text {
		if strings.ContainsRune(En, c) {
			decrypted = append(decrypted, (c-shift)%Alphabet+a)
		} else {
			decrypted = append(decrypted, c)
		}
	}
	return decrypted
}

func getKeyLetter(text []rune) rune {
	freq := getFrequency(text)
	var decrKeyLetter rune
	bestDiff := 1.

	for i := range freq {
		currDif := getNorm(freq, FreqEn, int(i-a))
		if currDif < bestDiff {
			bestDiff = currDif
			decrKeyLetter = i
		}
	}
	fmt.Printf("%c\n", decrKeyLetter)
	return decrKeyLetter
}
