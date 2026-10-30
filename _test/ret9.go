package main

import "fmt"

func two() ([]int, []int) {
	a, c := []int{1}, []int{7}
	return append(a, 2), append(c, 3)
}

func three() (int, []int, int) {
	s := []int{1, 2, 3}
	return len(s), append(s[:1], 9), cap(s)
}

func main() {
	fmt.Println(two())
	fmt.Println(three())
}

// Output:
// [1 2] [7 3]
// 3 [1 9] 3
