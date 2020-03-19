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

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
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

////////////////////////////////////////////////////////////////////////////////

func unionBits(words []int) int {
	word := 0
	for _, w := range words {
		word |= w
	}
	return word
}

func hasBit(n int, pos int) bool {
	val := n & (1 << uint(pos))
	return val > 0

}

func setBit(n int, pos int) int {
	return n | (1 << uint(pos))
}

func convert2Bit(word string) int {
	bits := 0
	for _, c := range word {
		bit := 1 << uint(c-'a')
		bits |= bit
	}
	return bits
}

func makeExistIdxs(words []int) []int {
	existBits := unionBits(words)
	idxs := make([]int, 0)
	for i := 0; i < A; i++ {
		if hasBit(existBits, i) {
			idxs = append(idxs, i)
		}
	}
	return idxs
}

func find(knownChars int, curIdx int, leftCount int) int {
	if leftCount > len(existIdxs)-curIdx && false {
		return 0
	} else if leftCount == 0 {
		count := 0
		for _, word := range words {
			unknownChars := ^knownChars
			if unknownChars&word == 0 {
				count++
			}
		}
		return count
	}

	ret := 0

	for i := curIdx; i < len(existIdxs); i++ {
		count := find(setBit(knownChars, existIdxs[i]), i+1, leftCount-1)
		ret = maxInt(ret, count)
	}
	return ret
}

func solve(K int, orgWords []string) int {
	words = make([]int, len(orgWords))
	for i, orgWord := range orgWords {
		words[i] = convert2Bit(orgWord)
	}
	existIdxs = makeExistIdxs(words)

	if len(existIdxs) <= K {
		return len(orgWords)
	}
	return find(0, 0, K)
}

// A : size of Alphabet
const A = 26

var (
	words     []int
	existIdxs []int
)

func main() {
	defer writer.Flush()

	var n, k int
	scanf("%d %d\n", &n, &k)

	words := make([]string, n)
	var s string
	for i := range words {
		scanf("%s\n", &s)
		s = removeCharacters(s, "antic")
		words[i] = s
	}

	printf("%d\n", solve(k-5, words))
}
