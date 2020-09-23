package main

import (
	"fmt"
	"log"
)

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
	fmt.Println(decrypt(text, keySize))
}
