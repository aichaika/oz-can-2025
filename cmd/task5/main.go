package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type grid struct {
	g  [][]byte
	sx int
	sy int
}

// type point struct {
// 	x int
// 	y int
// }

type Op struct {
	s int
	e int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var cnt int

	fmt.Fscan(in, &cnt)
	cnt = 1
	for i := 1; i <= cnt; i++ {
		doTask(in, out)
	}
}

func doTask(in *bufio.Reader, out *bufio.Writer) {
	var sy, sx, cnt int
	fmt.Fscan(in, &sy, &sx, &cnt)
	fmt.Fscanln(in)

	g := readGrid(in, sx, sy)
	readOps(in, cnt)

	// cg := g.copyGrid()
	g.printGrid(out)

	// g.cleanR(3)
	// g.printGrid(out)

	// g.cleanL(1)
	// g.printGrid(out)
	// cg.printGrid(out)

	// g.mirorRtoL(4)
	// g.mirorLtoR(3)
	g.mirorTtoB(4)
	// g.mirorBtoT(2)
	// g = g._expand(0, 3)
	// g.printGrid(out)

	// g._reduce()
	g.printGrid(out)

}

func (g *grid) mirorTtoB(line int) {
	ey := g.sy
	eg := g._expand(0, ey)

	line -= 1
	shift_y_max := line
	for shift_y := 0; shift_y < shift_y_max; shift_y++ {
		src_y := line - shift_y - 1
		dst_y := line + shift_y
		for x := 0; x < eg.sx; x++ {
			eg.g[ey+dst_y][x] = eg.g[ey+src_y][x]
		}
	}
	eg.cleanT(line + ey)
	eg._reduce()
	*g = *eg
}

func (g *grid) mirorBtoT(line int) {
	ey := g.sy
	eg := g._expand(0, ey)

	line -= 1
	shift_y_max := g.sy - line
	for shift_y := 0; shift_y < shift_y_max; shift_y++ {
		src_y := line + shift_y
		dst_y := line - shift_y - 1
		for x := 0; x < g.sx; x++ {
			eg.g[dst_y+ey][x] = eg.g[src_y+ey][x]
		}
	}
	eg.cleanB(line + ey)
	eg._reduce()
	*g = *eg
}

func (g *grid) mirorRtoL(col int) {
	ex := g.sx
	eg := g._expand(ex, 0)

	col -= 1
	shift_x_max := g.sx - col
	for y := 0; y < g.sy; y++ {
		for shift_x := 0; shift_x < shift_x_max; shift_x++ {
			src_x := col + shift_x
			dst_x := col - shift_x - 1
			eg.g[y][dst_x+ex] = eg.g[y][src_x+ex]
		}
	}
	eg.cleanR(col + ex)
	eg._reduce()
	*g = *eg
}

func (g *grid) mirorLtoR(col int) {
	ex := g.sx
	eg := g._expand(ex, 0)

	col -= 1
	shift_x_max := col
	for y := 0; y < g.sy; y++ {
		for shift_x := 0; shift_x < shift_x_max; shift_x++ {
			src_x := col - shift_x - 1
			dst_x := col + shift_x

			eg.g[y][dst_x+ex] = eg.g[y][src_x+ex]
		}
	}
	eg.cleanL(col + ex)
	eg._reduce()
	*g = *eg
}

func (g *grid) _expand(ex int, ey int) *grid {
	ng := grid{
		g: make([][]byte, 0, g.sy+(2*ey)),
	}
	for y := 0; y < ey; y++ {
		ng.g = append(ng.g, bytes.Repeat([]byte{'.'}, g.sx+(ex*2)))
	}

	for y := 0; y < g.sy; y++ {
		line := make([]byte, 0, g.sx+(2*ex))
		line = append(line, bytes.Repeat([]byte{'.'}, ex)...)
		line = append(line, g.g[y]...)
		line = append(line, bytes.Repeat([]byte{'.'}, ex)...)
		ng.g = append(ng.g, line)
	}

	for y := 0; y < ey; y++ {
		ng.g = append(ng.g, bytes.Repeat([]byte{'.'}, g.sx+(ex*2)))
	}
	ng.sx = g.sx + (2 * ex)
	ng.sy = g.sy + (2 * ey)

	return &ng
}

func (g *grid) _reduce() {
	ng := grid{
		g: [][]byte{},
	}
	fx := g.sx
	lx := 0
	for y := 0; y < g.sy; y++ {
		clrY := true
		for x := 0; x < g.sx; x++ {
			if g.g[y][x] != '.' {
				clrY = false
				if x < fx {
					fx = x
				}
				if x > lx {
					lx = x
				}
			}
		}
		if !clrY {
			ng.g = append(ng.g, g.g[y])
		}
	}

	g.g = ng.g
	g.sy = len(ng.g)
	g.sx = lx - fx + 1

	for y := 0; y < g.sy; y++ {
		g.g[y] = g.g[y][fx : lx+1]
	}

}

func (g *grid) cleanT(line int) {
	for y := 0; y < line; y++ {
		for x := 0; x < g.sx; x++ {
			g.g[y][x] = '.'
		}
	}
}

func (g *grid) cleanB(line int) {
	for y := line; y < g.sy; y++ {
		for x := 0; x < g.sx; x++ {
			g.g[y][x] = '.'
		}
	}
}

func (g *grid) cleanR(col int) {
	for y := 0; y < g.sy; y++ {
		for x := col; x < g.sx; x++ {
			g.g[y][x] = '.'
		}
	}
}

func (g *grid) cleanL(col int) {
	for y := 0; y < g.sy; y++ {
		for x := 0; x < col; x++ {
			g.g[y][x] = '.'
		}
	}
}

// func (g *grid) copyGrid() *grid {
// 	ng := grid{
// 		sx: g.sx,
// 		sy: g.sy,
// 		g:  make([][]byte, g.sy),
// 	}

// 	for i := range g.g {
// 		ng.g[i] = make([]byte, len(g.g[i]))
// 		copy(ng.g[i], g.g[i])
// 	}

// 	return &ng
// }

// func (g *grid) mergeGrid(in *grid) {
// 	for y := 0; y < g.sy; y++ {
// 		for x := 0; x < g.sx; x++ {
// 			g.g[y][x] = in.g[y][x]
// 		}
// 	}
// }

func readOps(in *bufio.Reader, cnt int) []Op {
	ops := make([]Op, 0, cnt)

	for i := 0; i < cnt; i++ {
		op := Op{}
		fmt.Fscan(in, &op.s, &op.e)
		fmt.Fscanln(in)
		ops = append(ops, op)
	}
	return ops
}

func readGrid(in *bufio.Reader, sx int, sy int) *grid {
	g := grid{
		g:  [][]byte{},
		sx: sx,
		sy: sy,
	}
	// fmt.Fscan(in, &g.sy, &g.sx)
	// fmt.Fscanln(in)

	for i := 0; i < g.sy; i++ {
		var line []byte
		fmt.Fscanln(in, &line)
		g.g = append(g.g, line)
	}
	return &g
}

func (g *grid) printGrid(out *bufio.Writer) {
	// fmt.Fprintln(out)

	fmt.Fprintf(out, "[%v x %v]\n", g.sx, g.sy)
	for i := 0; i < g.sy; i++ {
		fmt.Fprintln(out, string(g.g[i]))
	}
	fmt.Fprintln(out)
}
