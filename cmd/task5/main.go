package main

import (
	"bufio"
	"fmt"
	"os"
)

type Cell struct {
	T bool
	B bool
	L bool
	R bool
}

var (
	CellEmpty = Cell{}
	CellFull  = Cell{T: true, B: true, L: true, R: true}
)

type BGrid struct {
	g  [][]byte
	sx int
	sy int
}

type grid struct {
	g  [][]Cell
	sx int
	sy int
}

type Point struct {
	x int
	y int
}

type Op struct {
	s int
	e int
}

var gout *bufio.Writer

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	gout = out
	defer out.Flush()
	var cnt int

	// gg := grid{sx: 1, sy: 3}

	// _numToPoint(&gg, 5)
	// return

	fmt.Fscan(in, &cnt)
	cnt = 1
	for i := 1; i <= cnt; i++ {
		// fmt.Fprintf(out, "i: %v\n", i)
		doTask(in, out)
	}
}

func doTask(in *bufio.Reader, out *bufio.Writer) {
	var sy, sx, cnt int
	fmt.Fscan(in, &sy, &sx, &cnt)
	fmt.Fscanln(in)

	g := readGrid(in, sx, sy)
	// ops := readOps(in, cnt)
	// for _, op := range ops {
	// 	if !g.applyOp(op) {
	// 		fmt.Fprintln(out, "Unsuported ops (may be later)")
	// 		return
	// 	}
	// 	g.printGrid(out)
	// }

	readOps(in, cnt)

	// g.cleanUR(Point{x: 1, y: 6})
	// g.mirorDRtoL(Point{x: 1, y: 1})
	g.printGrid(out)

	g.mirorDLtoR(Point{x: 1, y: 2})
	g.printGrid(out)

	// g.mirorDRtoL(Point{x: 1, y: 2})
	// g.printGrid(out)

	// g.mirorURtoL(Point{x: 1, y: 6})
	// g.printGrid(out)

	// g.mirorURtoL(Point{x: 1, y: 5})
	// g.mirorULtoR(Point{x: 2, y: 3})

	// g.mirorDLtoR(Point{x: 3, y: 1})
	// g.mirorDLtoR(Point{x: 1, y: 1})
	// g.mirorDRtoL(Point{x: 3, y: 1})
	// g.cleanDR(Point{x: 0, y: 0})
	// g.cleanDL(Point{x: 0, y: 0})

	// g.printGrid(out)

}

func (g *grid) mirorDLtoR(p Point) {
	e := g.sy
	if e < g.sx {
		e = g.sx
	}
	eg := g._expand(e, e)

	p.x += -1 + e
	p.y += -1 + e

	for p.x != 0 && p.y != 0 {
		p.x -= 1
		p.y -= 1
	}

	for i := 0; (p.x+i) < eg.sx && (p.y+i) < eg.sy; i++ {
		dx := p.x + i
		dy := p.y + i

		eg.g[dy][dx] = eg.g[dy][dx].mirorDLtoR(true)

		// eg.g[dy][dx] = Cell{T: true, R: true}

		for y := 1; dy+y < eg.sy && dx+y < eg.sx; y++ {
			eg.g[dy][dx+y] = mergeCell(eg.g[dy][dx+y], eg.g[dy+y][dx].mirorDLtoR(false))
			eg.g[dy+y][dx] = CellEmpty
			// eg.g[dy+y][dx] = Cell{T: true, B: true}
		}

	}

	eg._reduce()
	*g = *eg
}

func (g *grid) mirorDRtoL(p Point) {
	e := g.sy
	if e < g.sx {
		e = g.sx
	}
	eg := g._expand(e, e)

	p.x += -1 + e
	p.y += -1 + e

	for p.x != 0 && p.y != 0 {
		p.x -= 1
		p.y -= 1
	}

	for i := 0; (p.x+i) < eg.sx && (p.y+i) < eg.sy; i++ {
		dx := p.x + i
		dy := p.y + i

		eg.g[dy][dx] = eg.g[dy][dx].mirorDRtoL(true)

		// eg.g[dy][dx] = Cell{T: true, R: true}

		for y := 1; dy-y >= 0 && dx-y >= 0; y++ {
			eg.g[dy][dx-y] = mergeCell(eg.g[dy][dx-y], eg.g[dy-y][dx].mirorDRtoL(false))

			eg.g[dy-y][dx] = CellEmpty
			// eg.g[dy][dx-y] = Cell{T: true, B: true}
		}

	}

	eg._reduce()
	*g = *eg
}

