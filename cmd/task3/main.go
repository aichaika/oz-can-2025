package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type bline []byte
type bgreed []bline

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var taskCount int

	fmt.Fscan(in, &taskCount)
	fmt.Fscanln(in)
	for i := 0; i < taskCount; i++ {
		doTask(in, out)
	}
}

func doTask(in *bufio.Reader, out *bufio.Writer) {
	var cnt, sy, sx int
	fmt.Fscan(in, &cnt, &sy, &sx)
	fmt.Fscanln(in)

	ga := make([]bgreed, 0, cnt)
	ga = append(ga, readGrid(in, sx, sy))
	for i := 0; i < cnt-1; i++ {
		fmt.Fscanln(in)
		ga = append(ga, readGrid(in, sx, sy))
	}
	slices.Reverse(ga)
	gr := ga[0]

	for i := 1; i < cnt; i++ {
		gr = mergeGrid(gr, ga[i], sx, sy)
	}

	printGrid(out, gr)
}

func mergeGrid(fg bgreed, g bgreed, sx int, sy int) []bline {
	for y := 0; y < sy; y++ {
		for x := 0; x < sx; x++ {
			if g[y][x] == '.' {
				continue
			}
			fg[y][x] = g[y][x]
		}
	}
	return fg
}

func readGrid(in *bufio.Reader, sx int, sy int) bgreed {
	grid := bgreed{}
	for i := 0; i < sy; i++ {
		var line bline
		fmt.Fscanln(in, &line)
		grid = append(grid, line)
	}
	return grid
}

func printGrid(out *bufio.Writer, grid bgreed) {
	for _, line := range grid {
		fmt.Fprintln(out, string(line))
	}
	fmt.Fprintln(out)
}
