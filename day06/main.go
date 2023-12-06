package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	bruteForce("day06/input-0") // Part 1
	bruteForce("day06/input-1") // Part 2

	maths("day06/input-0") // Part 1
	maths("day06/input-1") // Part 2
}

func maths(fileName string) {
	f, _ := os.Open(fileName)
	defer func() { _ = f.Close() }()

	var times, distances []int
	s := bufio.NewScanner(f)
	s.Scan()
	timesStrs := strings.Split(s.Text()[len(`Time:      `):], " ")
	for _, timeStr := range timesStrs {
		if timeStr == "" {
			continue
		}
		time, _ := strconv.Atoi(strings.TrimSpace(timeStr))
		times = append(times, time)
	}
	s.Scan()
	distancesStrs := strings.Split(s.Text()[len(`Distance:  `):], " ")
	for _, distanceStr := range distancesStrs {
		if distanceStr == "" {
			continue
		}
		distance, _ := strconv.Atoi(strings.TrimSpace(distanceStr))
		distances = append(distances, distance)
	}

	total := 1
	for i := 0; i < len(times); i++ {
		b1, b2 := bounds(times[i], distances[i])
		total *= b2 - b1 + 1
	}
	fmt.Println(total)

}

func bounds(t, d int) (int, int) {
	sqrt := math.Sqrt(float64(t*t - 4*d))
	floatT := float64(t)

	b1 := (floatT - sqrt) / 2.0
	b2 := (floatT + sqrt) / 2.0

	return int(math.Floor(b1 + 1)), int(math.Ceil(b2 - 1))
}

func bruteForce(fileName string) {
	f, _ := os.Open(fileName)
	defer func() { _ = f.Close() }()

	var times, distances []int
	s := bufio.NewScanner(f)
	s.Scan()
	timesStrs := strings.Split(s.Text()[len(`Time:      `):], " ")
	for _, timeStr := range timesStrs {
		if timeStr == "" {
			continue
		}
		time, _ := strconv.Atoi(strings.TrimSpace(timeStr))
		times = append(times, time)
	}
	s.Scan()
	distancesStrs := strings.Split(s.Text()[len(`Distance:  `):], " ")
	for _, distanceStr := range distancesStrs {
		if distanceStr == "" {
			continue
		}
		distance, _ := strconv.Atoi(strings.TrimSpace(distanceStr))
		distances = append(distances, distance)
	}

	total := 1
	for i := 0; i < len(times); i++ {
		var count int
		for j := 1; j < times[i]; j++ {
			dist := j * (times[i] - j)
			if dist > distances[i] {
				count++
			}
		}
		if count > 0 {
			total *= count
		}
	}
	fmt.Println(total)
}

func part2() {
	f, _ := os.Open("day06/input-1")
	defer func() { _ = f.Close() }()

	var times, distances []int
	s := bufio.NewScanner(f)
	s.Scan()
	timesStrs := strings.Split(s.Text()[len(`Time:      `):], " ")
	for _, timeStr := range timesStrs {
		if timeStr == "" {
			continue
		}
		time, _ := strconv.Atoi(strings.TrimSpace(timeStr))
		times = append(times, time)
	}
	s.Scan()
	distancesStrs := strings.Split(s.Text()[len(`Distance:  `):], " ")
	for _, distanceStr := range distancesStrs {
		if distanceStr == "" {
			continue
		}
		distance, _ := strconv.Atoi(strings.TrimSpace(distanceStr))
		distances = append(distances, distance)
	}

	total := 1
	for i := 0; i < len(times); i++ {
		var count int
		for j := 1; j < times[i]; j++ {
			dist := j * (times[i] - j)
			if dist > distances[i] {
				count++
			}
		}
		if count > 0 {
			total *= count
		}
	}
	fmt.Println(total)
}
