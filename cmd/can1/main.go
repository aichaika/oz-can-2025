package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Grid struct {
	grid   [][]byte
	size_x int
	size_y int
}

var in *bufio.Reader
var out *bufio.Writer

func main() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var task_cnt int

	fmt.Fscan(in, &task_cnt)
	fmt.Fscanln(in)
	// task_cnt = 1
	for i := 1; i <= task_cnt; i++ {
		// fmt.Fprintf(out, "test num: %v\n", i)
		doTask()
	}
}

func doTask() {
	var cnt_line, cnt_diag int
	fmt.Fscan(in, &cnt_line, &cnt_diag)
	fmt.Fscanln(in)

	size_y := cnt_diag*2 + 1
	size_x := cnt_line + cnt_diag*2

	g := MakeGrid(size_x, size_y, ' ')
	g.DrawHex(cnt_line, cnt_diag)

	g.Print()

}

func (g *Grid) DrawHex(cnt_line int, cnt_diag int) {
	for x := 0; x < cnt_line; x++ {
		g.grid[0][cnt_diag+x] = '_'
		g.grid[g.size_y-1][cnt_diag+x] = '_'
	}

	for y := 0; y < cnt_diag; y++ {
		g.grid[1+y][-1+cnt_diag-y] = '/'
		g.grid[1+y][cnt_diag+cnt_line+y] = '\\'
		g.grid[1+cnt_diag+y][y] = '\\'
		g.grid[1+cnt_diag+y][-1+cnt_diag+cnt_diag+cnt_line-y] = '/'
	}
}

func MakeGrid(size_x int, size_y int, char byte) *Grid {
	g := Grid{
		size_x: size_x,
		size_y: size_y,
	}

	g.grid = make([][]byte, size_y)
	for y := range g.grid {
		g.grid[y] = make([]byte, size_x)
		for x := range g.grid[y] {
			g.grid[y][x] = char
		}
	}

	return &g
}

func (g *Grid) Print() {
	// fmt.Fprintln)

	// fmt.Fprintf(out, "[%v x %v]\n", g.size_x, g.size_y)
	for i := 0; i < g.size_y; i++ {
		fmt.Fprintln(out, strings.TrimRight(string(g.grid[i]), " "))
	}
	// fmt.Fprintln()
}
