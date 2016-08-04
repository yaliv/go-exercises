package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int, ends ...*int) {
	// Send value of this point to the channel.
	ch <- t.Value
	// Every tree has at least one end.
	root := 1
	var en *int
	if len(ends) > 0 {
		en = ends[0]
	} else {
		en = &root
	}
	// And we start with no branch.
	br := 0

	// Here we have left branch.
	if t.Left != nil {
		go Walk(t.Left, ch, en)
		br++
	}
	// Here we have right branch.
	if t.Right != nil {
		go Walk(t.Right, ch, en)
		br++
	}

	switch br {
	case 0: // No branch at this point.
		*en--
	case 2: // A branch is added.
		*en++
	}

	// No more branch.
	if *en <= 0 {
		close(ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	// Create two channels.
	ch1, ch2 := make(chan int), make(chan int)
	// Kick off the walkers.
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	// Receive values from Channel 1 to a cache (map),
	// let's call it cv (cached values).
	cv1 := make(map[int]bool)
	for v1 := range ch1 {
		cv1[v1] = true
	}

	// Receive values from Channel 2
	// while compare with Channel 1.
	for v2 := range ch2 {
		if _, ok := cv1[v2]; !ok {
			return false
		}
	}
	// All values from the two channels are SAME.
	return true
}

func main() {
	// Test the `Walk` function.
	fmt.Println("`Walk` Function Test")
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	// Assume we don't know the number of tree values.
	var nums bool
	for v := range ch {
		if nums { fmt.Print(", ") }
		fmt.Print(v)
		nums = true
	}
	fmt.Println()

	// Test the `Same` function.
	fmt.Println("`Same` Function Test")
	k1, k2 := 1, 1
	fmt.Printf( "Tree %v & %v: %v\n", k1, k2, Same(tree.New(k1), tree.New(k2)) )
	k1, k2 = 1, 2
	fmt.Printf( "Tree %v & %v: %v\n", k1, k2, Same(tree.New(k1), tree.New(k2)) )
}
