package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func validPass(s string) bool {
	m := map[string]int{}
	words := strings.Split(s, " ")
	for _, word := range words {
		if m[word] == 1 {
			return false
		}

		for _, ana := range permutations(word) {
			m[ana] = 1
		}
	}
	return true
}

func join(ins []rune, c rune) (result []string) {
	for i := 0; i <= len(ins); i++ {
		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
	}
	return
}

func permutations(testStr string) []string {
	var n func(testStr []rune, p []string) []string
	n = func(testStr []rune, p []string) []string {
		if len(testStr) == 0 {
			return p
		} else {
			result := []string{}
			for _, e := range p {
				result = append(result, join([]rune(e), testStr[0])...)
			}
			return n(testStr[1:], result)
		}
	}

	output := []rune(testStr)
	return n(output[1:], []string{string(output[0])})
}

func readFile(s string) int {
	i := 0
	f, err := os.Open(s)
	check(err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		str := scanner.Text()
		if validPass(str) {
			i++
		}
	}

	f.Close()
	return i
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	n := readFile("day4.txt")
	fmt.Println(n)
}
