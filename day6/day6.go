package main

import "fmt"

type node struct {
	children []*node
	data     int
	visit    int
}

var tree = &node{}

func addSeq(arr []int, time int) (bool, int) {
	curr := tree

	for i, v := range arr {
		found := false
		for j, n := range curr.children {
			if n.data == v {
				if i == len(arr)-1 {
					return false, n.visit
				}
				found = true
				curr.children[j].visit = time
				curr = curr.children[j]
			}
		}
		if !found {
			curr.children = append(curr.children, &node{[]*node{}, v, time})
			curr = curr.children[len(curr.children)-1]
		}
	}
	return true, time
}

func main() {
	arr := []int{4, 10, 4, 1, 8, 4, 9, 14, 5, 1, 14, 15, 0, 15, 3, 5}

	addSeq(arr, 0)

	count := 1
	for true {
		most := 0
		fmt.Println(arr)
		for i, v := range arr {
			if v > arr[most] {
				most = i
			}
		}
		dist := arr[most]
		arr[most] = 0

		most++
		for dist > 0 {
			arr[most%len(arr)]++
			most++
			dist--
		}

		loop, cycles := addSeq(arr, count)
		if !loop {
			fmt.Println(count, count-cycles)
			break
		}
		count++
	}
	fmt.Println(count)
}
