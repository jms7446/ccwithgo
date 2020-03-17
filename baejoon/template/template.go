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
func readInt2() (int, int)      { return readInt(), readInt() }
func readInt3() (int, int, int) { return readInt(), readInt(), readInt() }

func readIntsN(n int) []int {
	var ints = make([]int, 0, n)
	for i := 0; i < n; i++ {
		ints = append(ints, readInt())
	}
	return ints
}

func main() {
	N, M := readInt2()
	K := readInt()
	nums := readIntsN(K)
	fmt.Println(N, M, K)
	fmt.Println(nums)
}
