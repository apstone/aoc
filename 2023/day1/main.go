package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := bytes.Split(input, []byte("\n"))

	// Create a regex for finding numbers
	re, err := regexp.Compile("[0-9]")
	if err != nil {
		panic(err)
	}

	totalValue := 0

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if len(line) > 0 {
			numMatches := re.FindAllString(string(line), -1)
			firstDigit := numMatches[0]
			lastDigit := numMatches[len(numMatches)-1]
			mergedDigits := firstDigit + lastDigit

			number, err := strconv.Atoi(mergedDigits)
			if err != nil {
				panic(err)
			}
			totalValue += number
		}

	}
	fmt.Println(totalValue)
}
