package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("day09/input-0")
	defer f.Close()

	var lines [][][]int
	s := bufio.NewScanner(f)
	for s.Scan() {
		var line []int
		for _, numStr := range strings.Split(s.Text(), " ") {
			num, _ := strconv.Atoi(numStr)
			line = append(line, num)
		}
		lines = append(lines, [][]int{line})
	}
	// Recursive calculate diffs
	for i := 0; i < len(lines); i++ {
		last := lines[i][len(lines[i])-1]
		if allZeroes(last) {
			continue
		} else {
			lines[i] = append(lines[i], diffs(last))
			i = -1
		}
	}

	// Part 1
	// Recursive extrapolate
	for i := 0; i < len(lines); i++ {
		for j := len(lines[i]) - 1; j >= 1; j-- {
			baseSeq := lines[i][j-1]
			base := baseSeq[len(baseSeq)-1]
			addSeq := lines[i][j]
			add := addSeq[len(addSeq)-1]
			ext := extrapolate(base, add)
			lines[i][j-1] = append(lines[i][j-1], ext)
		}
	}
	// Sum all extrapolations
	var sumPart1 int
	for _, line := range lines {
		sumPart1 += line[0][len(line[0])-1]
	}
	fmt.Println(sumPart1)

	// Part 2
	// Recursive extrapolate left
	for i := 0; i < len(lines); i++ {
		for j := len(lines[i]) - 1; j >= 1; j-- {
			baseSeq := lines[i][j-1]
			base := baseSeq[0]
			addSeq := lines[i][j]
			add := addSeq[0]
			ext := extrapolate(base, -add)
			lines[i][j-1] = append([]int{ext}, lines[i][j-1]...)
		}
	}
	// Sum all extrapolations
	var sumPart2 int
	for _, line := range lines {
		sumPart2 += line[0][0]
	}
	fmt.Println(sumPart2)
}

func extrapolate(base, add int) int {
	return base + add
}

func diffs(nums []int) []int {
	var diffs []int
	for i := 1; i < len(nums); i++ {
		diffs = append(diffs, nums[i]-nums[i-1])
	}
	return diffs
}

func allZeroes(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}
