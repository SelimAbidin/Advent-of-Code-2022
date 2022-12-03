package main

import (
	"bufio"
	"fmt"
	"os"
)

// func findCommonChars (s1 string, s2 string) {

// 	a := map[string]int{}

// 	s1.

// }

type ComparableString string

func (r ComparableString) commonChars(s string) map[string]int {

	a := map[string]int{}

	for _, c1 := range r {
		for _, c2 := range s {
			if c1 == c2 {
				val, exist := a[string(c1)]
				if !exist {
					a[string(c1)] = 1
				} else {
					a[string(c1)] = val + 1
				}

			}
		}
	}

	return a
}

func toScore(i rune) int {
	if i > 90 { // lowercase
		return int(i) - 96
	} else {
		return int(i) - 38
	}
}

func solution(path string) (int, int) {

	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	groupTotal := 0
	group := []string{}
	for scanner.Scan() {
		text := scanner.Text()
		half := len(text) / 2
		commons := ComparableString(text[0:half]).commonChars(text[half:])
		for i := range commons {
			total += toScore(rune(i[0]))
		}

		group = append(group, text)
		if len(group) == 3 {
			com1 := ComparableString(group[0]).commonChars(group[1])
			com2 := ComparableString(group[1]).commonChars(group[2])

			intersects := []rune{}
			for v := range com1 {
				_, exits := com2[v]
				if exits {
					intersects = append(intersects, rune(v[0]))
				}
			}

			if len(intersects) > 1 {
				panic("More than one element in group exists")
			}

			groupTotal += toScore(intersects[0])
			group = []string{}
		}
	}

	return total, groupTotal
}

func main() {
	total, groupTotal := solution("data/3/input.txt")

	fmt.Printf("Part - 1 solution : %d\n", total)
	fmt.Printf("Part - 2 solution : %d\n", groupTotal)

}
