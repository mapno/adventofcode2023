package main

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

type typ int

func (t typ) String() string {
	switch t {
	case HighCard:
		return "HighCard"
	case OnePair:
		return "OnePair"
	case TwoPairs:
		return "TwoPairs"
	case ThreeOfAKind:
		return "ThreeOfAKind"
	case FullHouse:
		return "FullHouse"
	case FourOfAKind:
		return "FourOfAKind"
	case FiveOfAKind:
		return "FiveOfAKind"
	}
	return ""
}

const (
	HighCard typ = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var ranksPart1 = map[rune]int{
	'A': 14, 'K': 13, 'Q': 12, 'J': 11, 'T': 10, '9': 9,
	'8': 8, '7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2,
}

var ranksPart2 = map[rune]int{
	'A': 14, 'K': 13, 'Q': 12, 'J': 1, 'T': 10, '9': 9,
	'8': 8, '7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2,
}

type hand struct {
	cards string
	bid   int
}

func (h *hand) typ1() typ {
	m := map[rune]int{}
	for _, c := range h.cards {
		m[c]++
	}
	switch len(m) {
	case 1:
		return FiveOfAKind
	case 2:
		var a, b int
		for _, v := range m {
			if a == 0 {
				a = v
			} else {
				b = v
			}
		}
		if a == 1 || b == 1 {
			return FourOfAKind
		} else {
			return FullHouse
		}
	case 3:
		var a, b, c int
		for _, v := range m {
			if a == 0 {
				a = v
			} else if b == 0 {
				b = v
			} else {
				c = v
			}
		}
		// TwoPairs
		if (a == 1 && b == 2 && c == 2) || (a == 2 && b == 1 && c == 2) || (a == 2 && b == 2 && c == 1) {
			return TwoPairs
		} else if a == 3 || b == 3 || c == 3 {
			return ThreeOfAKind
		} else if a == 2 || b == 2 || c == 2 {
			return TwoPairs
		} else {
			return HighCard
		}
	case 4:
		var a, b, c, d int
		for _, v := range m {
			if a == 0 {
				a = v
			} else if b == 0 {
				b = v
			} else if c == 0 {
				c = v
			} else {
				d = v
			}
		}
		if a == 2 || b == 2 || c == 2 || d == 2 {
			return OnePair
		} else {
			return HighCard
		}
	case 5:
		return HighCard
	}
	return HighCard
}

func (h *hand) typ2() typ {
	m := map[rune]int{}
	var jokers int
	for _, c := range h.cards {
		if c == 'J' {
			jokers++
		} else {
			m[c]++
		}
	}
	if len(m) == 0 {
		return FiveOfAKind
	}
	for i := 0; i < jokers; i++ {
		for k := range m {
			m[k]++
		}
	}
	switch len(m) {
	case 1:
		return FiveOfAKind
	case 2:
		var a, b int
		for _, v := range m {
			if a == 0 {
				a = v
			} else {
				b = v
			}
		}
		if a == 4 || b == 4 {
			return FourOfAKind
		} else {
			return FullHouse
		}
	case 3:
		var a, b, c int
		for _, v := range m {
			if a == 0 {
				a = v
			} else if b == 0 {
				b = v
			} else {
				c = v
			}
		}
		// TwoPairs
		if (a == 1 && b == 2 && c == 2) || (a == 2 && b == 1 && c == 2) || (a == 2 && b == 2 && c == 1) {
			return TwoPairs
		} else if a == 3 || b == 3 || c == 3 {
			return ThreeOfAKind
		} else if a == 2 || b == 2 || c == 2 {
			return TwoPairs
		} else {
			return HighCard
		}
	case 4:
		var a, b, c, d int
		for _, v := range m {
			if a == 0 {
				a = v
			} else if b == 0 {
				b = v
			} else if c == 0 {
				c = v
			} else {
				d = v
			}
		}
		if a == 2 || b == 2 || c == 2 || d == 2 {
			return OnePair
		} else {
			return HighCard
		}
	case 5:
		return HighCard
	}
	return HighCard
}

func (h *hand) compare1(other hand, r map[rune]int) int {
	typ1 := h.typ1()
	typ2 := other.typ1()
	if typ1 > typ2 {
		return 1
	} else if typ1 < typ2 {
		return -1
	} else {
		for i := 0; i < len(h.cards); i++ {
			if r[rune(h.cards[i])] > r[rune(other.cards[i])] {
				return 1
			} else if r[rune(h.cards[i])] < r[rune(other.cards[i])] {
				return -1
			}
		}
	}
	return 0
}

func (h *hand) compare2(other hand, r map[rune]int) int {
	typ1 := h.typ2()
	typ2 := other.typ2()
	if typ1 > typ2 {
		return 1
	} else if typ1 < typ2 {
		return -1
	} else {
		for i := 0; i < len(h.cards); i++ {
			if r[rune(h.cards[i])] > r[rune(other.cards[i])] {
				return 1
			} else if r[rune(h.cards[i])] < r[rune(other.cards[i])] {
				return -1
			}
		}
	}
	return 0
}

var _ heap.Interface = (*handQueue)(nil)

type handQueue struct {
	part  int
	hands []hand
	r     map[rune]int
}

func (h handQueue) Len() int {
	return len(h.hands)
}

func (h handQueue) Less(i, j int) bool {
	if h.part == 1 {
		return h.hands[i].compare1(h.hands[j], h.r) == -1
	}
	return h.hands[i].compare2(h.hands[j], h.r) == -1
}

func (h handQueue) Swap(i, j int) { h.hands[i], h.hands[j] = h.hands[j], h.hands[i] }

func (h *handQueue) Push(x any) { h.hands = append(h.hands, x.(hand)) }

func (h *handQueue) Pop() any {
	old := h.hands
	n := len(old)
	x := old[n-1]
	h.hands = old[0 : n-1]
	return x
}

func part1() {
	f, _ := os.Open("day07/input-0")
	defer func() { _ = f.Close() }()

	h := &handQueue{r: ranksPart1, part: 1}
	heap.Init(h)

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.Split(s.Text(), " ")
		bid, _ := strconv.Atoi(line[1])
		heap.Push(h, hand{cards: line[0], bid: bid})
	}
	var i, sum int
	for h.Len() > 0 {
		bid := heap.Pop(h).(hand)
		i++
		sum += bid.bid * i
	}
	println(sum)
}

func part2() {
	f, _ := os.Open("day07/input-0")
	defer func() { _ = f.Close() }()

	h := &handQueue{r: ranksPart2, part: 2}
	heap.Init(h)

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.Split(s.Text(), " ")
		bid, _ := strconv.Atoi(line[1])
		heap.Push(h, hand{cards: line[0], bid: bid})
	}
	var i, sum int
	for h.Len() > 0 {
		hand := heap.Pop(h).(hand)
		i++
		sum += hand.bid * i
	}
	println(sum)
}
