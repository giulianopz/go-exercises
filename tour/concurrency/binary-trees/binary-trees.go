package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {

	Traverse(t, ch)
	close(ch)
}

func Traverse(t *tree.Tree, ch chan int) {

	if t != nil {
		Traverse(t.Left, ch)
		ch <- t.Value
		Traverse(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	c1, c2 := make(chan int), make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)

	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2

		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

func main() {

	c := make(chan int)
	t1 := tree.New(1)
	go Walk(t1, c)
	for i := range c {
		fmt.Println(i)
	}

	fmt.Println("tree.New(1) == tree.New(1):", Same(t1, tree.New(1)))
	fmt.Println("tree.New(1) == tree.New(2):", Same(t1, tree.New(2)))
}
