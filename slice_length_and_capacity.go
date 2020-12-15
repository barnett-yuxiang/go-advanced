package main

import "fmt"

func main() {
	trace := make([]int, 3)
	fmt.Printf("%d, %2d, %p, %v\n", len(trace), cap(trace), &trace, trace)

	trace = append(trace, 1, 2, 3, 4)
	fmt.Printf("%d, %2d, %p, %v\n", len(trace), cap(trace), &trace, trace)

	trace = append(trace, 5)
	fmt.Printf("%d, %2d, %p, %v\n", len(trace), cap(trace), &trace, trace)

	trace = append(trace, 6)
	fmt.Printf("%d, %2d, %p, %v\n", len(trace), cap(trace), &trace, trace)
}
