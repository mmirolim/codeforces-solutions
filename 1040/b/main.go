package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	min, list := solve1pass(n, k)
	fmt.Printf("%d\n%s", min, strings.Trim(fmt.Sprintf("%v", list), "[]"))
}
func solve(n, k int) (min int, list []int) {
	list = make([]int, 0, n)
	offset := 0
	for {
		for i := k + 1 - offset; i <= n; i += 2*k + 1 {
			min++
			list = append(list, i)
		}
		if list[len(list)-1]+k >= n && len(list) != 0 {
			break
		}
		// reset
		min = 0
		list = list[:0]
		offset++
	}
	return
}

func solve1pass(n, k int) (min int, list []int) {
	segmentSize := 2*k + 1
	numSeg := n / segmentSize
	list = make([]int, 0, numSeg+1)
	diff := n % segmentSize
	offset := 0
	switch {
	case diff == 0:
		min = numSeg
	case diff < k+1:
		offset = k + 1 - diff
		min = numSeg + 1
	case diff >= k+1:
		min = numSeg + 1
	}

	for pos := segmentSize - k - offset; pos <= n; pos += segmentSize {
		list = append(list, pos)
	}

	return min, list
}
