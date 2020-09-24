package main

import (
	"fmt"
	"log"
)

var hits = make([]int, Tests)

func main() {
	//for i, key := range Keys {
	//	for j := 0; j < 2; j++ {
	//		go decrypt(i, j, key)
	//	}
	//}
	//time.Sleep(20 * time.Second)
	decrypt(0, 0, Keys[0])
	fmt.Println(hits)
}

func decrypt(charSize int, textNumber int, key string) {
	text, err := readFile(createFilePath(Prefix, Suffix, textNumber+1), (charSize+1)*1000)
	if err != nil {
		log.Fatal(err)
	}
	encrypted := encrypt(text, key)

	keySize := getKeySize(encrypted, LGramLength)

	decryptedKey := decryptKey(encrypted, keySize)

	fmt.Println(decryptedKey)
	fmt.Println(len(decryptedKey))
	fmt.Println(Keys[charSize])
	fmt.Println(len(key))

	if decryptedKey == key {
		hits[charSize]++
	}
}
