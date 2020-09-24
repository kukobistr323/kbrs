package main

import (
	"os"
)

const PrefixInput = "lab1-vigenere/resources/text"
const SuffixInput = ".txt"
const PrefixOutput = "lab1-vigenere/resources/"
const SuffixOutput = ".png"

func createFilePath(prefix string, suffix string, name string) string {
	return prefix + name + suffix
}

func readFile(path string, charSize int) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	text := make([]byte, charSize)
	_, err = file.Read(text)
	if err != nil {
		return "", err
	}
	return string(text), nil
}
