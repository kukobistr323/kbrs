package main

import (
	"os"
	"strconv"
)

func createFilePath(prefix string, suffix string, number int) string {
	return prefix + strconv.Itoa(number) + suffix
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
