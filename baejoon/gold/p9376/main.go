package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func minMax(array []int) (int, int) {
	var max = array[0]
	var min = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func removeCharacters(input string, characters string) string {
	filter := func(r rune) rune {
		if strings.IndexRune(characters, r) < 0 {
			return r
		}
		return -1
	}

	return strings.Map(filter, input)

}

////////////////////////////////////////////////////////////////////////////////

// Pos :
type Pos struct {
	r, c int
}

func (p Pos) move(x Pos) Pos {
	return Pos{p.r + x.r, p.c + x.c}
}

func (p Pos) inBound() bool {
	return 0 <= p.r && p.r < R && 0 <= p.c && p.c < C
}

func (p Pos) getMark(smap [][]byte) byte {
	if p.inBound() {
		return smap[p.r][p.c]
	}
	return '*'
}

func (p Pos) setMark(smap [][]byte, mark byte) {
	if p.inBound() {
		smap[p.r][p.c] = mark
	}
}

var directions = []Pos{Pos{0, 1}, Pos{1, 0}, Pos{0, -1}, Pos{-1, 0}}

func makeMap(prob []string) [][]byte {
	smap := make([][]byte, R)
	for r, row := range prob {
		smap[r] = make([]byte, C)
		for c := range row {
			smap[r][c] = row[c]
		}
	}
	return smap
}

// SearchInfo :
type SearchInfo struct {
	pos     Pos
	history []Pos
}

func newSearchInfo(pos Pos, preHistory []Pos) SearchInfo {
	history := make([]Pos, len(preHistory))
	copy(history, preHistory)
	history = append(history, pos)
	return SearchInfo{pos, history}
}

// InfoQueue :
type InfoQueue []SearchInfo

func (q InfoQueue) push(s SearchInfo) InfoQueue {
	return append(q, s)
}

func (q InfoQueue) pop() (InfoQueue, SearchInfo) {
	return q[1:], q[0]
}

// PosStack :
type PosStack []Pos

func (s PosStack) push(p Pos) PosStack {
	return append(s, p)
}

func (s PosStack) pop() (PosStack, Pos) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func searchStep(smap [][]byte, que []SearchInfo) (InfoQueue, []Pos) {
	newQue := make([]SearchInfo, 0)
	for _, info := range que {
		stack := PosStack{info.pos}
		var curPos Pos
		for len(stack) > 0 {
			stack, curPos = stack.pop()
			for _, d := range directions {
				pos := curPos.move(d)
				mark := pos.getMark(smap)
				pos.setMark(smap, '*')
				if mark == '*' {
					continue
				} else if mark == '#' {
					newQue = append(newQue, newSearchInfo(pos, info.history))
				} else if mark == '$' {
					newInfo := newSearchInfo(pos, info.history)
					return nil, newInfo.history
				} else if mark == '.' {
					stack = stack.push(pos)
				}
			}

		}
	}
	return newQue, nil
}

func search(smap [][]byte, start Pos) []Pos {
	var successHistory []Pos
	que := InfoQueue{SearchInfo{start, make([]Pos, 0)}}
	for len(que) > 0 {
		que, successHistory = searchStep(smap, que)
		if successHistory != nil {
			return successHistory
		}
	}
	return nil
}

func unlockDoors(smap [][]byte, history []Pos) {
	for _, pos := range history {
		pos.setMark(smap, '.')
	}
}

func solve(prob []string) int {
	var smap [][]byte
	start := Pos{0, 0}

	smap = makeMap(prob)
	preHistory := search(smap, start)
	fmt.Println(preHistory)

	smap = makeMap(prob)
	unlockDoors(smap, preHistory)
	PostHistory := search(smap, start)
	fmt.Println(PostHistory)

	return len(preHistory) + len(PostHistory) - 2
}

var (
	R, C int
)

func main() {
	defer writer.Flush()

	var T int
	scanf("%d\n", &T)

	for t := 0; t < T; t++ {
		scanf("%d %d\n", &R, &C)

		prob := make([]string, 0)
		prob = append(prob, strings.Repeat(".", C+2))
		var s string
		for i := 0; i < R; i++ {
			scanf("%s\n", &s)
			prob = append(prob, "."+s+".")
		}
		prob = append(prob, strings.Repeat(".", C+2))
		R = R + 2
		C = C + 2
		// fmt.Println(len(prob), len(prob[0]))
		printf("%d\n", solve(prob))
	}
}
