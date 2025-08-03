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

var gout *bufio.Writer

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	gout = out

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
	g.printGrid(out)

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
