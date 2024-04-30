package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	findCount("input.txt")
}

func findCount(filename string) {
	bytes, _ := os.ReadFile(filename)
	fileContent := string(bytes)

	replacements := [][]string{
		{"one", "1"},
		{"two", "2"},
		{"three", "3"},
		{"four", "4"},
		{"five", "5"},
		{"six", "6"},
		{"seven", "7"},
		{"eight", "8"},
		{"nine", "9"},
	}

	lines := strings.Split(fileContent, "\n")
	re := regexp.MustCompile("[0-9]")

	sum := 0
	for _, line := range lines {
		words := getWords(replacements, line)

		digits := re.FindAllString(line, -1)
		if len(digits) <= 0 {
			continue
		}

		first := digits[0]
		last := digits[len(digits)-1]

		if len(words) > 0 {
			firstWord := replacements[words[0]]
			if strings.Index(line, firstWord[0]) < strings.Index(line, first) {
				first = firstWord[1]
			}

			lastWord := replacements[words[len(words)-1]]
			if strings.LastIndex(line, lastWord[0]) > strings.LastIndex(line, last) {
				last = lastWord[1]
			}
		}

		crc := fmt.Sprintf("%s%s", first, last)
		num, err := strconv.Atoi(crc)
		if err != nil {
			continue
		}
		sum += num
	}

	fmt.Printf("Sum is %d", sum)
}

func getWords(replacements [][]string, line string) []int {
	words := []int{}
	for i := 0; i < len(line); i++ {
		for id, replacement := range replacements {
			if strings.HasPrefix(line[i:], replacement[0]) {
				words = append(words, id)
			}
		}
	}

	return words
}
