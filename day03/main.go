package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//part1()
	part2()
}

const (
	dot rune = -2
	ast rune = -6
)

func part1() {
	f, _ := os.Open("day03/input-0")
	defer func() { _ = f.Close() }()

	var grid [][]rune
	s := bufio.NewScanner(f)
	for s.Scan() {
		grid = append(grid, []rune{})
		for _, c := range s.Text() {
			grid[len(grid)-1] = append(grid[len(grid)-1], c-'0')
		}
	}

	// Traverse the grid
	var isNum bool
	var num, sum int
	var numxy [][2]int
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] >= 0 && grid[y][x] <= 9 {
				if isNum {
					num = num*10 + int(grid[y][x])
				} else {
					num = int(grid[y][x])
				}
				numxy = append(numxy, [2]int{x, y})
				isNum = true
			} else {
				if num > 0 {
					fmt.Println(num, numxy)
					for _, xy := range numxy {
						if has := hasAdjacentSymbol(grid, xy[0], xy[1]); has {
							fmt.Println("has adjacent symbol")
							sum += num
							break
						}
					}
				}
				isNum = false
				numxy = nil
				num = 0
			}
		}
	}

	fmt.Println(sum)
}

func hasAdjacentSymbol(grid [][]rune, x, y int) bool {
	// Search in all directions
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dy == 0 && dx == 0 {
				continue
			}

			if y+dy >= 0 && y+dy < len(grid) &&
				x+dx >= 0 && x+dx < len(grid[y+dy]) &&
				(grid[y+dy][x+dx] < 0 || grid[y+dy][x+dx] > 9) && grid[y+dy][x+dx] != dot {
				return true
			}
		}
	}

	return false
}

func part2() {
	f, _ := os.Open("day03/input-0")
	defer func() { _ = f.Close() }()

	var grid [][]rune
	s := bufio.NewScanner(f)
	for s.Scan() {
		grid = append(grid, []rune{})
		for _, c := range s.Text() {
			grid[len(grid)-1] = append(grid[len(grid)-1], c-'0')
		}
	}

	var isNum bool
	var num int
	var numxy [][2]int
	numMap := make(map[[2]int]int)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] >= 0 && grid[y][x] <= 9 {
				if isNum {
					num = num*10 + int(grid[y][x])
				} else {
					num = int(grid[y][x])
				}
				numxy = append(numxy, [2]int{x, y})
				isNum = true
			} else {
				if num > 0 {
					for _, xy := range numxy {
						numMap[xy] = num
					}
				}
				isNum = false
				numxy = nil
				num = 0
			}
		}
	}

	var sum int
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == ast {
				nums := dfs(grid, numMap, x, y)
				if len(nums) == 2 {
					sum += nums[0] * nums[1]
				}
			}
		}
	}
	fmt.Println(sum)

}

func dfs(grid [][]rune, numMap map[[2]int]int, x, y int) []int {
	// Search in all directions
	var nums []int
	for dy := -1; dy <= 1; dy++ {
		var rowNum int
		for dx := -1; dx <= 1; dx++ {
			if dy == 0 && dx == 0 {
				continue
			}

			if grid[y+dy][x+dx] < 0 || grid[y+dy][x+dx] > 9 {
				rowNum = -1
				continue
			}

			num, ok := numMap[[2]int{x + dx, y + dy}]
			if ok && num != rowNum {
				nums = append(nums, num)
			}
			rowNum = num
		}
	}
	if len(nums) > 1 {
		return dedup(nums)
	}
	return nil
}

func dedup(s []int) []int {
	// Check if all are the same
	same := true
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			same = false
			break
		}
	}
	if same {
		return s[:2]
	}

	var m = make(map[int]bool)
	var r []int
	for _, v := range s {
		if _, ok := m[v]; !ok {
			m[v] = true
			r = append(r, v)
		}
	}
	return r
}
