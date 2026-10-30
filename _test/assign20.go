package main

import "fmt"

type P struct{ X int }

func id(x int) int { return x }

func main() {
	x, y := 1, 2
	x, y = id(y), id(x)
	fmt.Println(x, y)

	x, y = id(y), x
	fmt.Println(x, y)

	p, q := P{1}, P{2}
	p, q = P{q.X}, P{p.X}
	fmt.Println(p, q)

	s, t := []int{1}, []int{}
	s, t = append(s, 2), s
	fmt.Println(s, t)

	a := []int{1, 2}
	a[0], a[1] = id(a[1]), id(a[0])
	fmt.Println(a)

	ch := make(chan int, 1)
	ch <- 5
	r, k := <-ch, 1
	fmt.Println(r, k)
}

// Output:
// 2 1
// 1 2
// {2} {1}
// [1 2] [1]
// [2 1]
// 5 1
