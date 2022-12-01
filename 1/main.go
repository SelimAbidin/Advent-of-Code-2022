package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sum(array []uint32) uint32 {
	var total uint32 = 0
	for _, el := range array {
		total += el
	}
	return total
}

func partOne() []uint32 {

	data, err := os.ReadFile("data/1/part-1.txt")

	if err != nil {
		panic(err)
	}

	var caloryGroups []string = strings.Split(string(data), "\n")
	var totalCals []uint32
	var currentTotal uint32 = 0
	for _, s := range caloryGroups {
		if len(s) == 0 {
			totalCals = append(totalCals, currentTotal)
			currentTotal = 0
		} else {
			i, err := strconv.ParseUint(s, 0, 32)

			if err != nil {
				panic(err)
			}

			currentTotal = currentTotal + uint32(i)
		}
	}

	totalCals = append(totalCals, currentTotal)
	sort.Slice(totalCals, func(i, j int) bool { return totalCals[i] > totalCals[j] })
	fmt.Printf("Part-One Answer is %d\n", totalCals[0])
	fmt.Printf("Part-Two Answer is %d\n", sum(totalCals[0:3]))

	return totalCals
}

func main() {
	partOne()
}
