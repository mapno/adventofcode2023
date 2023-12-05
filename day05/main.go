package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type entry struct {
	destRangeStart int
	origRangeStart int
	rangNum        int
}

type rangeMappper []entry

func (m *rangeMappper) addEntry(e entry) {
	*m = append(*m, e)
}

func (m *rangeMappper) getDest(orig int) int {
	for _, e := range *m {
		if orig >= e.origRangeStart && orig < e.origRangeStart+e.rangNum {
			return e.destRangeStart + orig - e.origRangeStart
		}
	}
	return orig
}

func buildRangeMappper(s *bufio.Scanner) rangeMappper {
	var m rangeMappper

	for s.Scan() {
		if s.Text() == "" {
			break
		}
		parts := strings.Split(s.Text(), " ")
		destRangeStart, _ := strconv.Atoi(parts[0])
		origRangeStart, _ := strconv.Atoi(parts[1])
		rangNum, _ := strconv.Atoi(parts[2])
		m.addEntry(entry{destRangeStart, origRangeStart, rangNum})
	}

	return m
}

func main() {
	//part1()
	part2()
}

func part1() {
	f, _ := os.Open("day05/input-0")
	defer func() { _ = f.Close() }()

	var seeds []int

	s := bufio.NewScanner(f)
	s.Scan()
	for _, seedStr := range strings.Split(s.Text()[len("seeds: "):], " ") {
		seed, _ := strconv.Atoi(seedStr)
		seeds = append(seeds, seed)
	}

	m1, m2, m3, m4, m5, m6, m7 := buildMappers(s)

	lowestLoc := math.MaxInt64
	for _, seed := range seeds {
		loc := locForSeed(m1, m2, m3, m4, m5, m6, m7, seed)
		if loc < lowestLoc {
			lowestLoc = loc
		}
	}

	fmt.Println(lowestLoc)
}

func buildMappers(s *bufio.Scanner) (rangeMappper, rangeMappper, rangeMappper, rangeMappper, rangeMappper, rangeMappper, rangeMappper) {
	s.Scan() // Empty line
	s.Scan() // `seed-to-soil map:` line
	m1 := buildRangeMappper(s)

	s.Scan() // `soil-to-fertilizer map:` line
	m2 := buildRangeMappper(s)

	s.Scan() // `fertilizer-to-water map:` line
	m3 := buildRangeMappper(s)

	s.Scan() // `water-to-light map:` line
	m4 := buildRangeMappper(s)

	s.Scan() // `light-to-temperature map:` line
	m5 := buildRangeMappper(s)

	s.Scan() // `temperature-to-humidity map:` line
	m6 := buildRangeMappper(s)

	s.Scan() // `humidity-to-location map:` line
	m7 := buildRangeMappper(s)

	return m1, m2, m3, m4, m5, m6, m7
}

func locForSeed(m1, m2, m3, m4, m5, m6, m7 rangeMappper, seed int) int {
	d2 := m1.getDest(seed)
	d3 := m2.getDest(d2)
	d4 := m3.getDest(d3)
	d5 := m4.getDest(d4)
	d6 := m5.getDest(d5)
	d7 := m6.getDest(d6)
	return m7.getDest(d7)
}

func part2() {
	f, _ := os.Open("day05/input-0")
	defer func() { _ = f.Close() }()

	s := bufio.NewScanner(f)
	s.Scan()

	var seeds []int
	seedStrs := strings.Split(s.Text()[len("seeds: "):], " ")
	for i := 0; i < len(seedStrs); i += 2 {
		seed, _ := strconv.Atoi(seedStrs[i])
		seedRange, _ := strconv.Atoi(seedStrs[i+1])
		for j := 0; j < seedRange; j++ {
			seeds = append(seeds, seed+j)
		}
	}

	m1, m2, m3, m4, m5, m6, m7 := buildMappers(s)

	//seeds = dedupSlice(seeds)
	fmt.Println("Total seeds:", len(seeds))
	lowestLoc := math.MaxInt64
	for i, seed := range seeds {
		loc := locForSeed(m1, m2, m3, m4, m5, m6, m7, seed)
		if loc < lowestLoc {
			lowestLoc = loc
		}

		if i%1_000_000 == 0 {
			fmt.Println("Processed", i, "seeds")
		}
	}

	fmt.Println(lowestLoc)
}

//func dedupSlice(s []int) []int {
//	m := make(map[int]bool)
//	var res []int
//	for _, v := range s {
//		if _, ok := m[v]; !ok {
//			m[v] = true
//			res = append(res, v)
//		}
//	}
//	return res
//}
