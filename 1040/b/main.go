package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	min, list := solve(n, k)
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
