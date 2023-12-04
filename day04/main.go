package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	f, _ := os.Open("day04/input-0")
	defer func() { _ = f.Close() }()

	var sum int
	s := bufio.NewScanner(f)
	for s.Scan() {
		nums := strings.Split(strings.Split(s.Text(), ":")[1], "|")
		winningNums := make(map[int]bool)
		numsIHave := make(map[int]bool)
		for _, num := range strings.Split(strings.TrimSpace(nums[0]), " ") {
			if num == "" {
				continue
			}
			n, _ := strconv.Atoi(num)
			winningNums[n] = true
		}
		for _, num := range strings.Split(strings.TrimSpace(nums[1]), " ") {
			if num == "" {
				continue
			}
			n, _ := strconv.Atoi(num)
			numsIHave[n] = true
		}
		var count int
		for num := range numsIHave {
			if _, ok := winningNums[num]; ok {
				count++
			}
		}
		sum += powInt(count - 1)
	}
	fmt.Println(sum)
}

func powInt(y int) int {
	return int(math.Pow(float64(2), float64(y)))
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func part2() {
	f, _ := os.Open("day04/input-0")
	defer func() { _ = f.Close() }()

	var sum int
	s := bufio.NewScanner(f)
	lines := make([]string, 0)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	queue := &IntHeap{}
	heap.Init(queue)
	for i := range lines {
		heap.Push(queue, i)
	}
	countCache := make(map[int]int)

	for queue.Len() > 0 {
		sum++

		// Pop
		curr := heap.Pop(queue).(int)

		if count, ok := countCache[curr]; ok {
			for i := curr + 1; i <= curr+count; i++ {
				heap.Push(queue, i)
			}
			continue
		}

		nums := strings.Split(strings.Split(lines[curr], ":")[1], "|")
		winningNums := make(map[int]bool)
		numsIHave := make(map[int]bool)
		for _, num := range strings.Split(strings.TrimSpace(nums[0]), " ") {
			if num == "" {
				continue
			}
			n, _ := strconv.Atoi(num)
			winningNums[n] = true
		}
		for _, num := range strings.Split(strings.TrimSpace(nums[1]), " ") {
			if num == "" {
				continue
			}
			n, _ := strconv.Atoi(num)
			numsIHave[n] = true
		}

		var count int
		for num := range numsIHave {
			if _, ok := winningNums[num]; ok {
				count++
			}
		}
		countCache[curr] = count

		for i := curr + 1; i <= curr+count; i++ {
			heap.Push(queue, i)
		}
	}

	fmt.Println(sum)
}
