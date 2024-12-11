package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseMultiplication(str string) int {
	parts := strings.Split(str[4:len(str)-1], ",")
	num1, _ := strconv.Atoi(parts[0])
	num2, _ := strconv.Atoi(parts[1])

	return num1 * num2
}

func sumOfMultiplications(lines []string) int {
	r := regexp.MustCompile(`mul\([\d]{1,3},[\d]{1,3}\)`)
	sum := 0
	for _, line := range lines {
		matches := r.FindAllString(line, -1)
		for _, match := range matches {
			sum += parseMultiplication(match)
		}
	}
	return sum
}

func sumOfMultiplicationsWithExclusion(lines []string) int {
	r := regexp.MustCompile(`(mul\([\d]{1,3},[\d]{1,3}\))|(don\'t)|(do)`)
	sum := 0
	enabled := true
	for _, line := range lines {
		matches := r.FindAllString(line, -1)
		for _, match := range matches {
			if strings.Contains(match, "don't") {
				enabled = false
			} else if strings.Contains(match, "do") {
				enabled = true
			} else {
				if enabled {
					sum += parseMultiplication(match)
				}
			}
		}
	}
	return sum
}

func main() {
	file, _ := os.Open("input.txt")
	bufReader := bufio.NewReader(file)

	lines := []string{}
	for {
		line, _, err := bufReader.ReadLine()
		if err != nil {
			break
		}
		lines = append(lines, string(line))
	}

	fmt.Printf("sum : %d\n", sumOfMultiplications(lines))
	fmt.Printf("Test sum : %d\n", sumOfMultiplicationsWithExclusion(lines))

}
