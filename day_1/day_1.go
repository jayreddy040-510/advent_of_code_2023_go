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
	defer f.Close()
	if err != nil {
		log.Fatalf("error opening puzzle input: %v", err)
	}
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
	// log.Printf("final sum is %d", sum)
}

func PuzzleTwo() {
	f, err := os.Open("puzzle_input.txt")
	if err != nil {
		log.Fatalf("cant open puzzle_input.txt: %v", err)
	}
	scanner := bufio.NewScanner(f)
	re := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|[1-9])`)
	var sum int
	for scanner.Scan() {
		b := scanner.Bytes()
		matches := re.FindAll(b, -1)
		frontNumBytes, backNumBytes := matches[0], matches[len(matches)-1]
		var frontNum, backNum string
		if len(frontNumBytes) > 1 {
			frontNum = stringToDigitMapping[string(frontNumBytes)]
		} else {
			frontNum = string(frontNumBytes)
		}
		if len(backNumBytes) > 1 {
			backNum = stringToDigitMapping[string(backNumBytes)]
		} else {
			backNum = string(backNumBytes)
		}
		//	log.Printf("ln: %v\nfrontNum: %v\nbackNum: %v", string(b), frontNum, backNum)
		num, err := strconv.Atoi(frontNum + backNum)
		if err != nil {
			log.Fatalf("error converting number in line to string: %v", err)
		}
		sum += num
	}
	log.Printf("sum is %d", sum)
}

func reverseString(s string) string {
	runeSlice := []rune(s)
	for i, j := 0, len(runeSlice)-1; i < j; i, j = i+1, j-1 {
		runeSlice[i], runeSlice[j] = runeSlice[j], runeSlice[i]
	}
	return string(runeSlice)
}

func reverseBytes(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}

func PuzzleTwoReverseString() {
	f, err := os.Open("puzzle_input.txt")
	defer f.Close()
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
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
	defer f.Close()
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
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
		rb := reverseBytes(b)
		backNumBytes := backwardsRe.Find(rb)
		backNumBytes = reverseBytes(backNumBytes)
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
