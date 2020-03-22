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

const UintMax = ^uint(0)
const UnitMin = 0
const IntMax = int(^uint(0) >> 1)
const IntMin = -IntMax - 1

func printProb(prob []string) {
	for _, s := range prob {
		printf("%s\n", s)
	}
}

// Stack :
type Stack []Point

func (s *Stack) push(e Point) {
	*s = append(*s, e)
}

func (s *Stack) pop() Point {
	ret := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ret
}

// Queue :
type Queue []interface{}

func newQueue() Queue {
	return make(Queue, 0)
}

func newQueueWith(xs ...interface{}) Queue {
	que := newQueue()
	for _, x := range xs {
		que.push(x)
	}
	return que
}

func (q *Queue) push(e interface{}) {
	*q = append(*q, e)
}

func (q *Queue) pop() interface{} {
	ret := (*q)[0]
	*q = (*q)[1:]
	return ret
}

////////////////////////////////////////////////////////////////////////////////

// Point :
type Point struct {
	r, c int
}

func (p Point) move(x Point) Point {
	return Point{p.r + x.r, p.c + x.c}
}

var directions = []Point{Point{0, 1}, Point{1, 0}, Point{0, -1}, Point{-1, 0}}

func (p Point) inBound() bool {
	return 0 <= p.r && p.r < R && 0 <= p.c && p.c < C
}

func (p Point) getMark(prob []string) byte {
	if p.inBound() {
		return prob[p.r][p.c]
	}
	return '*'
}

func (p Point) isVisited(visited [][]bool) bool {
	return visited[p.r][p.c]
}

func (p Point) setVisited(visited [][]bool) {
	if p.inBound() {
		visited[p.r][p.c] = true
	}
}

func findMark(prob []string, mark byte) []Point {
	points := make([]Point, 0)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if prob[i][j] == mark {
				points = append(points, Point{i, j})
			}
		}
	}
	return points
}

func readMapProb(R, C int) []string {
	prob := make([]string, R)
	var s string
	for i := 0; i < R; i++ {
		scanf("%s\n", &s)
		prob[i] = s
	}
	return prob
}

func makeVisited(R, C int) [][]bool {
	visited := make([][]bool, R)
	for r := 0; r < R; r++ {
		visited[r] = make([]bool, C)
	}
	return visited
}

func makeMultiVisited(R, C, n int) [][][]bool {
	visited := make([][][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = makeVisited(R, C)
	}
	return visited
}

func printVisited(visited [][]bool) {
	for _, row := range visited {
		for _, b := range row {
			if b {
				printf("O")
			} else {
				printf("X")
			}
		}
		printf("\n")
	}
}

////////////////////////////////////////////////////////////////////////////////

func solve(prob []string) int {
	Ls := findMark(prob, 'L')
	start := Ls[0]
	dest := Ls[1]
	que := newQueueWith(start)
	visited := makeVisited(R, C)

	wQue := newQueue()
	for _, w := range findMark(prob, '.') {
		wQue.push(w)
	}
	wQue.push(start)
	wQue.push(dest)
	wVisited := makeVisited(R, C)

	step := 0
	for len(wQue) > 0 {
		step++
		wQueLen := len(wQue)
		for i := 0; i < wQueLen; i++ {
			curPos := wQue.pop().(Point)
			for _, d := range directions {
				pos := curPos.move(d)
				if !pos.inBound() || pos.isVisited(wVisited) {
					continue
				}
				pos.setVisited(wVisited)
				mark := pos.getMark(prob)
				if mark == 'X' {
					wQue.push(pos)
				}
			}
		}
		queLen := len(que)
		for i := 0; i < queLen; i++ {
			subQue := newQueueWith(que.pop())
			for len(subQue) > 0 {
				curPos := subQue.pop().(Point)
				for _, d := range directions {
					pos := curPos.move(d)
					if !pos.inBound() || pos.isVisited(visited) {
						continue
					}
					pos.setVisited(visited)
					mark := pos.getMark(prob)
					if mark == '.' || (mark == 'X' && pos.isVisited(wVisited)) {
						subQue.push(pos)
					} else if mark == 'X' {
						que.push(pos)
					} else if mark == 'L' && pos.r == dest.r && pos.c == dest.c {
						return step
					}
				}
			}
		}
	}
	return -1
}

var (
	R, C int
)

func main() {
	defer writer.Flush()

	scanf("%d %d\n", &R, &C)
	prob := readMapProb(R, C)
	printf("%d\n", solve(prob))
}
