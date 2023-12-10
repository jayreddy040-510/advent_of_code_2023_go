package day2

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func PuzzleOne() {
	f, err := os.Open("puzzle_input.txt")
	defer f.Close()
	if err != nil {
		log.Fatalf("could not read input file: %v", err)
	}
	scanner := bufio.NewScanner(f)
	re := regexp.MustCompile(`(\d+\s{1}red|\d+\s{1}blue|\d+\s{1}green)`)
	var sum int
	for scanner.Scan() {
		var limitFlag bool
		ln := scanner.Text()
		colonSplitStrings := strings.Split(ln, ":")
		cubesSlice := re.FindAllString(colonSplitStrings[1], -1)
		log.Printf("cubeSlice: %v", cubesSlice)
		for _, cube := range cubesSlice {
			eachCube := strings.Split(cube, " ")
			num, err := strconv.Atoi(eachCube[0])
			if err != nil {
				log.Fatalf("erro converting str to num: %v", err)
			}
			if eachCube[1] == "red" && num > 12 {
				limitFlag = true
				break
			} else if eachCube[1] == "green" && num > 13 {
				limitFlag = true
				break
			} else if eachCube[1] == "blue" && num > 14 {
				limitFlag = true
				break
			}
		}
		if !limitFlag {
			gameStringsArray := strings.Split(colonSplitStrings[0], " ")
			gameNum, err := strconv.Atoi(gameStringsArray[1])
			if err != nil {
				log.Fatalf("error converting game str to num: %v", err)
			}
			sum += gameNum
		}
	}
	log.Printf("SUM SUM SUM SUM SUM SUM is %d", sum)
}
