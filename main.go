package main

import "fmt"

func main() {
	s := []int{1, 2}[2:]
	fmt.Printf("s: %v, len: %d, cap: %d.\n", s, len(s), cap(s))

	foo := make([]int, 5)
	foo[3] = 42
}
