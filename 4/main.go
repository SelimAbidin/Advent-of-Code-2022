package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type Section struct {
	Start int
	End int
}

func contains(a Section, b Section) bool  {
	return a.Start >= b.Start && a.End <= b.End
}

func intersects(a Section, b Section) bool  {
	return b.Start <= a.End && b.End >= a.End
}

func solution(path string) (int, int) {

	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	works := [][]Section{}
	for scanner.Scan() {
		text := scanner.Text()

		sections := strings.Split(text, ",")

		section1 := strings.Split(sections[0], "-")
		section2 := strings.Split(sections[1], "-")
		
		a1,_ := strconv.ParseInt(section1[0],10, 32)
		a2,_ := strconv.ParseInt(section1[1],10, 32)

		b1,_ := strconv.ParseInt(section2[0],10, 32)
		b2,_ := strconv.ParseInt(section2[1],10, 32)

		a := Section{Start: int(a1), End: int(a2) }
		b := Section{Start: int(b1), End: int(b2) }

		works = append(works, []Section{a, b})
	}

	totalContains :=0 
	for _, v := range works {
		if contains(v[0], v[1]) || contains(v[1], v[0]) {
			totalContains++
		}
	}

	totalIntersects :=0 
	for _, v := range works {
		if intersects(v[0], v[1]) || intersects(v[1], v[0]) {
			totalIntersects++
		}
	}
	
	return totalContains, totalIntersects
}

func main() {
	containsCount, intersectCount  := solution("data/4/input.txt")
	fmt.Printf("Part 1 solution is %d\n", containsCount)
	fmt.Printf("Part 2 solution is %d\n", intersectCount)

}