func (g *grid) mirorULtoR(p Point) {
	e := g.sy
	if e < g.sx {
		e = g.sx
	}
	eg := g._expand(e, e)

	p.x += -1 + e
	p.y += -1 + e

	for p.x != 0 && p.y != eg.sy {
		p.x -= 1
		p.y += 1
	}

	for i := 0; (p.x+i) < eg.sx && (p.y-i) > 0; i++ {
		dx := p.x + i
		dy := p.y - i - 1

		eg.g[dy][dx] = eg.g[dy][dx].mirorULtoR(true)

		// eg.g[dy][dx] = Cell{T: true, L: true}

		for y := 1; dy-y >= 0 || dx+y >= eg.sx; y++ {
			eg.g[dy][dx+y] = mergeCell(eg.g[dy][dx+y], eg.g[dy-y][dx].mirorULtoR(false))

			// eg.g[dy][dx+y] = Cell{T: true}
			// eg.g[dy-y][dx] = Cell{T: true, B: true}

			eg.g[dy-y][dx] = CellEmpty
		}

	}

	eg._reduce()
	*g = *eg
}

func (g *grid) mirorURtoL(p Point) {
	e := g.sy
	if e < g.sx {
		e = g.sx
	}
	eg := g._expand(e, e)

	p.x += -1 + e
	p.y += -1 + e

	for p.x != 0 && p.y != eg.sy {
		p.x -= 1
		p.y += 1
	}

	for i := 0; (p.x+i) < eg.sx && (p.y-i) > 0; i++ {
		dx := p.x + i
		dy := p.y - i - 1

		eg.g[dy][dx] = eg.g[dy][dx].mirorURtoL(true)

		// eg.g[dy][dx] = Cell{T: true, L: true}

		for y := 1; dy+y < eg.sy && dx-y >= 0; y++ {
			eg.g[dy][dx-y] = mergeCell(eg.g[dy][dx-y], eg.g[dy+y][dx].mirorURtoL(false))
			eg.g[dy+y][dx] = CellEmpty
			// eg.g[dy][dx-y] = Cell{T: true, B: true}
		}

	}

	// eg._reduce()
	*g = *eg
}

// func (g *grid) mirorDLtoR(p Point) {
// 	e := g.sy
// 	if e < g.sx {
// 		e = g.sx
// 	}
// 	eg := g._expand(e, e)

// 	p.x -= 1
// 	p.y -= 1

// 	for i := 0; (i+p.x) < g.sx && (i+p.y) < g.sy; i++ {
// 		eg.g[e+i+p.y][e+i+p.x] = eg.g[e+i+p.y][e+i+p.x].mirorDLtoR(true)
// 		for y := i + 1; y < g.sy; y++ {
// 			eg.g[e+p.y+i][e+p.x+y] = mergeCell(eg.g[e+p.y+i][e+p.x+y], eg.g[e+p.y+y][e+p.x+i].mirorDLtoR(false))
// 		}
// 	}

// 	p.x += e
// 	p.y += e
// 	eg.cleanDL(p)
// 	eg._reduce()
// 	*g = *eg
// }

func (g *grid) cleanUR(p Point) {
	p.x -= 1
	p.y -= 1

	for y := 0; y < g.sy; y++ {
		for x := 0; x < g.sx; x++ {
			if p.y-y < x+1 {
				g.g[y][x] = CellEmpty
			}
		}
	}
}

func (g *grid) cleanUL(p Point) {
	p.x -= 1
	p.y -= 1

	for y := 0; y < g.sy; y++ {
		for x := 0; x < g.sx; x++ {
			if p.y-y > x+1 {
				g.g[y][x] = CellEmpty
			}
		}
	}
}

func (g *grid) cleanDR(p Point) {
	p.x -= 1
	p.y -= 1

	for y := 0; y < g.sy; y++ {
		for x := 0; x < g.sx; x++ {
			if (y - p.y) < (x - p.x) {
				g.g[y][x] = CellEmpty
			}
		}
	}
}

