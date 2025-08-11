package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var DEBUG_ON = false

var in *bufio.Reader
var out *bufio.Writer

func main() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	if DEBUG_ON {
		fmt.Fprintln(out, "DEBUG ON")
	}

	doTask()
}

func doTask() {
	var hex_cnt_line, hex_cnt_diag, hex_cnt int
	var size_x, size_y int

	fmt.Fscan(in,
		&size_x, &size_y,
		&hex_cnt_line, &hex_cnt_diag, &hex_cnt,
	)
	fmt.Fscanln(in)

	// size_y = 13
	// size_x = 8

	hex_size_y := hex_cnt_diag*2 + 1
	hex_size_x := hex_cnt_line + hex_cnt_diag*2

	h := MakeGrid(hex_size_x, hex_size_y, ' ')
	h.DrawHex(hex_cnt_line, hex_cnt_diag)

	b := MakeGrid(size_x+2, size_y+2, ' ')
	b.drawBorder()

	g := MakeGrid(size_x, size_y, ' ')

	max_cnt_x2 := (g.size_x-h.size_x)/(h.size_x+hex_cnt_line) + 1
	max_cnt_x1 := (g.size_x - hex_cnt_diag) / (h.size_x + hex_cnt_line)
	max_cnt_x := max_cnt_x2
	if max_cnt_x1 > max_cnt_x {
		max_cnt_x = max_cnt_x1
	}

	max_cnt_y := (g.size_y-1)/hex_cnt_diag - 1

	l := MakeGrid(max_cnt_x, max_cnt_y, 'x')
	// l.grid[0][0] = 'x'
	// l.grid[0][1] = 'x'

	for y := 0; y < l.size_y; y++ {
		for x := 0; x < l.size_x; x++ {
			if l.grid[y][x] == ' ' {
				continue
			}
			var shift_x, shift_y int
			if y%2 == 0 {
				if x > max_cnt_x2 {
					continue
				}
				shift_x = (h.size_x + hex_cnt_line) * x
				shift_y = y * hex_cnt_diag
				g.DrawGrid(h, shift_x, shift_y)
			} else {
				if x >= max_cnt_x1 {
					continue
				}
				shift_x = (h.size_x+hex_cnt_line)*x + hex_cnt_line + hex_cnt_diag
				shift_y = y * hex_cnt_diag
				g.DrawGrid(h, shift_x, shift_y)
			}
		}
	}

	b.DrawGrid(g, 1, 1)
	b.Print()
}

// func (m *grid) mergeGrid(g *grid, sx int, sy int) {
// 	for y := 0; y < g.sy; y++ {
// 		for x := 0; x < g.sx; x++ {
// 			if g.g[y][x] == ' ' {
// 				continue
// 			}
// 			m.g[y+sy][x+sx] = g.g[y][x]
// 		}
// 	}
// }

func (g *Grid) drawBorder() {
	for x := 0; x < g.size_x; x++ {
		g.grid[0][x] = '-'
		g.grid[g.size_y-1][x] = '-'
	}
	for y := 0; y < g.size_y; y++ {
		g.grid[y][0] = '|'
		g.grid[y][g.size_x-1] = '|'
	}

	g.grid[0][0] = '+'
	g.grid[0][g.size_x-1] = '+'
	g.grid[g.size_y-1][0] = '+'
	g.grid[g.size_y-1][g.size_x-1] = '+'
}

// func (g *grid) makeFigure(sl int, ss int) {
// 	for x := 0; x < sl; x++ {
// 		g.g[0][ss+x] = '_'
// 		g.g[g.sy-1][ss+x] = '_'
// 	}

// 	for y := 1; y < g.sy/2+1; y++ {
// 		g.g[y][g.sy/2-y] = '/' // +

// 		g.g[y][sl+ss+y-1] = 'x' // -
// 		// g.g[y][sl+ss+y-1] = 'x' // -

// 		g.g[y][sl+ss+y-1] = '\\'        // -
// 		g.g[y+ss][sl+ss+g.sy/2-y] = '/' // -

// 		g.g[y+ss][y-1] = '\\' // +
// 	}
// }

// func makeGrid(sx int, sy int) *grid {
// 	g := grid{
// 		sx: sx,
// 		sy: sy,
// 	}

// 	g.g = make([][]byte, sy)
// 	for y := range g.g {
// 		g.g[y] = make([]byte, sx)
// 		for x := range g.g[y] {
// 			g.g[y][x] = ' '
// 		}
// 	}

// 	return &g
// }

// func (g *grid) printGrid(out *bufio.Writer) {
// 	// fmt.Fprintln(out)

// 	// fmt.Fprintf(out, "[%v x %v]\n", g.sx, g.sy)
// 	for i := 0; i < g.sy; i++ {
// 		fmt.Fprintln(out, strings.TrimRight(string(g.g[i]), " "))
// 	}
// 	// fmt.Fprintln(out)
// }

type Grid struct {
	grid   [][]byte
	size_x int
	size_y int
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

func (g *Grid) DrawGrid(in_g *Grid, shift_x int, shift_y int) {
	for y := 0; y < in_g.size_y; y++ {
		for x := 0; x < in_g.size_x; x++ {
			if in_g.grid[y][x] == ' ' {
				continue
			}
			target_x := shift_x + x
			target_y := shift_y + y

			if DEBUG_ON && (target_x > g.size_x) {
				fmt.Fprintf(out, "overflov grid by x (%v > %v)\n", target_x > g.size_x)
			}
			if DEBUG_ON && (target_y > g.size_y) {
				fmt.Fprintf(out, "overflov grid by x (%v > %v)\n", target_y > g.size_x)
			}

			g.grid[target_y][target_x] = in_g.grid[y][x]
		}
	}
}

func (g *Grid) Print() {
	// fmt.Fprintln)

	if DEBUG_ON {
		fmt.Fprintf(out, "[%v x %v]\n", g.size_x, g.size_y)
	}

	for i := 0; i < g.size_y; i++ {
		fmt.Fprintln(out, strings.TrimRight(string(g.grid[i]), " "))
	}
	// fmt.Fprintln()
}
