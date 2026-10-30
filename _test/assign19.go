package main

import "fmt"

func tower(n int) []int {
	s := []int{}
	for i := n; i > 0; i-- {
		s = append(s, i)
	}
	return s
}

func main() {
	a, b := tower(2), []int{}
	c, d := []int{7}, tower(3)
	m, n := map[string]int{"k": 1}, tower(1)
	var e, f = []int{1}, tower(2)
	fmt.Println(a, b, c, d, m, n, e, f)
}

// Output:
// [2 1] [] [7] [3 2 1] map[k:1] [1] [1] [2 1]
