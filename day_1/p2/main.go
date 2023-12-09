package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

// var stringToDigitMapping
func main() {
	f, err := os.Open("../puzzle_input.txt")
	defer f.Close()
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	var sum int
	re := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|[1-9])`)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ln := scanner.Text()
		frontNum := re.FindString(ln)
		var backNum string
		var foundBackNum bool
		for i := len(ln) - 1; foundBackNum == false; i-- {
			backNum = re.FindString(ln[i:])
			if len(backNum) > 0 {
				break
			}
		}
		if len(frontNum) > 1 {
		}
		num, err := strconv.Atoi(frontNum + backNum)
		if err != nil {
			log.Fatalf("convert nums in line from string to int: %v", err)
		}
		sum += num
	}
}
