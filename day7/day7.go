package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

var tower = make([]*floor, 0)

func exists(s string, f *floor) (bool, *floor) {
	if f == nil {
		return false, nil
	}
	if f.name == s {
		return true, f
	}
	for i := range f.above {
		b, fl := exists(s, f.above[i])
		if b {
			return b, fl
		}
	}
	return false, nil
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
		wt, _ := strconv.Atoi(names[1])
		bFl := &floor{names[0], wt, make([]*floor, 0)}

		for i, v := range names {
			if i == 0 {
				b := false
				for j := range tower {
					b, fl := exists(v, tower[j])
					if b {
						fl.weight = wt
						bFl = fl
						fmt.Println(fl.name, "(", fl.weight, ")")
						break
					}
				}
				if !b {
					tower = append(tower, bFl)
				}
			} else if i >= 2 {
				b := false
				for j := range tower {
					b, fl := exists(v, tower[j])
					if b {
						bFl.putOn(fl)
						tower = append(tower[:j], tower[j+1:]...)
						break
					}
				}
				if !b {
					bFl.putOn(&floor{v, 0, make([]*floor, 0)})
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
	for _, v := range tower {
		if v != nil {
			fmt.Print(v.name)
			for _, v2 := range v.above {
				if v2 != nil {
					fmt.Print(" ", v2.name)
				}
			}
			fmt.Println()
		}
	}
	fmt.Println(len(tower))
}
