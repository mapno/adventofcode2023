package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//part1()
	part2()
}

func part1() {
	f, _ := os.Open("day02/input-0")
	defer func() { _ = f.Close() }()

	cubeSet := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	var total int
	s := bufio.NewScanner(f)
	for s.Scan() {
		var game int
		_, _ = fmt.Sscanf(s.Text(), `Game %d`, &game)
		sets := strings.Split(strings.Split(s.Text(), ":")[1], ";")
		for _, set := range sets {
			for _, cube := range strings.Split(set, ",") {
				cube = strings.TrimSpace(cube)
				cubeSplit := strings.Split(cube, " ")
				num, _ := strconv.Atoi(cubeSplit[0])
				if cubeSet[cubeSplit[1]] < num {
					goto NEXT
				}
			}
		}
		total += game
	NEXT:
	}
	fmt.Println(total)
}

func part2() {
	f, _ := os.Open("day02/input-0")
	defer func() { _ = f.Close() }()

	var total int
	s := bufio.NewScanner(f)
	for s.Scan() {
		minCubes := map[string]int{"red": 0, "green": 0, "blue": 0}
		sets := strings.Split(strings.Split(s.Text(), ":")[1], ";")
		for _, set := range sets {
			for _, cube := range strings.Split(set, ",") {
				cube = strings.TrimSpace(cube)
				cubeSplit := strings.Split(cube, " ")
				num, _ := strconv.Atoi(cubeSplit[0])
				if num > minCubes[cubeSplit[1]] {
					minCubes[cubeSplit[1]] = num
				}
			}
		}
		count := 1
		for _, v := range minCubes {
			count *= v
		}
		total += count
	}
	fmt.Println(total)
}
