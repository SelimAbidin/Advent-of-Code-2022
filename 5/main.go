package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func toInt(a []string) []int {
	n := []int{}
	for _, v := range a {
		d, _ := strconv.ParseInt(v, 10, 32)
		n = append(n, int(d))
	}
	return n
}

func IfElse[T any](condition bool, a T, b T) T {

	if condition {
		return a
	} else {
		return b
	}

}

func reverse[T any](a []T) []T {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

type StackSet map[int][]string

func (s *StackSet) add(id int, v string) {
	(*s)[id] = append((*s)[id], v)
}

func (s *StackSet) moveFromToByReverse(from int, to int, number int) {
	(*s)[to], (*s)[from] = append((*s)[to], reverse((*s)[from][len((*s)[from])-number:])...), (*s)[from][0:len((*s)[from])-number]
}

func (s *StackSet) moveFromTo(from int, to int, number int) {
	(*s)[to], (*s)[from] = append((*s)[to], (*s)[from][len((*s)[from])-number:]...), (*s)[from][0:len((*s)[from])-number]
}

func (s *StackSet) lastStacks(a []int) string {

	var t string = ""
	for _, v := range a {
		t += (*s)[v][len((*s)[v])-1]
	}

	return t
}

func solution(path string) (string, string) {

	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	createLines := []string{}

	for scanner.Scan() {
		text := scanner.Text()

		if len(text) == 0 {
			break
		}

		createLines = append(createLines, text)
	}

	rex := regexp.MustCompile(`\d+`)
	stacks := toInt(rex.FindAllString(createLines[len(createLines)-1], -1))
	stackSets9000 := StackSet{}
	stackSets9001 := StackSet{}

	for _, v := range reverse(createLines[0 : len(createLines)-1]) {

		for i, sid := range stacks {
			end := (i + 1) * 4
			create := v[i*4 : IfElse(end >= len(v), len(v), end)]
			trimmed := strings.TrimSpace(create)

			if len(trimmed) > 0 {
				stackSets9000.add(sid, string(trimmed[1]))
				stackSets9001.add(sid, string(trimmed[1]))
			}
		}
	}

	for scanner.Scan() {
		text := scanner.Text()
		// move {0} from {1} to {2}
		command := toInt(rex.FindAllString(text, -1))

		stackSets9000.moveFromToByReverse(command[1], command[2], command[0])
		stackSets9001.moveFromTo(command[1], command[2], command[0])
	}

	return stackSets9000.lastStacks(stacks), stackSets9001.lastStacks(stacks)
}

func main() {
	c9000, c9001 := solution("data/5/input.txt")
	fmt.Printf("Part 1 solution is %s\n", c9000)
	fmt.Printf("Part 2 solution is %s\n", c9001)
	// fmt.Printf("Part 2 solution is %d\n", intersectCount)

}
