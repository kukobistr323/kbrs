package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const KEY = "oil"
const INPUT = "lab1-vigenere/resources/input.txt"
const ENCRYPTED = "lab1-vigenere/resources/encrypted.txt"
const EN = "abcdefghijklmnopqrstuvwxyz"
const ALPHABET = int32(26)
const a = int32('a')

func main() {
	text, err := readFile(INPUT)
	if err != nil {
		log.Fatal(err)
	}
	if writeToFile(ENCRYPTED, encrypt(text, KEY)) != nil {
		log.Fatal(err)
	}
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
		if strings.ContainsRune(EN, result[i]) {
			result[i] = (((result[i] - a) + (int32(key[i%len(key)]) - a)) % ALPHABET) + a
		}
	}
	return string(result)
}
