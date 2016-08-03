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
		br++
		go Walk(t.Left, ch, en)
	}
	// Here we have right branch.
	if t.Right != nil {
		br++
		go Walk(t.Right, ch, en)
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
	return t1 == t2
}

func main() {
	ch := make(chan int)

	// Test the `Walk` function.
	fmt.Println("`Walk` Function Test")
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
	//fmt.Println("`Same` Function Test")
}
