// Package tree implements equivalence check of binary trees
package tree

import "golang.org/x/tour/tree"

// New returns a sorted binary tree holding values from nums.
func New(nums ...int) *tree.Tree {
	var t *tree.Tree
	for _, n := range nums {
		t = insert(t, n)
	}
	return t
}

func insert(t *tree.Tree, v int) *tree.Tree {
	if t == nil {
		return &tree.Tree{Left: nil, Value: v, Right: nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

// Walk walks the tree t sending all values from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	walkRecursive(t, ch)
}

func walkRecursive(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	walkRecursive(t.Left, ch)
	ch <- t.Value
	walkRecursive(t.Right, ch)
}

// Same determines whether the trees t1 and t2 store the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if ok1 != ok2 || v1 != v2 {
			return false
		}

		if !ok1 {
			break
		}
	}
	return true
}
