package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

////////////////////////////////////////////////////////////////////////////////

func solve(N int, K int, nums []int) int {
	if N == 1 {
		return 0
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	totalLen := nums[N-1] - nums[0]

	gaps := make([]int, N-1)
	for i := range gaps {
		gaps[i] = nums[i+1] - nums[i]
	}
	sort.Slice(gaps, func(i, j int) bool { return gaps[i] > gaps[j] })

	for i := 0; i < K-1; i++ {
		totalLen -= gaps[i]
	}

	return totalLen
}

func main() {
	N, K := readInt(), readInt()
	nums := readIntsN(N)
	fmt.Println(solve(N, K, nums))
}
