package day1

import (
	"bufio"
	"fmt"
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

func PuzzleOne() {
	f, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatalf("error opening puzzle input: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		var foundI, foundJ bool
		for i, j := 0, len(line)-1; i <= j; {
			if !foundI {
				_, err := strconv.Atoi(string(line[i]))
				if err != nil {
					i++
				} else {
					foundI = true
				}
			}
			if !foundJ {
				_, err := strconv.Atoi(string(line[j]))
				if err != nil {
					j--
				} else {
					foundJ = true
				}
			}
			if foundI && foundJ {
				num, err := strconv.Atoi(string(line[i]) + string(line[j]))
				if err != nil {
					log.Fatalf("error creating addend to sum %v", err)
				}
				sum += num
				break
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading input: %v", err)
	}
	fmt.Printf("total sum is %d", sum)
}

func OldPuzzleTwo() {
	f, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	defer f.Close()
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
	// log.Printf("final sum is %d", sum)
}

func reverseString(s string) string {
	runeSlice := []rune(s)
	for i, j := 0, len(runeSlice)-1; i < j; i, j = i+1, j-1 {
		runeSlice[i], runeSlice[j] = runeSlice[j], runeSlice[i]
	}
	return string(runeSlice)
}

func reverseBytes(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func PuzzleTwoReverseString() {
	f, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	defer f.Close()
	var sum int
	re := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|[1-9])`)
	backwardsRe := regexp.MustCompile(`(eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|[1-9])`)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ln := scanner.Text()
		frontNum := re.FindString(ln)
		if len(frontNum) == 0 {
			log.Printf("could not find digit in line: %v", ln)
			continue
		}
		reverseLn := reverseString(ln)
		backNumString := backwardsRe.FindString(reverseLn)
		backNum := reverseString(backNumString)
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
	// log.Printf("final sum is %d", sum)
}

func PuzzleTwoReverseBytes() {
	f, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	defer f.Close()
	var sum int
	re := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|[1-9])`)
	backwardsRe := regexp.MustCompile(`(eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|[1-9])`)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		b := scanner.Bytes()
		frontNumBytes := re.Find(b)
		frontNum := string(frontNumBytes)
		if len(frontNum) == 0 {
			log.Printf("could not find digit in line: %v", b)
			continue
		}
		reverseBytes(b)
		backNumBytes := backwardsRe.Find(b)
		reverseBytes(backNumBytes)
		backNum := string(backNumBytes)
		if len(frontNum) > 1 {
			if _, ok := stringToDigitMapping[frontNum]; ok {
				frontNum = stringToDigitMapping[frontNum]
			}
		}
		if len(backNum) > 1 {
			backNum = stringToDigitMapping[backNum]
			if _, ok := stringToDigitMapping[backNum]; ok {
				backNum = stringToDigitMapping[backNum]
			}
		}
		num, err := strconv.Atoi(frontNum + backNum)
		if err != nil {
			log.Fatalf("error converting nums in line from string to int: %v", err)
		}
		sum += num
	}
	// log.Printf("final sum is %d", sum)
}
