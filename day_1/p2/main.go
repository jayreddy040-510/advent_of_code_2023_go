package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

var stringToDigitMapping = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

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
		if len(frontNum) == 0 {
			log.Printf("could not find digit in line: %v", ln)
		}
		var backNum string
		var foundBackNum bool
		for i := len(ln) - 1; foundBackNum == false; i-- {
			backNum = re.FindString(ln[i:])
			if len(backNum) > 0 {
				break
			}
		}
		if len(frontNum) > 1 {
			frontNum = stringToDigitMapping[frontNum]
		}
		if len(backNum) > 1 {
			backNum = stringToDigitMapping[backNum]
		}
		num, err := strconv.Atoi(frontNum + backNum)
		if err != nil {
			log.Fatalf("error converting nums in line from string to int: %v", err)
		}
		sum += num
	}
	log.Printf("final sum is %d", sum)
}