func (g *grid) cleanDL(p Point) {
	p.x -= 1
	p.y -= 1

	for y := 0; y < g.sy; y++ {
		for x := 0; x < g.sx; x++ {
			if (y - p.y) > (x - p.x) {
				g.g[y][x] = CellEmpty
			}
		}
	}
}

func (c Cell) ToByte() byte {
	switch {
	case !c.T && !c.R && !c.B && !c.L:
		return '.'
	case c.T && c.R && c.B && c.L:
		return '#'
	case !c.T && !c.R && c.B && !c.L, c.T && c.R && !c.B && c.L:
		return '^'
	case !c.T && !c.R && !c.B && c.L, c.T && c.R && c.B && !c.L:
		return '>'
	case c.T && !c.R && !c.B && !c.L, !c.T && c.R && c.B && c.L:
		return 'v'
	case !c.T && c.R && !c.B && !c.L, c.T && !c.R && c.B && c.L:
		return '<'
	case c.T && !c.R && !c.B && c.L, !c.T && c.R && c.B && !c.L:
		return '/'
	case c.T && c.R && !c.B && !c.L, !c.T && !c.R && c.B && c.L:
		return '\\'
	case c.T && !c.R && c.B && !c.L, !c.T && c.R && !c.B && c.L:
		return 'x'
	}
	return '?'
}

func (g *grid) ToBGrid() *BGrid {
	bg := BGrid{
		sx: g.sx,
		sy: g.sy,
		g:  make([][]byte, g.sy),
	}

	for y := 0; y < g.sy; y++ {
		bg.g[y] = make([]byte, len(g.g[y]))
		for x := 0; x < g.sx; x++ {
			bg.g[y][x] = g.g[y][x].ToByte()
		}
	}
	return &bg
}

func (g *grid) applyOp(op Op) bool {
	s := _numToPoint(g, op.s)
	e := _numToPoint(g, op.e)
	if s.y == e.y {
		if s.x < e.x {
			g.mirorBtoT(s.y)
			return true
		} else {
			g.mirorTtoB(s.y)
			return true
		}
	} else if s.x == e.x {
		if s.y > e.y {
			g.mirorRtoL(s.x)
			return true
		} else {
			g.mirorLtoR(s.x)
			return true
		}
	}

	return false

}

func _numToPoint(g *grid, n int) Point {
	w := g.sx
	h := g.sy

	if n <= w+1 {
		return Point{x: n, y: 1}
	}
	n -= w

	if n <= h {
		return Point{x: w + 1, y: n}
	}
	n -= h

	if n <= w {
		return Point{x: w - n + 2, y: h + 1}
	}
	n -= w

	if n <= h {
		return Point{x: 1, y: h - n + 2}
	}

	panic(
		fmt.Sprintf("cant parce point: g[%v:%v], n=%v", g.sx, g.sy, n),
	)
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
			// eg.g[ey+dst_y][x] = mereCell(eg.g[ey+dst_y][x], eg.g[ey+src_y][x])
			eg.g[ey+dst_y][x] = mergeCell(eg.g[ey+dst_y][x], eg.g[ey+src_y][x].mirorTtoB())
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
			// eg.g[dst_y+ey][x] = mereCell(eg.g[dst_y+ey][x], eg.g[src_y+ey][x])
			eg.g[dst_y+ey][x] = mergeCell(eg.g[dst_y+ey][x], eg.g[src_y+ey][x].mirorTtoB())
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
			// eg.g[y][dst_x+ex] = mereCell(eg.g[y][dst_x+ex], eg.g[y][src_x+ex])
			eg.g[y][dst_x+ex] = mergeCell(eg.g[y][dst_x+ex], eg.g[y][src_x+ex].mirorLtoR())
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
			eg.g[y][dst_x+ex] = mergeCell(eg.g[y][dst_x+ex], eg.g[y][src_x+ex].mirorLtoR())
		}
	}
	eg.cleanL(col + ex)
	eg._reduce()
	*g = *eg
}

