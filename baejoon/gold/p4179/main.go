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

////////////////////////////////////////////////////////////////////////////////

// Pos :
type Pos struct {
	r, c int
}

func (p Pos) inBound() bool {
	return 0 <= p.r && p.r < R && 0 <= p.c && p.c < C
}

func (p Pos) move(other Pos) Pos {
	return Pos{p.r + other.r, p.c + other.c}
}

func initProb(rawProb []string) (prob [][]byte, jPos Pos, fPoses []Pos) {
	fPoses = make([]Pos, 0)
	prob = make([][]byte, R)
	for r, str := range rawProb {
		prob[r] = make([]byte, C)
		for c := range str {
			mark := str[c]
			prob[r][c] = mark
			if mark == 'J' {
				jPos = Pos{r, c}
			} else if mark == 'F' {
				fPoses = append(fPoses, Pos{r, c})
			}
		}
	}
	return
}

func fireMove(fq []Pos, prob [][]byte) []Pos {
	newFq := make([]Pos, 0)
	for _, curPos := range fq {
		for _, d := range directions {
			pos := curPos.move(d)
			if pos.inBound() && prob[pos.r][pos.c] == '.' {
				prob[pos.r][pos.c] = 'F'
				newFq = append(newFq, pos)
			}
		}
	}
	return newFq
}

func jMove(jq []Pos, prob [][]byte) ([]Pos, bool) {
	newJq := make([]Pos, 0)
	for _, curPos := range jq {
		for _, d := range directions {
			pos := curPos.move(d)
			if !pos.inBound() {
				return nil, true
			} else if prob[pos.r][pos.c] == '.' {
				prob[pos.r][pos.c] = 'J'
				newJq = append(newJq, pos)
			}
		}
	}
	return newJq, false
}

func solve(rawProb []string) int {
	prob, jPos, fPoses := initProb(rawProb)

	jq := make([]Pos, 0)
	jq = append(jq, jPos)
	prob[jPos.r][jPos.c] = 'J'

	fq := make([]Pos, 0)
	for _, p := range fPoses {
		fq = append(fq, p)
		prob[p.r][p.c] = 'F'
	}

	stepCount := 0
	var success bool
	for len(jq) > 0 {
		stepCount++
		fq = fireMove(fq, prob)
		jq, success = jMove(jq, prob)
		if success {
			return stepCount
		}
	}
	return -1
}

var (
	R, C       int
	directions = [4]Pos{Pos{0, 1}, Pos{1, 0}, Pos{-1, 0}, Pos{0, -1}}
)

func main() {
	defer writer.Flush()

	scanf("%d %d\n", &R, &C)

	prob := make([]string, 0)
	var s string
	for i := 0; i < R; i++ {
		scanf("%s\n", &s)
		prob = append(prob, s)
	}

	ret := solve(prob)
	if ret != -1 {
		printf("%d\n", ret)
	} else {
		printf("IMPOSSIBLE\n")
	}
}
