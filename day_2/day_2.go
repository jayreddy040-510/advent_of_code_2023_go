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
	if err != nil {
		log.Fatalf("could not read input file: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	re := regexp.MustCompile(`(\d+)\s{1}(\w+)`)
	var sum int
	for scanner.Scan() {
		var limitFlag bool
		ln := scanner.Text()
		colonSplitStrings := strings.Split(ln, ":")
		cubesSlice := re.FindAllStringSubmatch(colonSplitStrings[1], -1)
		// log.Printf("cubeSlice: %v", cubesSlice)
		for _, cube := range cubesSlice {
			num, err := strconv.Atoi(cube[1])
			if err != nil {
				log.Fatalf("erro converting str to num: %v", err)
			}
			if cube[2] == "red" && num > 12 {
				limitFlag = true
				break
			} else if cube[2] == "green" && num > 13 {
				limitFlag = true
				break
			} else if cube[2] == "blue" && num > 14 {
				limitFlag = true
				break
			}
		}
		if !limitFlag {
			gameStringsArray := strings.Fields(colonSplitStrings[0])
			gameNum, err := strconv.Atoi(gameStringsArray[1])
			if err != nil {
				log.Fatalf("error converting game str to num: %v", err)
			}
			sum += gameNum
		}
	}
	log.Printf("SUM SUM SUM SUM SUM SUM is %d", sum)
}

func PuzzleTwo() {
	f, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatalf("could not read input file: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	re := regexp.MustCompile(`(\d+)\s{1}(\w+)`)
	var sum int
	var maxRed, maxGreen, maxBlue int
	for scanner.Scan() {
		ln := scanner.Text()
		colonSplitStrings := strings.Split(ln, ":")
		cubesSlice := re.FindAllStringSubmatch(colonSplitStrings[1], -1)
		// log.Printf("cubeSlice: %v", cubesSlice)
		for _, cube := range cubesSlice {
			num, err := strconv.Atoi(cube[1])
			if err != nil {
				log.Fatalf("erro converting str to num: %v", err)
			}
			if cube[2] == "red" && num > maxRed {
				maxRed = num
			} else if cube[2] == "green" && num > maxGreen {
				maxGreen = num
			} else if cube[2] == "blue" && num > maxBlue {
				maxBlue = num
			}
		}
		sum += maxRed * maxGreen * maxBlue
		maxRed, maxGreen, maxBlue = 0, 0, 0
	}
	// log.Printf("SUM SUM SUM SUM SUM SUM is %d", sum)
}
