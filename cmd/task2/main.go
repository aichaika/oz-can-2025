package main

import (
	"bufio"
	"fmt"
	"os"
)

type grid struct {
	g  [][]byte
	sx int
	sy int
}

type point struct {
	x int
	y int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var cnt int

	fmt.Fscan(in, &cnt)
	for i := 1; i <= cnt; i++ {
		// fmt.Fprintf(out, "test num: %v\n", i)
		doTask(in, out)
	}
}

func doTask(in *bufio.Reader, out *bufio.Writer) {
	g := readGrid(in)

	A := g.findChar(out, 'A')
	B := g.findChar(out, 'B')

	if A.y < B.y || (A.y == B.y && A.x < B.x) {
		move00(g, A, 'A', 'a')
		moveEE(g, B, 'B', 'b')
	} else {
		move00(g, B, 'B', 'b')
		moveEE(g, A, 'A', 'a')
	}

	g.printGrid(out)
}

func move00(g *grid, P point, M byte, m byte) {
	p := P
	if p.y != 0 && g.g[p.y-1][p.x] != '#' {
		p = moveU(g, p, m)
		p = moveL(g, p, m)
	} else {
		p = moveL(g, p, m)
		p = moveU(g, p, m)
	}
	g.g[P.y][P.x] = M
}

func moveEE(g *grid, P point, M byte, m byte) {
	p := P
	if p.y != (g.sy-1) && g.g[p.y+1][p.x] != '#' {
		p = moveD(g, p, m)
		p = moveR(g, p, m)
	} else {
		p = moveR(g, p, m)
		p = moveD(g, p, m)
	}
	g.g[P.y][P.x] = M
}

func moveU(g *grid, p point, m byte) point {
	for y := p.y; y >= 0; y-- {
		g.g[y][p.x] = m
	}
	return point{x: p.x, y: 0}
}

func moveD(g *grid, p point, m byte) point {
	for y := p.y; y < g.sy; y++ {
		g.g[y][p.x] = m
	}

	return point{x: p.x, y: g.sy - 1}
}

func moveL(g *grid, p point, m byte) point {
	for x := p.x; x >= 0; x-- {
		g.g[p.y][x] = m
	}
	return point{x: 0, y: p.y}
}

func moveR(g *grid, p point, m byte) point {
	for x := p.x; x < g.sx; x++ {
		g.g[p.y][x] = m
	}
	return point{x: g.sx - 1, y: p.y}
}

func (g *grid) findChar(out *bufio.Writer, c byte) point {
	for y := 0; y < g.sy; y++ {
		for x := 0; x < g.sx; x++ {
			if g.g[y][x] == c {
				return point{x: x, y: y}
			}
		}
	}
	panic(0)
}

func readGrid(in *bufio.Reader) *grid {
	g := grid{
		g: [][]byte{},
	}
	fmt.Fscan(in, &g.sy, &g.sx)
	fmt.Fscanln(in)

	for i := 0; i < g.sy; i++ {
		var line []byte
		fmt.Fscanln(in, &line)
		g.g = append(g.g, line)
	}
	return &g
}

func (g *grid) printGrid(out *bufio.Writer) {
	// fmt.Fprintf(out, "[%v x %v]\n", g.sx, g.sy)
	for i := 0; i < g.sy; i++ {
		fmt.Fprintln(out, string(g.g[i]))
	}
	// fmt.Fprintln(out)
}

// func (g *grid) getPoint(x int, y int) byte {
// 	if x < 0 || x >= g.sx || y < 0 || y >= g.sy {
// 		return 0
// 	}
// 	return g.g[y][x]
// }

// func (g *grid) setPoint(x int, y int, c byte) {
// 	if x < 0 || x >= g.sx || y < 0 || y >= g.sy {
// 		return
// 	}
// 	g.g[y][x] = c
// }
