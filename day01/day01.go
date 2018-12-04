package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/storypixel/adventofcode-2018/aoclib"
)

func main() {
	var currFreq int // keep our running total
	var freqDelta int
	var freqs []int
	var changes []int

	file, err := os.Open("./testdata/fixture1")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	freqs = append(freqs, 0)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		freqDelta, err = strconv.Atoi(line)
		changes = append(changes, freqDelta)
		currFreq += freqDelta
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %v\n", currFreq)

	// Part 2
	currFreq = 0
	clen := len(changes)
	i := 0
	newFreq := 0
	freq := 0
	for {
		change := changes[i]
		freq = freqs[len(freqs)-1]
		newFreq = freq + change
		if aoclib.Contains(freqs, newFreq) {
			break
		}
		freqs = append(freqs, newFreq)
		i = (i + 1) % clen // circular index
	}

	fmt.Printf("Part 2: %v\n", newFreq)
}
