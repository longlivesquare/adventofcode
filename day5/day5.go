package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile(s string) []int {
	arr := make([]int, 0)
	f, err := os.Open(s)
	check(err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		arr = append(arr, i)
	}

	f.Close()
	return arr
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func steps(arr []int) int {
	n := 0
	i := 0
	for true {
		if i < 0 || i >= len(arr) {
			return n
		}
		off := arr[i]
		if arr[i] >= 3 {
			arr[i]--
		} else {
			arr[i]++
		}
		i += off
		n++
	}
	return n
}

func main() {
	arr := readFile("day5.txt")
	//fmt.Println(len(arr))
	fmt.Println(steps(arr))
}
