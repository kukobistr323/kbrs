package main

import (
	"io/ioutil"
	"os"
)

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
