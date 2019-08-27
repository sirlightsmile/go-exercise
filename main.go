package main

import (
	"fmt"
	"sync"
)

//
// Go routines
//
//
// This program implements simple binary tree data structure and helper functions to populate it and print it.
// Your goal is to implement a function `goPopulate` in this program. Navigate to `goPopulate` function comment for more details.
// Build and review this program first to understand how it works.
//

type node struct {
	data  int
	left  *node
	right *node
}

// populate binary tree until it reaches height
// every left child contains parent data * 2
// every right child contains (parent data * 2) + 1
//
// example height = 1:
//
//         i
//       /  \
//     2*i  2*i+1
//
func (root *node) populate(height int) {

	if height <= 0 {
		return
	}

	root.left = &node{
		data: root.data * 2,
	}
	root.right = &node{
		data: root.data*2 + 1,
	}

	root.left.populate(height - 1)
	root.right.populate(height - 1)
}

// print binary tree to console
func (root *node) print() {

	var queue = []*node{root}
	var depth = 1

	for len(queue) > 0 {

		var n *node
		n, queue = queue[0], queue[1:]

		fmt.Print(n.data, " ")

		if n.left != nil && n.right != nil {
			queue = append(queue, n.left)
			queue = append(queue, n.right)
			depth++

			if (depth & (depth - 1)) == 0 {
				fmt.Println()
			}
		}
	}

	fmt.Println()
}

// populate binary tree until it reaches height analogically to `populate` function
// restrictions:
//
// - transition from one node to another node (parent -> child) can be done only with new goroutine
// - that means that every child has to be created in a new goroutine
// - for example to populate a tree with 7 nodes, you will need to create at least 6 goroutines (1 is for root)
// - program must be free of race conditions (go run --race main.go)
// - you can write additional helper functions if necessary
//
func (root *node) goPopulate(height int, wg *sync.WaitGroup) {

	if height <= 0 {
		wg.Done()
		return
	}

	root.left = &node{
		data: root.data * 2,
	}
	root.right = &node{
		data: root.data*2 + 1,
	}

	var subWG sync.WaitGroup
	subWG.Add(2)
	go root.left.goPopulate(height-1, &subWG)
	go root.right.goPopulate(height-1, &subWG)
	subWG.Wait()

	wg.Done()
}

func main() {

	tree := node{data: 1}

	//tree.populate(5)

	var mainWG sync.WaitGroup
	mainWG.Add(1)
	go tree.goPopulate(5, &mainWG)
	mainWG.Wait()

	tree.print()
}
