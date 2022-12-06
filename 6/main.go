package main

import (
	"fmt"
	"os"
)


func containSameChar[T comparable](a []T) bool  {
	
	keys := map[T]bool{}

	for _, v := range a {
		_,exist := keys[v]
		if exist {
			return true
		} else {
			keys[v] = true
		}
	}

	return false
}

func solution(path string, count int) int {

	file, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}


	temp := file[0:count - 1]
	for i, v := range string(file[count - 1:]) {
		
		temp = append(temp, byte(v))
		
		if !containSameChar(temp) {
			return i + count
		}

		temp = temp[1:]
	}
	

	return -1
}

func main() {
	example11 := solution("data/6/example-1-1.txt", 4) // 5
	example12 := solution("data/6/example-1-2.txt", 4) // 6
	example13 := solution("data/6/example-1-3.txt", 4) // 10
	example14 := solution("data/6/example-1-4.txt", 4) // 11
	fmt.Printf("Part 1 Example-1 Solution is %d\n", example11)
	fmt.Printf("Part 1 Example-2 Solution is %d\n", example12)
	fmt.Printf("Part 1 Example-3 Solution is %d\n", example13)
	fmt.Printf("Part 1 Example-4 Solution is %d\n", example14)

	result1 := solution("data/6/input.txt", 4)
	fmt.Printf("Part 1 Result Solution is %d\n", result1)

	fmt.Println("----------------------------------------")

	example21 := solution("data/6/example-2-1.txt", 14) // 19
	example22 := solution("data/6/example-2-2.txt", 14) // 23
	example23 := solution("data/6/example-2-3.txt", 14) // 23
	example24 := solution("data/6/example-2-4.txt", 14) // 29
	example25 := solution("data/6/example-2-5.txt", 14) // 26

	fmt.Printf("Part 2 Example-1 Solution is %d\n", example21)
	fmt.Printf("Part 2 Example-2 Solution is %d\n", example22)
	fmt.Printf("Part 2 Example-3 Solution is %d\n", example23)
	fmt.Printf("Part 2 Example-4 Solution is %d\n", example24)
	fmt.Printf("Part 2 Example-5 Solution is %d\n", example25)


	result2 := solution("data/6/input.txt", 14)
	fmt.Printf("Part 2 Result Solution is %d\n", result2)

}

