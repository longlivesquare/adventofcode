package main

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

type floor struct {
	name   string
	weight int
	above  []*floor
}

func (f *floor) putOn(topF *floor) {
	f.above = append(f.above, topF)
}

var tower = make([]*floor, 1288)

func exists(s string) int {
	for i, v := range tower {
		if v.name == s {
			return i
		}
	}
	return -1
}

func readFile(s string) {
	f, err := os.Open(s)

	fun := func(c rune) bool {
		return unicode.IsSpace(c) || c == ',' || c == '(' || c == ')' || c == '-' || c == '>'
	}
	check(err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		str := scanner.Text()
		names := strings.FieldsFunc(str, fun)

		for i, v := range names {
			if i == 0 || i >= 2 {
				fl := exists(v)
				if f1 > 0 {
					if i == 0 {

					}
				}
			}
		}
	}

	f.Close()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	readFile("day7.txt")

}
