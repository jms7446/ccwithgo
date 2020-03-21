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
	return prob[p.r][p.c]
}

func (p Point) isVisited(visited [][]bool) bool {
	return visited[p.r][p.c]
}

func (p Point) setVisited(visited [][]bool) {
	visited[p.r][p.c] = true
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

func printProb(prob []string) {
	for _, s := range prob {
		printf("%s\n", s)
	}
}

////////////////////////////////////////////////////////////////////////////////

func getKeyIndex(c byte, preIndex int) int {
	return preIndex | (1 << uint(c-'a'))
}

func hasKey(c byte, keyIndex int) bool {
	return (keyIndex & (1 << uint(c-'A'))) != 0
}

func makeMultiVisited(n int) [][][]bool {
	visited := make([][][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([][]bool, R)
		for r := 0; r < R; r++ {
			visited[i][r] = make([]bool, C)
		}
	}
	return visited
}

// Info :
type Info struct {
	pos      Point
	keyIndex int
}

func solve() int {
	visited := makeMultiVisited(NUM_MAP)
	start := findMark(prob, '0')
	que := Queue{Info{start[0], 0}}
	step := 0
	for len(que) > 0 {
		numQue := len(que)
		step++
		for i := 0; i < numQue; i++ {
			curInfo := que.pop().(Info)
			for _, d := range directions {
				pos := curInfo.pos.move(d)
				if !pos.inBound() || pos.isVisited(visited[curInfo.keyIndex]) {
					continue
				}
				pos.setVisited(visited[curInfo.keyIndex])
				mark := pos.getMark(prob)
				if mark == '1' {
					return step
				} else if mark == '#' {
					continue
				} else if mark == '.' || mark == '0' {
					que.push(Info{pos, curInfo.keyIndex})
				} else if 'A' <= mark && mark <= 'F' {
					if hasKey(mark, curInfo.keyIndex) {
						que.push(Info{pos, curInfo.keyIndex})
					}
				} else if 'a' <= mark && mark <= 'f' {
					keyIndex := getKeyIndex(mark, curInfo.keyIndex)
					que.push(Info{pos, keyIndex})
					pos.setVisited(visited[keyIndex])
				}
			}
		}
	}

	return -1
}

var (
	R, C int
	prob []string
)

const (
	NUM_KEY = 6
	NUM_MAP = 1 << 6
)

func main() {
	defer writer.Flush()

	scanf("%d %d\n", &R, &C)

	prob = make([]string, R)
	var s string
	for i := 0; i < R; i++ {
		scanf("%s\n", &s)
		prob[i] = s
	}
	printf("%d\n", solve())
}
