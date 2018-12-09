package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Does two lines 'match' like this where abxde 'matches' abzde per puzzle
// where a certain letter in common position differs
func matchesWellAlmost(line string, line2 string) int {
	indexOfNonMatch := -1
	matchCount := 0

	for i := range line {
		// fmt.Printf("%v: \n", line[i] == line2[i])
		isRuneMatch := line[i] == line2[i]
		if isRuneMatch {
			matchCount++
		} else {
			indexOfNonMatch = i
		}
	}

	// Do we 'almost' match?
	if matchCount == (len(line) - 1) {
		return indexOfNonMatch
	}

	return -1
}

func main() {
	var rawdata []string

	file, err := os.Open("./testdata/fixture1")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		rawdata = append(rawdata, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Part 1
	twos := 0
	threes := 0

	for _, line := range rawdata {
		var freqs = make(map[string]int)
		for _, letter := range line {
			key := string(letter)
			freqs[key]++
			// fmt.Printf("%d: %c\n", i, rune)
		}
		//fmt.Printf("line map: %v\n", freqs)
		has2 := false
		has3 := false
		for _, freq := range freqs {
			// fmt.Printf("%d ", freq)
			if freq == 2 {
				has2 = true
			}

			if freq == 3 {
				has3 = true
			}
		}

		if has2 {
			twos++
		}

		if has3 {
			threes++
		}

		//fmt.Printf("\n")
	}

	// Part 2
	targetLine := ""
	sharedCharacterPosition := -1

	for i, line := range rawdata {
		// fmt.Printf("%v\n", line)
		starti := (i + 1) % len(line)
		matchesCriteria := false
		for _, line2 := range rawdata[starti:] {
			indexOfNonMatch := matchesWellAlmost(line, line2)
			matchesCriteria = indexOfNonMatch != -1
			// fmt.Printf("  %v vs %v is %v\n", line, line2, matchesWellAlmost(line, line2))
			if matchesCriteria {
				targetLine = line
				sharedCharacterPosition = indexOfNonMatch
				// fmt.Printf("  targetLine is %v, %v: %v\n", targetLine, sharedCharacterPosition, targetLine[:sharedCharacterPosition]+targetLine[sharedCharacterPosition+1:])
			}
			// matchesWellAlmost(line, line2)
			// targetLine = ""
			// sharedCharacterPosition = -1
		}

		//var freqs = make(map[string]int)
		// for i, rune := range line {
		// 	fmt.Printf("%d: %c\n", i, rune)
		// }
	}

	fmt.Printf("Part 1: %v %v %v\n", twos, threes, twos*threes)

	fmt.Printf("Part 2: targetLine is %v, %v: %v\n", targetLine, sharedCharacterPosition, targetLine[:sharedCharacterPosition]+targetLine[sharedCharacterPosition+1:])
}
