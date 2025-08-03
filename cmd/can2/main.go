package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type grid struct {
	g  [][]byte
	sx int
	sy int
}

var gout *bufio.Writer

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	gout = out

	doTask(in, out)
}

func doTask(in *bufio.Reader, out *bufio.Writer) {
	var sl, ss, sx, sy, c int
	fmt.Fscan(in, &sx, &sy, &ss, &sl, &c)
	fmt.Fscanln(in)

	hsy := ss*2 + 1
	hsx := sl + ss*2

	// fmt.Fprintf(gout, "sl: %v\n", sl)
	// fmt.Fprintf(gout, "ss: %v\n", ss)
	// fmt.Fprintf(gout, "sy: %v\n", sy)
	// fmt.Fprintf(gout, "sx: %v\n", sx)

	// fmt.Printf("sy: %v\n", sy)
	// fmt.Printf("sx: %v\n", sx)

	h := makeGrid(hsx, hsy)
	h.makeFigure(sl, ss)

	m := makeGrid(sx+2, sy+2)
	m.makeMap()

	// cx := sx / hsx
	// cy := sy / hsy

	psx := 1
	psy := 1

	// y := 0
	// x := 0
	pc := 0
	for {

		for i := 0; i < c; i++ {
			var x, y int
			if i%2 == 0 {
				x = psx + (i/2)*(hsx+sl)
				y = psy
			} else {
				x = psx + ss + 1 + (i/2)*(hsx+sl)
				y = psy + ss
			}

			if (x + hsx) > sx+1 {
				break
			} else {
				m.mergeGrid(h, x, y)
				pc += 1
			}
		}
		c -= pc
		if c <= 0 {
			break
		}
		psy += 2 * ss

	}

	// fmt.Fprintf(gout, "sl: %v\n", c)

	m.printGrid(out)
}

func (m *grid) mergeGrid(g *grid, sx int, sy int) {
	for y := 0; y < g.sy; y++ {
		for x := 0; x < g.sx; x++ {
			if g.g[y][x] == ' ' {
				continue
			}
			m.g[y+sy][x+sx] = g.g[y][x]
		}
	}
}

func (g *grid) makeMap() {
	for x := 0; x < g.sx; x++ {
		g.g[0][x] = '-'
		g.g[g.sy-1][x] = '-'
	}
	for y := 0; y < g.sy; y++ {
		g.g[y][0] = '|'
		g.g[y][g.sx-1] = '|'
	}

	g.g[0][0] = '+'
	g.g[0][g.sx-1] = '+'
	g.g[g.sy-1][0] = '+'
	g.g[g.sy-1][g.sx-1] = '+'
}

func (g *grid) makeFigure(sl int, ss int) {
	for x := 0; x < sl; x++ {
		g.g[0][ss+x] = '_'
		g.g[g.sy-1][ss+x] = '_'
	}

	for y := 1; y < g.sy/2+1; y++ {
		g.g[y][g.sy/2-y] = '/' // +

		g.g[y][sl+ss+y-1] = 'x' // -
		// g.g[y][sl+ss+y-1] = 'x' // -

		g.g[y][sl+ss+y-1] = '\\'        // -
		g.g[y+ss][sl+ss+g.sy/2-y] = '/' // -

		g.g[y+ss][y-1] = '\\' // +
	}
}

func makeGrid(sx int, sy int) *grid {
	g := grid{
		sx: sx,
		sy: sy,
	}

	g.g = make([][]byte, sy)
	for y := range g.g {
		g.g[y] = make([]byte, sx)
		for x := range g.g[y] {
			g.g[y][x] = ' '
		}
	}

	return &g
}

func (g *grid) printGrid(out *bufio.Writer) {
	// fmt.Fprintln(out)

	// fmt.Fprintf(out, "[%v x %v]\n", g.sx, g.sy)
	for i := 0; i < g.sy; i++ {
		fmt.Fprintln(out, strings.TrimRight(string(g.g[i]), " "))
	}
	// fmt.Fprintln(out)
}