func (g *grid) _expand(ex int, ey int) *grid {
	ng := grid{
		g: make([][]Cell, 0, g.sy+(2*ey)),
	}
	for y := 0; y < ey; y++ {
		// ng.g = append(ng.g, bytes.Repeat([]Cell{}, g.sx+(ex*2)))
		ng.g = append(ng.g, repeatCell(Cell{}, g.sx+(ex*2)))

	}

	for y := 0; y < g.sy; y++ {
		line := make([]Cell, 0, g.sx+(2*ex))
		line = append(line, repeatCell(Cell{}, ex)...)
		line = append(line, g.g[y]...)
		line = append(line, repeatCell(Cell{}, ex)...)
		ng.g = append(ng.g, line)
	}

	for y := 0; y < ey; y++ {
		ng.g = append(ng.g, repeatCell(Cell{}, g.sx+(ex*2)))
	}
	ng.sx = g.sx + (2 * ex)
	ng.sy = g.sy + (2 * ey)

	return &ng
}

func (g *grid) _reduce() {
	fx := g.sx
	lx := 0

	fy := g.sy
	ly := 0

	for y := 0; y < g.sy; y++ {
		clrY := true
		for x := 0; x < g.sx; x++ {
			if g.g[y][x] != CellEmpty {
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
			if y < fy {
				fy = y
			}
			if y > ly {
				ly = y
			}
		}
	}

	g.g = g.g[fy : ly+1]
	g.sy = len(g.g)
	g.sx = lx - fx + 1

	for y := 0; y < g.sy; y++ {
		g.g[y] = g.g[y][fx : lx+1]
	}

}

func (g *grid) cleanT(line int) {
	for y := 0; y < line; y++ {
		for x := 0; x < g.sx; x++ {
			g.g[y][x] = CellEmpty
		}
	}
}

func (g *grid) cleanB(line int) {
	for y := line; y < g.sy; y++ {
		for x := 0; x < g.sx; x++ {
			g.g[y][x] = CellEmpty
		}
	}
}

func (g *grid) cleanR(col int) {
	for y := 0; y < g.sy; y++ {
		for x := col; x < g.sx; x++ {
			g.g[y][x] = CellEmpty
		}
	}
}

func (g *grid) cleanL(col int) {
	for y := 0; y < g.sy; y++ {
		for x := 0; x < col; x++ {
			g.g[y][x] = CellEmpty
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
		g:  make([][]Cell, sy),
		sx: sx,
		sy: sy,
	}
	// fmt.Fscan(in, &g.sy, &g.sx)
	// fmt.Fscanln(in)

	for y := 0; y < g.sy; y++ {
		var bline []byte
		fmt.Fscanln(in, &bline)
		g.g[y] = make([]Cell, sx)
		for x := 0; x < g.sx; x++ {
			if bline[x] == '#' {
				g.g[y][x] = CellFull
			} else {
				g.g[y][x] = CellEmpty
			}
		}
	}
	return &g
}

func (g *grid) printGrid(out *bufio.Writer) {
	// fmt.Fprintln(out)

	bg := g.ToBGrid()

	// fmt.Fprintf(out, "[%v x %v]\n", g.sx, g.sy)
	for i := 0; i < bg.sy; i++ {
		fmt.Fprintln(out, string(bg.g[i]))
	}
	fmt.Fprintln(out)
}

func (c Cell) mirorDLtoR(clr bool) Cell {
	nc := Cell{T: c.L, R: c.B}
	if !clr {
		nc.L = c.T
		nc.B = c.R
	}
	return nc
}

func (c Cell) mirorDRtoL(clr bool) Cell {
	nc := Cell{L: c.T, B: c.R}
	if !clr {
		nc.T = c.L
		nc.R = c.B
	}
	return nc
}

func (c Cell) mirorURtoL(clr bool) Cell {
	nc := Cell{T: c.R, L: c.B}
	if !clr {
		nc.R = c.T
		nc.B = c.L
	}
	return nc
}

func (c Cell) mirorULtoR(clr bool) Cell {
	nc := Cell{R: c.T, B: c.L}
	if !clr {
		nc.T = c.R
		nc.L = c.B
	}
	return nc
}

func (c Cell) mirorTtoB() Cell {
	tmp := c.T
	c.T = c.B
	c.B = tmp
	return c
}

func (c Cell) mirorLtoR() Cell {
	tmp := c.L
	c.L = c.R
	c.R = tmp
	return c
}

func mergeCell(a Cell, b Cell) Cell {
	return Cell{
		T: a.T || b.T,
		B: a.B || b.B,
		L: a.L || b.L,
		R: a.R || b.R,
	}
}

func repeatCell(c Cell, cnt int) []Cell {
	res := make([]Cell, cnt)

	for i := range res {
		res[i] = c
	}

	return res
}
