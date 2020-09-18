package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const Key = "oil"
const LGramLength = 3

const Input = "lab1-vigenere/resources/input.txt"
const Encrypted = "lab1-vigenere/resources/encrypted.txt"

var FreqEn = map[rune]float64{'a': 0.08167, 'b': 0.01492, 'c': 0.02782, 'd': 0.04253, 'e': 0.12702, 'f': 0.0228,
	'g': 0.02015, 'h': 0.06094, 'i': 0.06966, 'j': 0.00153, 'k': 0.00772, 'l': 0.04025, 'm': 0.02406, 'n': 0.06749,
	'o': 0.07507, 'p': 0.01929, 'q': 0.00095, 'r': 0.05987, 's': 0.06327, 't': 0.09056, 'u': 0.02758, 'v': 0.00978,
	'w': 0.0236, 'x': 0.0015, 'y': 0.01974, 'z': 0.00074}

const En = "abcdefghijklmnopqrstuvwxyz"
const Alphabet = int32(26)
const a = int32('a')

type FreqEntry struct {
	letter rune
	freq   int
}

func main() {
	text, err := readFile(Input)
	if err != nil {
		log.Fatal(err)
	}
	if writeToFile(Encrypted, encrypt(text, Key)) != nil {
		log.Fatal(err)
	}

	text, err = readFile(Encrypted)
	if err != nil {
		log.Fatal(err)
	}

	keySize := getKeySize(text, LGramLength)
	fmt.Println(keySize)
	fmt.Println(findKeyWord(text, keySize))
}

func readFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func writeToFile(path string, text string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = file.WriteString(text)
	if err != nil {
		file.Close()
		return err
	}
	err = file.Close()
	return err
}

func encrypt(text string, key string) string {
	result := []int32(strings.ToLower(text))

	for i := 0; i < len(result); i++ {
		if strings.ContainsRune(En, result[i]) {
			result[i] = (((result[i] - a) + (int32(key[i%len(key)]) - a)) % Alphabet) + a
		}
	}
	return string(result)
}

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

func findKeyWord(text string, keySize int) string {

	caesar := make([][]rune, keySize, keySize)

	for i, c := range text {
		caesar[i%keySize] = append(caesar[i%keySize], c)
	}

	return ""
}

func getFrequency(text string) map[rune]int {
	var freq = make(map[rune]int)
	for _, c := range text {
		if strings.ContainsRune(En, c) {
			freq[c]++
		}
	}
	return freq
}

//func decryptCaesar(text string) string {
//	freq := getFrequency(text)
//
//}

//func sortLettersByFreq(freq map[rune]int) []rune {
//	freqSlice := make([]FreqEntry, 0, len(freq))
//	for k, v := range freq {
//		freqSlice = append(freqSlice, FreqEntry{
//			letter: k,
//			freq:   v,
//		})
//	}
//	sort.Slice(freqSlice, func(i, j int) bool {
//		return freqSlice[i].freq < freqSlice[j].freq
//	})
//	letters := make([]rune, 0, len(freqSlice))
//	for _, elem := range freqSlice {
//		letters := append(letters, elem.letter)
//	}
//
//}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
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
