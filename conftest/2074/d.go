package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Event struct {
	x     int
	etype int //etype is 0 -open, 1 - close
}

type ByEvent []Event

func (s ByEvent) Len() int      { return len(s) }
func (s ByEvent) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s Event) String() string {
	return fmt.Sprintf("%d %d", s.x, s.etype)
}

func (s ByEvent) Less(i, j int) bool {
	if s[i].x == s[j].x {
		return s[i].etype < s[j].etype
	}
	return s[i].x < s[j].x
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func solve(x []int, r []int) int {
	//events := make([]Event, len(r)*2)
	//for i := 0; i < len(x); i += 1 {
	//	events[2*i] = Event{x: x[i] - r[i], etype: -1}
	//	events[2*i+1] = Event{x: x[i] + r[i], etype: 1}
	//}
	//sort.Sort(ByEvent(events))
	xp := make(map[int]int)

	ans := 0
	for i := range len(x) {
		for j := x[i] - r[i]; j <= x[i]+r[i]; j += 1 {
			xp[j] = max(xp[j], 2*int(math.Sqrt(float64(r[i]*r[i]-(j-x[i])*(j-x[i]))))+1)
		}
	}
	for x := range xp {
		ans += xp[x]
	}
	return ans
}

func main() {
	in := bufio.NewReader(os.Stdin)
	tc := readInt(in)

	for i := 0; i < tc; i++ {
		nmStr := readLineNumbs(in)
		if len(nmStr) < 2 {
			continue // Safeguard against empty lines
		}

		x := readArrInt(in)

		arr := readArrInt(in)
		ans := solve(x, arr)
		fmt.Println(ans)
	}
}

func readInt(in *bufio.Reader) int {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	n, _ := strconv.Atoi(nStr)
	return n
}

func readLineNumbs(in *bufio.Reader) []string {
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	numbs := strings.Split(line, " ")
	return numbs
}

func readArrInt(in *bufio.Reader) []int {
	numbs := readLineNumbs(in)
	arr := make([]int, len(numbs))
	for i, n := range numbs {
		val, _ := strconv.Atoi(n)
		arr[i] = val
	}
	return arr
}

func readArrInt64(in *bufio.Reader) []int64 {
	numbs := readLineNumbs(in)
	arr := make([]int64, len(numbs))
	for i, n := range numbs {
		val, _ := strconv.ParseInt(n, 10, 64)
		arr[i] = val
	}
	return arr
}
