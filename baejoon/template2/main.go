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

func (q *Queue) push(e interface{}) {
	*q = append(*q, e)
}

func (q *Queue) pop() interface{} {
	ret := (*q)[0]
	*q = (*q)[1:]
	return ret
}

////////////////////////////////////////////////////////////////////////////////

func main() {
	defer writer.Flush()

	var t int
	scanf("%d\n", &t)

	var n, m int
	for i := 0; i < t; i++ {
		scanf("%d %d\n", &n, &m)
	}
}
