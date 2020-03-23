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

////////////////////////////////////////////////////////////////////////////////

func decompose(s string) []string {
	a1 := strings.Split(s, "=")
	a2 := strings.Split(a1[0], "+")
	return []string{a2[0], a2[1], a1[1]}
}

func analyze(s string) (map[byte]int, int) {
	t := 0
	m := make(map[byte]int)
	for i := len(s) - 1; i >= 0; {
		if '1' <= s[i] && s[i] <= '9' {
			count := int(s[i] - '0')
			m[s[i-1]] += count
			t += count
			i -= 2
		} else {
			m[s[i]]++
			t++
			i--
		}
	}
	return m, t
}

func solve(s string) (int, int, int) {
	ss := decompose(s)
	m1, t1 := analyze(ss[0])
	m2, t2 := analyze(ss[1])
	m3, t3 := analyze(ss[2])
	for x1 := 1; x1 <= 10; x1++ {
		for x2 := 1; x2 <= 10; x2++ {
			found := true
			leftSum := x1*t1 + x2*t2
			x3 := leftSum / t3
			if leftSum%t3 != 0 || x3 > 10 {
				continue
			}
			for k, v := range m3 {
				if x1*m1[k]+x2*m2[k] != x3*v {
					found = false
					break
				}
			}
			if found {
				return x1, x2, x3
			}
		}
	}
	return -1, -1, -1
}

func main() {
	defer writer.Flush()

	var s string
	scanf("%s\n", &s)
	x1, x2, x3 := solve(s)
	printf("%d %d %d\n", x1, x2, x3)
}
