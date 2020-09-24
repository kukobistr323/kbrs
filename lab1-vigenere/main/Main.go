package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
)

var hits = ConcurrentSlice{items: make([]int, Tests*Tests, Tests*Tests)}

func main() {
	var wg sync.WaitGroup

	for i := range Keys {
		for j := 0; j < Tests; j++ {
			for k := 0; k < Tests; k++ {
				wg.Add(1)
				go decrypt(j, k, i, &wg)
			}
		}
	}
	wg.Wait()
	fmt.Println(hits.items)

	//var result = []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 9, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 8, 10, 10, 10, 10, 10, 10, 10, 10, 10, 4, 8, 10, 10, 9, 9, 9, 9, 9, 10, 8, 8, 10, 10, 10, 10, 10, 10, 10, 10, 1, 5, 7, 8, 8, 8, 9, 9, 9, 10, 2, 5, 7, 9, 9, 9, 10, 10, 10, 10, 4, 8, 9, 9, 10, 10, 10, 10, 10, 10, 1, 6, 9, 9, 9, 9, 9, 9, 9, 10}
	drawPlot(preparePointsFirstPlot(hits.items), "Key Length", "Hits", 13, 2,
		createFilePath(PrefixOutput, SuffixOutput, "first"))
	drawPlot(preparePointsSecondPlot(hits.items), "Text Length", "Hits", 11000, 2,
		createFilePath(PrefixOutput, SuffixOutput, "second"))
}

func decrypt(charSize int, textNumber int, keyInd int, wg *sync.WaitGroup) {
	defer wg.Done()

	text, err := readFile(createFilePath(PrefixInput, SuffixInput, strconv.Itoa(textNumber+1)), (charSize+1)*1000)
	if err != nil {
		log.Fatal(err)
	}
	encrypted := encrypt(text, Keys[keyInd])
	keySize := getKeySize(encrypted, LGramLength)
	decryptedKey := decryptKey(encrypted, keySize)
	if decryptedKey == Keys[keyInd] {
		hits.Increment(keyInd*10 + charSize)
	}
}
