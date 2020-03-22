package template

// Point :
type Point struct {
	r, c int
}

func (p Point) move(x Point) Point {
	return Point{p.r + x.r, p.c + x.c}
}

var directions = []Point{Point{0, 1}, Point{1, 0}, Point{0, -1}, Point{-1, 0}}

func (p Point) inBound() bool {
	return 0 <= p.r && p.r < R && 0 <= p.c && p.c < C
}

func (p Point) getMark(prob []string) byte {
	if p.inBound() {
		return prob[p.r][p.c]
	}
	return '*'
}

func (p Point) isVisited(visited [][]bool) bool {
	return visited[p.r][p.c]
}

func (p Point) setVisited(visited [][]bool) {
	if p.inBound() {
		visited[p.r][p.c] = true
	}
}

func findMark(prob []string, mark byte) []Point {
	points := make([]Point, 0)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if prob[i][j] == mark {
				points = append(points, Point{i, j})
			}
		}
	}
	return points
}

func readMapProb(R, C int) []string {
	prob := make([]string, R)
	var s string
	for i := 0; i < R; i++ {
		scanf("%s\n", &s)
		prob[i] = s
	}
	return prob
}

func makeVisited(R, C int) [][]bool {
	visited := make([][]bool, R)
	for r := 0; r < R; r++ {
		visited[r] = make([]bool, C)
	}
	return visited
}

func makeMultiVisited(R, C, n int) [][][]bool {
	visited := make([][][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = makeVisited(R, C)
	}
	return visited
}

func printVisited(visited [][]bool) {
	for _, row := range visited {
		for _, b := range row {
			if b {
				printf("O")
			} else {
				printf("X")
			}
		}
		printf("\n")
	}
}
