package main

import "fmt"

type node struct {
	val   int
	left  *node
	right *node
}

func traverse(root *node, processVal func(v int)) {
	if root != nil {
		traverse(root.left, processVal)
		processVal(root.val)
		traverse(root.right, processVal)
	}
}

func inorderTraversal(root *node) []int {
	var values []int

	processVal := func(v int) {
		values = append(values, v)
	}

	traverse(root, processVal)

	return values
}

func main() {
	r := &node{1, nil, nil}
	r.right = &node{2, nil, nil}
	r.right.left = &node{3, nil, nil}

	fmt.Println(inorderTraversal(r))
}
