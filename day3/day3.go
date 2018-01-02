package main

import (
	"fmt"
)

type Node struct {
	up    *Node
	down  *Node
	left  *Node
	right *Node
	value int
}

func checkNeighbors(n *Node) {

}

func main() {
	maxleft := 0
	maxright := 0
	maxup := 0
	maxdown := 0

	right := 1
	left := 0
	up := 0
	down := 0

	leftright := 0
	updown := 0

	for i := 1; i < 361527; i++ {
		if right == 1 {
			leftright++
		} else if up == 1 {
			updown++
		} else if left == 1 {
			leftright--
		} else if down == 1 {
			updown--
		}

		if leftright > maxright {
			maxright = leftright
			right = 0
			up = 1
		} else if updown > maxup {
			maxup = updown
			up = 0
			left = 1
		} else if leftright < maxleft {
			maxleft = leftright
			left = 0
			down = 1
		} else if updown < maxdown {
			maxdown = updown
			down = 0
			right = 1
		}

	}

	if leftright < 0 {
		leftright = 0 - leftright
	}
	if updown < 0 {
		updown = 0 - updown
	}

	fmt.Println(leftright + updown)
}
