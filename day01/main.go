package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	f, _ := os.Open("day01/input-0")

	defer func() { _ = f.Close() }()

	var total int
	s := bufio.NewScanner(f)
	for s.Scan() {
		f, l := -1, -1
		for _, c := range s.Text() {
			if num := int(c - '0'); num < 10 {
				if f == -1 {
					f = num
				}
				l = num
			}
		}
		total += f*10 + l
	}
	fmt.Println(total)
}

func part2() {
	f, _ := os.Open("day01/input-0")

	defer func() { _ = f.Close() }()

	numMap := map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5,
		"six": 6, "seven": 7, "eight": 8, "nine": 9,
		"1": 1, "2": 2, "3": 3, "4": 4, "5": 5,
		"6": 6, "7": 7, "8": 8, "9": 9,
	}

	var total int
	s := bufio.NewScanner(f)
	for s.Scan() {
		idxs := make([]int, len(s.Text()))
		l := s.Text()
		for i := 0; i < len(l); i++ {
			ll := l[i:]
			for nStr, n := range numMap {
				if idx := strings.Index(ll, nStr); idx != -1 {
					idxs[idx+i] = n
				}
			}
		}
		first, last := -1, -1
		for _, idx := range idxs {
			if idx != 0 {
				if first == -1 {
					first = idx
				}
				last = idx
			}
		}

		total += first*10 + last
	}
	fmt.Println(total)
}
