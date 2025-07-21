package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

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
	var cnt int
	fmt.Fscan(in, &cnt)
	fmt.Fscanln(in)

	counters := map[string]int{}
	curact := ""
	for i := 0; i < cnt; i++ {
		line, _ := in.ReadString('\n')
		subj, obj, act, ens := parceLine(line)

		if _, ok := counters[subj]; !ok {
			counters[subj] = 0
		}
		if _, ok := counters[obj]; !ok {
			counters[obj] = 0
		}

		curact = act

		if subj == obj {
			if ens {
				counters[obj] += 2
			} else {
				counters[obj] -= 1
			}
		} else {
			if ens {
				counters[obj] += 1
			} else {
				counters[obj] -= 1
			}
		}
	}

	top := findMax(counters)
	for _, obj := range top {
		fmt.Fprintf(out, "%s is %s.\n", obj, curact)
	}
}

func findMax(counters map[string]int) []string {
	max := -100000
	obj := []string{}

	for _, count := range counters {
		if count > max {
			max = count
		}
	}

	for key, count := range counters {
		if count == max {
			obj = append(obj, key)
		}
	}

	sort.Strings(obj)
	return obj

}

func parceLine(line string) (subj string, obj string, act string, ens bool) {
	line = strings.TrimRight(line, "!\n")

	parts := strings.Split(line, " ")
	partsCnt := len(parts)

	subj = strings.TrimSuffix(parts[0], ":")
	act = parts[partsCnt-1]
	ens = partsCnt == 4

	if parts[2] == "am" {
		obj = subj
	} else {
		obj = parts[1]
	}

	return
}
