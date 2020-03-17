package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func init() {
	scanner.Split(bufio.ScanWords)
}

func readInt() int {
	scanner.Scan()
	x, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	return x
}

func readIntsN(n int) []int {
	var ints = make([]int, n)
	for i := range ints {
		ints[i] = readInt()
	}
	return ints
}

func min(xs ...int) int {
	ret := xs[0]
	for i := 1; i < len(xs); i++ {
		if xs[i] < ret {
			ret = xs[i]
		}
	}
	return ret
}

////////////////////////////////////////////////////////////////////////////////

func main() {
	N, M := readInt(), readInt()
	K := readInt()
	nums := readIntsN(K)
	fmt.Println(N, M, K)
	fmt.Println(nums)
}
