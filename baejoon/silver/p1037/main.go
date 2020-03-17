package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func main() {
	defer writer.Flush()

	var n int
	scanf("%d\n", &n)

	nums := make([]int, 0, 50)
	var x int
	for i := 0; i < n; i++ {
		scanf("%d", &x)
		nums = append(nums, x)
	}

	sort.Ints(nums)
	printf("%d\n", nums[0]*nums[n-1])
}
