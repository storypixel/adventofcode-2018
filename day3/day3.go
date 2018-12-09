package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var rawdata []string

	file, err := os.Open("./testdata/fixture2")
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
	type Claim struct {
		id     int
		xpos   int
		ypos   int
		width  int
		height int
		pure   bool
	}

	const dimX int = 1000
	const dimY int = 1000 // these could be dynamic
	cloth := [dimX][dimY][]int{}
	instructions := make([]Claim, len(rawdata))
	instructions[0] = Claim{id: 0, xpos: 1, ypos: 2, width: 3, height: 4}
	// fmt.Println(cloth, instructions)

	for i, line := range rawdata {
		// fmt.Println("line: ", line)
		// Claim Num
		r, _ := regexp.Compile(`\#\d+`)
		match := r.FindString(line)
		claimNum, _ := strconv.Atoi(match[1:])
		// fmt.Println(claimNum)

		// X and Y pos
		r, _ = regexp.Compile(`\d+,\d+`)
		match = r.FindString(line)
		poses := strings.Split(match, ",")
		posX, _ := strconv.Atoi(poses[0])
		posY, _ := strconv.Atoi(poses[1])
		// fmt.Println(posX, posY)

		// Width and Height
		r, _ = regexp.Compile(`\d+x\d+`)
		match = r.FindString(line)
		dims := strings.Split(match, "x")
		width, _ := strconv.Atoi(dims[0])
		height, _ := strconv.Atoi(dims[1])
		// fmt.Println(width, height)

		instructions[i] = Claim{
			id:     claimNum,
			xpos:   posX,
			ypos:   posY,
			width:  width,
			height: height,
			pure:   true}
		// fmt.Println(instructions[i])

		xindex := posX
		yindex := posY
		stopx := xindex + width
		stopy := yindex + height

		for xindex < stopx {
			for yindex < stopy {
				var dog = claimNum //[dimX][dimY][]int{}
				cloth[xindex][yindex] = append(cloth[xindex][yindex], dog)
				yindex++
			}
			yindex = posY
			xindex++
		}
	}

	totalOverlaps := 0
	for _, row := range cloth {
		// fmt.Println(row)
		for _, line := range row {
			if len(line) > 1 {
				totalOverlaps++
			}
		}
	}
	fmt.Printf("Part 1: %v", totalOverlaps)

	// Part 2

	// almost same thing in part 1 but take different conclusions
	for _, row := range cloth {
		fmt.Println(row)
		for _, line := range row {
			if len(line) > 1 {
				for _, cid := range line {
					instructionsIndex := cid - 1
					instructions[instructionsIndex].pure = false
					fmt.Println(instructions[instructionsIndex])
				}
			}
		}
	}

	for _, istr := range instructions {
		if istr.pure {
			fmt.Printf("%v is pure", istr.id)
		}
	}
}
