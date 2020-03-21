package template

////////////////////////////////////////////////////////////////////////////////
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

func printProb(prob []string) {
	for _, s := range prob {
		printf("%s\n", s)
	}
}

func makeMultiVisited(n int) [][][]bool {
	visited := make([][][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([][]bool, R)
		for r := 0; r < R; r++ {
			visited[i][r] = make([]bool, C)
		}
	}
	return visited
}
