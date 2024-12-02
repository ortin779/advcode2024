package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getTotalDistance(ls []int, rs []int) int {
	slices.Sort(ls)
	slices.Sort(rs)

	sum := 0
	for idx := range ls {
		sum += int(math.Abs(float64(ls[idx] - rs[idx])))
	}
	return sum
}

func getSimilarityScore(ls []int, rs []int) int {
	countMap := make(map[int]int, len(rs))
	for _, v := range rs {
		countMap[v] = countMap[v] + 1
	}
	simScore := 0
	for _, v := range ls {
		simScore += countMap[v] * v
	}
	return simScore
}

func main() {
	file, _ := os.Open("inp.txt")
	bufReader := bufio.NewReader(file)

	ls, rs := []int{}, []int{}
	for {
		line, _, err := bufReader.ReadLine()
		if err != nil {
			break
		}
		lineWithoutNewLine := strings.Split(string(line), "\n")
		words := strings.Split(lineWithoutNewLine[0], "   ")
		n1, err := strconv.Atoi(words[0])
		if err != nil {
			log.Fatalln("err", err)
		}
		ls = append(ls, n1)

		n2, err := strconv.Atoi(words[1])
		if err != nil {
			log.Fatalln("err", err)
		}
		rs = append(rs, n2)
	}

	fmt.Println("totalDistance", getTotalDistance(ls, rs))
	fmt.Println("similarityScore", getSimilarityScore(ls, rs))

}
