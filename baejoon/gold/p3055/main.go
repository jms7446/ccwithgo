package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

// Queue :
type Queue []interface{}

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

var directions = []Point{Point{0, 1}, Point{1, 0}, Point{0, -1}, Point{-1, 0}}

func (p Point) move(x Point) Point {
	return Point{p.r + x.r, p.c + x.c}
}

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

func readMapProb(R, C int) []string {
	prob := make([]string, R)
	var s string
	for i := 0; i < R; i++ {
		scanf("%s\n", &s)
		prob[i] = s
	}
	return prob
}

////////////////////////////////////////////////////////////////////////////////

func solve(prob []string) int {
	que := Queue{findMark(prob, 'S')[0]}
	wQue := make(Queue, 0)
	for _, p := range findMark(prob, '*') {
		wQue = append(wQue, p)
	}
	visited := makeVisited(R, C)
	step := 0
	for len(que) > 0 {
		step++
		lenWQue := len(wQue)
		for i := 0; i < lenWQue; i++ {
			curPos := wQue.pop().(Point)
			for _, d := range directions {
				pos := curPos.move(d)
				if !pos.inBound() || pos.isVisited(visited) {
					continue
				}
				mark := pos.getMark(prob)
				if mark == '.' || mark == 'S' {
					wQue.push(pos)
				}
				if mark != 'D' {
					pos.setVisited(visited)
				}
			}
		}
		lenQue := len(que)
		for i := 0; i < lenQue; i++ {
			curPos := que.pop().(Point)
			for _, d := range directions {
				pos := curPos.move(d)
				if !pos.inBound() || pos.isVisited(visited) {
					continue
				}
				pos.setVisited(visited)
				mark := pos.getMark(prob)
				if mark == '.' {
					que.push(pos)
				} else if mark == 'D' {
					return step
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
	ret := solve(prob)
	if ret >= 0 {
		printf("%d\n", ret)
	} else {
		printf("KAKTUS\n")
	}
}
