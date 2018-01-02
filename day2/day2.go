package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readFile(s string) [][]int {
	f, err := os.Open(s)
	check(err)
	scanner := bufio.NewScanner(f)
	grid := make([][]int, 0)
	for scanner.Scan() {
		str := scanner.Text()
		split := strings.Fields(str)
		//fmt.Println(split)
		row := make([]int, 0)

		for i, _ := range split {
			//fmt.Println(c)
			num, _ := strconv.Atoi(split[i])
			row = append(row, num)
		}
		grid = append(grid, row)
	}

	f.Close()
	return grid
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	grid := readFile("day2.txt")
	tot := 0

	for _, row := range grid {
		for i, num1 := range row {
			for j, num2 := range row {
				if i != j {
					if num1%num2 == 0 {
						tot += num1 / num2
					}
				}
			}
		}

	}
}
