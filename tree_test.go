package tree

import (
	"fmt"
	"testing"
)

// TestWalk tests Walk function.
func TestWalk(t *testing.T) {

	tests := []struct {
		name string
		tree []int
		exp  []int
	}{
		{
			name: "empty",
			tree: []int{},
			exp:  []int{},
		},
		{
			name: "multiple nodes",
			tree: []int{8, 13, 3, 5, 1, 1, 2},
			exp:  []int{1, 1, 2, 3, 5, 8, 13},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := make(chan int)
			var nums []int

			go Walk(New(tt.tree...), ch)

			for {
				val, ok := <-ch
				if !ok {
					break
				}
				nums = append(nums, val)
			}

			for i, n := range nums {
				if n != tt.exp[i] {
					t.Errorf("Expected %v, got %v", tt.exp, nums)
					break
				}
			}
		})
	}
}

// ExampleWalk provides Walk function example.
func ExampleWalk() {
	ch := make(chan int)
	go Walk(New(5, -9, 8, 0), ch)

	for {
		val, ok := <-ch
		if !ok {
			break
		}
		fmt.Print(" ", val)
	}
	// Output: -9 0 5 8
}

// TesSame tests Same function.
func TestSame(t *testing.T) {

	tests := []struct {
		name string
		t1   []int
		t2   []int
		exp  bool
	}{
		{
			name: "empty",
			t1:   []int{},
			t2:   []int{},
			exp:  true,
		},
		{
			name: "equal trees",
			t1:   []int{8, 13, 3, 5, 1, 1, 2},
			t2:   []int{8, 3, 1, 1, 2, 5, 13},
			exp:  true,
		},
		{
			name: "different length",
			t1:   []int{8, 13, 3, 5, 1, 1, 2},
			t2:   []int{8, 13, 3, 5, 1},
			exp:  false,
		},
		{
			name: "different node values",
			t1:   []int{8, 13, 3, 5, 1, 1, 2},
			t2:   []int{8, 14, 8, 0, -9, 5, 16},
			exp:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Same(New(tt.t1...), New(tt.t2...))

			if b != tt.exp {
				t.Errorf("Expected %v, got %v", b, tt.exp)
			}
		})
	}
}

// ExampleSame provides Same function example.
func ExampleSame() {
	fmt.Println(Same(New(0, 0, 1, 0), New(0, 0, 1, 1)),
		Same(New(1, 2, 3, 4), New(4, 3, 1, 2)))

	// Output: false true
}
