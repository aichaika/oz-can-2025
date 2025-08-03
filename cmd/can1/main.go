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

	var cnt int

	fmt.Fscan(in, &cnt)
	fmt.Fscanln(in)
	// cnt = 1
	for i := 1; i <= cnt; i++ {
		// fmt.Fprintf(out, "test num: %v\n", i)
		doTask(in, out)
	}
}

func doTask(in *bufio.Reader, out *bufio.Writer) {
	var sl, ss int
	fmt.Fscan(in, &sl, &ss)
	fmt.Fscanln(in)

	sy := ss*2 + 1
	sx := sl + ss*2

	// fmt.Fprintf(gout, "sl: %v\n", sl)
	// fmt.Fprintf(gout, "ss: %v\n", ss)
	// fmt.Fprintf(gout, "sy: %v\n", sy)
	// fmt.Fprintf(gout, "sx: %v\n", sx)

	// fmt.Printf("sy: %v\n", sy)
	// fmt.Printf("sx: %v\n", sx)

	g := makeGrid(sx, sy)
	g.makeFigure(sl, ss)

	g.printGrid(out)

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
