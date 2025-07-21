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
	// cnt = 1
	for i := 1; i <= cnt; i++ {
		// fmt.Fprintf(out, "test num: %v\n", i)
		doTask(in, out)
	}
}

func doTask(in *bufio.Reader, out *bufio.Writer) {
	g := readGrid(in)

	p := point{x: 0, y: 0}
	ki := map[byte]bool{}

	for {
		// g.printGrid(out)
		p, c, ok := findIsland(out, g, p)
		if !ok {
			break
		}
		if ki[c] {
			// 	fmt.Fprintf(out, "map: %v x:%v\n", ki, string(c))
			// 	g.printGrid(out)
			fmt.Fprintln(out, "NO")
			return
		}
		ki[c] = true
		cleanIsland(out, g, p)
	}
	fmt.Fprintln(out, "YES")

	// A := g.findChar(out, 'A')
	// B := g.findChar(out, 'B')

	// if A.y < B.y || (A.y == B.y && A.x < B.x) {
	// 	move00(g, A, 'A', 'a')
	// 	moveEE(g, B, 'B', 'b')
	// } else {
	// 	move00(g, B, 'B', 'b')
	// 	moveEE(g, A, 'A', 'a')
	// }

	// g.printGrid(out)
	// fmt.Fprintln(out, "YES")
}

func cleanIsland(out *bufio.Writer, g *grid, p point) {
	c := g.g[p.y][p.x]
	todo := []point{p}

	for i := 0; i < len(todo); i++ {
		todo = append(todo, cleanBinded(out, g, todo[i], c)...)
	}
}

func cleanBinded(out *bufio.Writer, g *grid, ps point, c byte) []point {
	g.g[ps.y][ps.x] = '.'
	pb := make([]point, 0, 6)
	pc := make([]point, 0, 6)

	pb = append(pb, point{x: ps.x - 1, y: ps.y - 1})
	pb = append(pb, point{x: ps.x + 1, y: ps.y - 1})

	pb = append(pb, point{x: ps.x - 2, y: ps.y})
	pb = append(pb, point{x: ps.x + 2, y: ps.y})

	pb = append(pb, point{x: ps.x - 1, y: ps.y + 1})
	pb = append(pb, point{x: ps.x + 1, y: ps.y + 1})

	for _, p := range pb {
		if p.x < 0 || p.y < 0 || p.x >= g.sx || p.y >= g.sy {
			continue
		}
		if g.g[p.y][p.x] == c {
			pc = append(pc, p)
			g.g[p.y][p.x] = '.'
		}
	}
	return pc
}

func findIsland(out *bufio.Writer, g *grid, p point) (point, byte, bool) {
	for y := p.y; y < g.sy; y++ {
		for x := p.x; x < g.sx; x++ {
			if g.g[y][x] != '.' {
				return point{x: x, y: y}, g.g[y][x], true
			}
		}
	}
	return point{}, 0, false
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
	fmt.Fprintln(out)

	fmt.Fprintf(out, "[%v x %v]\n", g.sx, g.sy)
	for i := 0; i < g.sy; i++ {
		fmt.Fprintln(out, string(g.g[i]))
	}
	// fmt.Fprintln(out)
}
