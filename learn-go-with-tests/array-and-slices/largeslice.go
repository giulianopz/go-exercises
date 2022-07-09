package main

import (
	"fmt"
)

func main() {
	a := make([]int, 1e6) // slice "a" with len = 1 million
	b := a[:2]            // even though "b" len = 2, it points to the same the underlying array "a" points to

	c := make([]int, len(b)) // create a copy of the slice so "a" can be garbage collected
	copy(c, b)
	fmt.Println(c)
}
