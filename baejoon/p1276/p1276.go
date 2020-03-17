package p1276

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

// MaxSlice :
func MaxSlice(xs []bridge, valueFunc func(bridge) int) (m int) {
	m = valueFunc(xs[0])
	var v int
	for _, x := range xs[1:] {
		v = valueFunc(x)
		if v > m {
			m = v
		}
	}
	return
}

type bridge struct {
	y, x1, x2 int
}

func newBridge(y, x1, x2 int) *bridge {
	b := new(bridge)
	b.y, b.x1, b.x2 = y, x1, x2
	return b
}

// func solve1(bs []bridge) int {
// 	sort.Slice(bs, func(i, j int) bool { return bs[i].y < bs[j].y })
// 	maxWidth := MaxSlice(bs, func(b bridge) int { return b.x2 })
// 	// fmt.Println(maxWidth)
// 	floors := make([]int, maxWidth)
// 	result := 0
// 	for _, b := range bs {
// 		result += b.y - floors[b.x1]
// 		result += b.y - floors[b.x2-1]
// 		for i := b.x1; i < b.x2; i++ {
// 			floors[i] = b.y
// 		}
// 	}
// 	return result
// }

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func solve2(bs []bridge) int {
	result := 0
	for _, a := range bs {
		var x1Max, x2Max int
		for _, b := range bs {
			if b.y >= a.y {
				continue
			}
			if a.x1 >= b.x1 && a.x1 < b.x2 {
				x1Max = max(x1Max, b.y)
			}
			if a.x2 > b.x1 && a.x2 <= b.x2 {
				x2Max = max(x2Max, b.y)
			}
		}
		result += 2*a.y - x1Max - x2Max
	}
	return result
}

func main() {
	N := readInt()
	bs := make([]bridge, N)
	for i := 0; i < N; i++ {
		bs[i] = *newBridge(readInt3())
	}
	fmt.Println(solve2(bs))
}
