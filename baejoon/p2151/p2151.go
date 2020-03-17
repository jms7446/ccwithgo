package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func init() {
	scanner.Split(bufio.ScanWords)
}

func readWord() string {
	scanner.Scan()
	return scanner.Text()
}

func readInt() int {
	scanner.Scan()
	x, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	return x
}

////////////////////////////////////////////////////////////////////////////////

// Pos :
type Pos struct {
	r, c int
}

// Info :
type Info struct {
	pos Pos
	d   int
	m   int
}

// InfoQueue :
type InfoQueue []Info

func (q *InfoQueue) push(x Info) {
	*q = append(*q, x)
}

func (q *InfoQueue) pop() Info {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

func isValidIndex(n int, idxs ...int) bool {
	for _, i := range idxs {
		if i < 0 || i >= n {
			return false
		}
	}
	return true
}

func findMark(c byte, hmap []string) []Pos {
	startEnd := make([]Pos, 0)
	for r, row := range hmap {
		for c, mark := range row {
			if mark == '#' {
				startEnd = append(startEnd, Pos{r, c})
			}
		}
	}
	return startEnd
}

func stepNext(pos Pos, d int) Pos {
	nextPos := Pos{
		pos.r + directions[d].r,
		pos.c + directions[d].c,
	}
	return nextPos
}

func solve(hmap []string) int {
	N := len(hmap)
	sharps := findMark('#', hmap)
	start, end := sharps[0], sharps[1]

	que := make(InfoQueue, 0)
	for d := range directions {
		que.push(Info{start, d, 0})
	}

	visited := make(map[Pos]bool)
	for len(que) > 0 {
		info := que.pop()
		var pos = info.pos
		for {
			pos = stepNext(pos, info.d)
			if !isValidIndex(N, pos.r, pos.c) {
				break
			}
			// fmt.Printf("    : %c, %+v\n", hmap[pos.r][pos.c], pos)
			mark := hmap[pos.r][pos.c]
			if mark == '*' {
				break
			} else if mark == '#' {
				if pos.r == end.r && pos.c == end.c {
					return info.m
				}
			} else if mark == '!' {
				if visited[pos] {
					break
				} else {
					visited[pos] = true
				}
				for _, d := range reflects[info.d] {
					que.push(Info{pos, d, info.m + 1})
				}
			}
		}
	}
	return -1
}

var directions []Pos
var reflects [][]int

func main() {
	directions = []Pos{Pos{0, 1}, Pos{1, 0}, Pos{0, -1}, Pos{-1, 0}}
	reflects = [][]int{[]int{1, 3}, []int{0, 2}, []int{1, 3}, []int{0, 2}}

	N := readInt()
	hmap := make([]string, N)
	for i := 0; i < N; i++ {
		hmap[i] = readWord()
	}
	fmt.Println(solve(hmap))
}
