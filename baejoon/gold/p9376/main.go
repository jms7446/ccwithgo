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

// IntMax :
const IntMax = int(^uint(0) >> 1)

// MinInt :
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

////////////////////////////////////////////////////////////////////////////////

// Point :
type Point struct {
	r, c int
}

func (p Point) move(x Point) Point {
	return Point{p.r + x.r, p.c + x.c}
}

func (p Point) inBound() bool {
	return 0 <= p.r && p.r < R && 0 <= p.c && p.c < C
}

func (p Point) getMark(smap []string) byte {
	if p.inBound() {
		return smap[p.r][p.c]
	}
	return '*'
}

func (p Point) setMark(smap [][]byte, mark byte) {
	if p.inBound() {
		smap[p.r][p.c] = mark
	}
}

func (p Point) getCost(arr [][]int) int {
	return arr[p.r][p.c]
}

func (p Point) setCost(arr [][]int, cost int) {
	arr[p.r][p.c] = cost
}

func (p Point) getCostSum(arrs [][][]int) int {
	costSum := 0
	for _, arr := range arrs {
		cost := p.getCost(arr)
		if cost == IntMax {
			return IntMax
		}
		costSum += cost
	}
	return costSum
}

// PointStack :
type PointStack []Point

func (s *PointStack) push(e Point) {
	*s = append(*s, e)
}

func (s *PointStack) pop() Point {
	ret := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ret
}

func makeCostMap(n int) [][][]int {
	costMap := make([][][]int, n)
	for k := 0; k < n; k++ {
		costMap[k] = make([][]int, R)
		for i := 0; i < R; i++ {
			costMap[k][i] = make([]int, C)
			for j := 0; j < C; j++ {
				costMap[k][i][j] = IntMax
			}
		}
	}
	return costMap
}

func findPoints(prob []string, target byte) []Point {
	prisoners := make([]Point, 0)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if prob[i][j] == target {
				prisoners = append(prisoners, Point{i, j})
			}
		}
	}
	return prisoners
}

func allEmpty(ques [][]Point) bool {
	for _, que := range ques {
		if len(que) > 0 {
			return false
		}
	}
	return true
}

func calcCost(pos Point, mark byte, costMap [][][]int) int {
	costSum := pos.getCostSum(costMap)
	if costSum < IntMax {
		if mark == '#' {
			return costSum - 2
		}
		return costSum
	}
	return costSum
}

func takeStep(startPos Point, step int, costs [][]int, prob []string) []Point {
	que := make([]Point, 0)
	stack := PointStack{startPos}
	for len(stack) > 0 {
		curPos := stack.pop()
		for _, direction := range directions {
			pos := curPos.move(direction)
			if !pos.inBound() {
				continue
			}
			mark := pos.getMark(prob)
			cost := pos.getCost(costs)
			if cost != IntMax || mark == '*' {
				continue
			}
			if mark == '#' {
				pos.setCost(costs, step+1)
				que = append(que, pos)
			} else {
				pos.setCost(costs, step)
				stack.push(pos)
				if mark == '$' {
					que = append(que, pos)
				}
			}
		}
	}
	return que
}

func solve(prob []string) int {
	// (N * R * C)
	costMap := makeCostMap(N)

	starts := append(findPoints(prob, '$'), Point{0, 0})
	ques := make([][]Point, N)
	for i, start := range starts {
		ques[i] = []Point{start}
	}

	minTotalCost := IntMax
	for step := 0; !allEmpty(ques) && step < minTotalCost; step++ {
		for i := 0; i < N; i++ {
			nextQue := make([]Point, 0)
			for _, curPos := range ques[i] {
				mark := curPos.getMark(prob)
				cost := calcCost(curPos, mark, costMap)
				// printf("step: %02d, cost: %d\n", step, cost)
				minTotalCost = MinInt(minTotalCost, cost)
				if step == 0 || mark == '#' {
					poss := takeStep(curPos, step, costMap[i], prob)
					for _, pos := range poss {
						nextQue = append(nextQue, pos)
					}
				}
			}
			ques[i] = nextQue
		}
	}
	return minTotalCost
}

// N :
const N = 3

var (
	// R :
	R int
	// C :
	C          int
	directions = []Point{Point{0, 1}, Point{1, 0}, Point{0, -1}, Point{-1, 0}}
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
