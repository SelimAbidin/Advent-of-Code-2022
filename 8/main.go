package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Tree struct {
	x      int
	y      int
	size   int
	left   *Tree
	right  *Tree
	bottom *Tree
	top    *Tree
}

func (t *Tree)getVisibilityAndCalculation() (bool, int)  {
	
	visible := false
	if t.isSide() {
		return true, 0
	}

	rDist, rSize := t.rightBiggest()
	lDist, lSize := t.leftBiggest()
	bDist, bSize :=t.bottomBiggest()
	tDist, tSize := t.topBiggest()
	
	if t.size > rSize {
		visible = true
	} else if t.size > lSize {
		visible = true
	}else if t.size > bSize {
		visible = true
	}else if t.size > tSize {
		visible = true
	}

	return visible, rDist *lDist *bDist * tDist
}

func (t *Tree) findThrough(next func(at *Tree) *Tree)(int, int)  {
	
	temp := next(t)

	if temp == nil {
		return 0, t.size
	}

	distance := 0
	for temp != nil {
		distance++
		if t.size <= temp.size{
			break
		}
		temp = next(temp)
	}

	if temp == nil {
		return distance, 0
	}

	return distance, temp.size
}

func (t *Tree)rightBiggest() (int, int)  {
	return t.findThrough(func(at *Tree) *Tree {
		return at.right
	})
}


func (t *Tree)leftBiggest() (int,int)  {
	return t.findThrough(func(at *Tree) *Tree {
		return at.left
	})
}

func (t *Tree)bottomBiggest() (int,int)   {
	return t.findThrough(func(at *Tree) *Tree {
		return at.bottom
	})
}

func (t *Tree)topBiggest() (int,int)  {
	return t.findThrough(func(at *Tree) *Tree {
		return at.top
	})
}

func (t *Tree)isSide() bool  {
	return  t.left == nil || t.right == nil || t.bottom == nil || t.top == nil 
}

func toUnsafeInt(s string) int {
	n, _ := strconv.ParseInt(s, 10, 32)
	return int(n)
}

func readFileByLine(path string) []string {

	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scan := bufio.NewScanner(file)
	data := []string{}
	for scan.Scan() {
		data = append(data, scan.Text())
	}

	return data
}

func toGrid(lines []string) [][]*Tree {

	treeGrid := [][]*Tree{}
	for y, line := range lines {
		temp := []*Tree{}
		for x, r := range line {
			temp = append(temp, &Tree{x: x, y: y, size: toUnsafeInt(string(r))})
		}
		treeGrid = append(treeGrid, temp)
	}

	return treeGrid
}

func setNeighbors(grid [][]*Tree) [][]*Tree {

	for _, line := range grid {

		for _, tree := range line {

			left := tree.x - 1
			top := tree.y - 1
			right := tree.x + 1
			bottom := tree.y + 1

			if left > -1 {
				tree.left = grid[tree.y][left]
			}

			if top > -1 {
				tree.top = grid[top][tree.x]
			}

			if right < len(grid[tree.y]) {
				tree.right = grid[tree.y][right]
			}

			if bottom < len(grid) {
				tree.bottom = grid[bottom][tree.x]
			}
		}

	}

	return grid
}

func main() {
	fileByLine := readFileByLine("data/8/input.txt")
	grid := toGrid(fileByLine)
	setNeighbors(grid)

	highestNum := math.MinInt
	totalVisibleTrees := 0
	for _, line := range grid {
		for _, t := range line {
			visible, scenic := t.getVisibilityAndCalculation()
			if visible {
				totalVisibleTrees++
			}

			if scenic > highestNum {
				highestNum = scenic
			}
		}
	}
	
	fmt.Println("Part - 1 solution :", totalVisibleTrees)
	fmt.Println("Part - 2 solution :", highestNum)
}
