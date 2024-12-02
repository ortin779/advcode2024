package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isSafeReport(report []int) bool {
	asc, desc := 1, -1

	diff := report[1] - report[0]
	prevDirection := asc
	if diff < 0 {
		prevDirection = desc
	}

	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]
		direction := asc
		if diff < 0 {
			direction = desc
		}
		switch diff {
		case 1, 2, 3, -1, -2, -3:
			if prevDirection != direction {
				return false
			}
		default:
			return false
		}
		prevDirection = direction
	}

	return true
}

func isSafeReportWithSkip(report []int) bool {
	if isSafeReport(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
		newReport := []int{}
		newReport = append(newReport, report[:i]...)
		newReport = append(newReport, report[i+1:]...)
		if isSafeReport(newReport) {
			return true
		}
	}

	return false
}

func getSafeReports(reports [][]int) int {
	count := 0
	for _, v := range reports {
		if isSafeReportWithSkip(v) {
			count += 1
		}
	}
	return count
}

func main() {
	file, _ := os.Open("input.txt")
	bufReader := bufio.NewReader(file)

	matrix := [][]int{}
	testMatrix := [][]int{{7, 6, 4, 2, 1}, {1, 2, 7, 8, 9}, {9, 7, 6, 2, 1}, {1, 3, 2, 4, 5}, {8, 6, 4, 4, 1}}
	for {
		line, _, err := bufReader.ReadLine()
		if err != nil {
			break
		}
		lineWithoutNewLine := strings.Split(string(line), "\n")
		nums := strings.Split(lineWithoutNewLine[0], " ")
		row := []int{}
		for _, v := range nums {
			num, _ := strconv.Atoi(v)
			row = append(row, num)
		}
		matrix = append(matrix, row)
	}

	fmt.Printf("testMatrix count : %d\n", getSafeReports(testMatrix))
	fmt.Printf("matrix count : %d\n", getSafeReports(matrix))

}
