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
	f, _ := os.Open("day08/input-0")
	defer f.Close()

	nodes := make(map[string][]string)
	s := bufio.NewScanner(f)
	s.Scan()
	order := s.Text()
	s.Scan()
	for s.Scan() {
		ss := strings.Split(s.Text(), " = ")
		curr := ss[0]
		ss = strings.Split(ss[1], ", ")
		left, right := ss[0], ss[1]
		left = strings.TrimFunc(left, func(r rune) bool { return r == '(' || r == ')' })
		right = strings.TrimFunc(right, func(r rune) bool { return r == '(' || r == ')' })
		nodes[curr] = []string{left, right}
	}

	const dest = "ZZZ"
	curr := "AAA"
	var idx, count int
	for curr != dest {
		switch order[idx] {
		case 'L':
			curr = nodes[curr][0]
		case 'R':
			curr = nodes[curr][1]
		}
		idx++
		idx %= len(order)
		count++
	}
	fmt.Println(count)
}

func part2() {
	f, _ := os.Open("day08/input-0")
	defer f.Close()

	var currNodes []string
	nodes := make(map[string][]string)
	s := bufio.NewScanner(f)
	s.Scan()
	order := s.Text()
	s.Scan()
	for s.Scan() {
		ss := strings.Split(s.Text(), " = ")
		curr := ss[0]
		if curr[2] == 'A' {
			currNodes = append(currNodes, curr)
		}
		ss = strings.Split(ss[1], ", ")
		left, right := ss[0], ss[1]
		left = strings.TrimFunc(left, func(r rune) bool { return r == '(' || r == ')' })
		right = strings.TrimFunc(right, func(r rune) bool { return r == '(' || r == ')' })
		nodes[curr] = []string{left, right}
	}

	var counts []int
	var idx, count, orderIdx int
	curr := currNodes[idx]
	for idx < len(currNodes) {
		if isEndNode(curr) {
			counts = append(counts, count)
			idx++
			if idx < len(currNodes) {
				curr = currNodes[idx]
			}
			orderIdx = 0
			count = 0
		}
		switch order[orderIdx] {
		case 'L':
			curr = nodes[curr][0]
		case 'R':
			curr = nodes[curr][1]
		}
		orderIdx++
		orderIdx %= len(order)
		count++
	}
	fmt.Println(LCM(counts[0], counts[1], counts[2:]...))
}

func isEndNode(node string) bool {
	return node[2] == 'Z'
}

// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
