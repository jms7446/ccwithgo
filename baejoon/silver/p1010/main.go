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

////////////////////////////////////////////////////////////////////////////////

// L :
const L = 30

var table [L][L]int

func makeTable() {
	for j := 1; j < L; j++ {
		table[0][j] = 1
		table[j][j] = 1
	}
	for i := 1; i < L; i++ {
		for j := i + 1; j < L; j++ {
			table[i][j] = table[i-1][j-1] + table[i][j-1]
		}
	}
}

func main() {
	defer writer.Flush()

	makeTable()

	var t int
	scanf("%d\n", &t)

	var n, m int
	for i := 0; i < t; i++ {
		scanf("%d %d\n", &n, &m)
		printf("%d\n", table[n][m])
	}
}
