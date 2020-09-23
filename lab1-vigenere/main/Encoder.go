package main

import "strings"

func encrypt(text string, key string) string {
	result := []rune(strings.ToLower(text))

	for i := 0; i < len(result); i++ {
		if strings.ContainsRune(En, result[i]) {
			result[i] = (((result[i] - a) + (rune(key[i%len(key)]) - a)) % Alphabet) + a
		}
	}
	return string(result)
}
