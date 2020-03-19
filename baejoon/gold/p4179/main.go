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

func (p Pos) add(other Pos) Pos {
	return Pos{p.r + other.r, p.c + other.c}
}

func solve(prob []string) int {
	jq := make([]Pos, 0)
	fq := make([]Pos, 0)
	visited := make([][]bool, R)
	for i := range visited {
		visited[i] = make([]bool, C)
	}
	for r, str := range prob {
		for c := range str {
			if str[c] == 'J' {
				jq = append(jq, Pos{r, c})
				visited[r][c] = true
			} else if str[c] == 'F' {
				fq = append(fq, Pos{r, c})
				visited[r][c] = true
			}
		}
	}

	stepCount := 0
	for len(jq) > 0 {
		stepCount++

		fqSize := len(fq)
		for i := 0; i < fqSize; i++ {
			curPos := fq[0]
			fq = fq[1:]
			for _, d := range directions {
				pos := curPos.add(d)
				if pos.inBound() && prob[pos.r][pos.c] != '#' && !visited[pos.r][pos.c] {
					visited[pos.r][pos.c] = true
					fq = append(fq, pos)
				}
			}
		}

		jqSize := len(jq)
		for i := 0; i < jqSize; i++ {
			curPos := jq[0]
			jq = jq[1:]
			for _, d := range directions {
				pos := curPos.add(d)
				if !pos.inBound() {
					return stepCount
				}
				if !visited[pos.r][pos.c] && prob[pos.r][pos.c] == '.' {
					visited[pos.r][pos.c] = true
					jq = append(jq, pos)
				}
			}
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
