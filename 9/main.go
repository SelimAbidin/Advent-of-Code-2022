package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Knot struct {
	x      int
	y      int
	visited []Point 
}

func (k *Knot) setPos(x int, y int)  {
	k.x = x
	k.y = y
	k.visited = append(k.visited, Point{x: k.x, y: k.y})
}


func createKnot(x int, y int) *Knot  {
	return &Knot{x:x, y:y, visited: []Point{{x:x, y:y}}}
}


type Rope struct {
	knots []*Knot
}

func (r* Rope) moveHead (command Command) {

	head := r.knots[0]
	for i := 0; i < command.speed; i++ {
		head.setPos(head.x + command.loc.x, head.y + command.loc.y)
		r.moveNextKnot(1)
	}
}

func (r* Rope) moveNextKnot (currentNodeIndex int) {

	if currentNodeIndex >= len(r.knots) {
		return
	}

	pk := r.knots[currentNodeIndex - 1]
	ck := r.knots[currentNodeIndex]

	d := math.Round(dist(pk, ck))

	if d > 1 {
		
		diffX := maxOne(pk.x - ck.x)
		diffY := maxOne(pk.y - ck.y)
		
		newX := ck.x + diffX
		newY := ck.y + diffY
		
		ck.setPos(newX, newY)
		r.moveNextKnot(currentNodeIndex + 1)
	}
}

func maxOne(i int) int {

	if i == 0 {
		return 0
	}

	return  int(i / int(math.Abs(float64(i))))
}

func dist(p1 *Knot, p2 *Knot) float64 {
	dx := float64(p1.x - p2.x)
	dy := float64(p1.y - p2.y)

	return math.Sqrt(dx * dx + dy * dy)
}

type Point struct {
	x int
	y int
}

type Command struct {
	speed int
	loc Point
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

func applyCommand(rope *Rope, commandStr string) {
	command := parseCommandString(commandStr)
	rope.moveHead(command)
}

func parseCommandString(command string) Command {
	
	cAr := strings.Split(command, " ")
	route := cAr[0]
	speed := toUnsafeInt(cAr[1])

	x := 0
	y:= 0

	if route == "U" {
		y = -1
	} else if route == "D" {
		y = 1
	} else if route == "L" {
		x = -1
	} else if route == "R" {
		x = 1
	}

	return Command{loc: Point{x: x, y:y}, speed: speed}
}

func applyCommands(rope *Rope, commands []string) {
	
	for _, c := range commands {
		applyCommand(rope, c)	
	}
}


func filterDuplicates(point []Point) []Point {
	
	uniques := []Point{}
	uniqueMap := map[string]bool{}
	for _, v := range point {
		key := strconv.Itoa(v.x) + "-" + strconv.Itoa(v.y)
		_, exist := uniqueMap[key]
		if !exist {
			uniqueMap[key] = true
			uniques = append(uniques, v)
		}

	}
	return uniques
}

func printMap(point []Point, actor string) {

	maxX := math.MinInt32
	minX := math.MaxInt32
	maxY := math.MinInt32
	minY := math.MaxInt32

	pointMap := map[string]bool{}
	for _, v := range point {
		key := strconv.Itoa(v.x) + "-" + strconv.Itoa(v.y)
		pointMap[key] = true
	}

	for _, v := range point {
		
		if v.x > maxX {
			maxX = v.x
		}

		if v.x < minX {
			minX = v.x
		}

		if v.y > maxY {
			maxY = v.y
		}
		
		if v.y < minY {
			minY = v.y
		}
	}

	for y := minY; y <= maxY; y++ {

		line := "" 
		for x := minX; x <= maxX; x++ {
			
			key := strconv.Itoa(x) + "-" + strconv.Itoa(y)
			_, exist := pointMap[key]
			if exist {
				line += actor
			} else {
				line += "."
			}

		}

		fmt.Println(line)
	}



}



func main() {
	// commands := readFileByLine("data/9/example.txt")
	// commands2 := readFileByLine("data/9/example-2.txt")
	
	commands := readFileByLine("data/9/input.txt")
	commands2 := commands

	shortRope := &Rope{knots: []*Knot{createKnot(0,0), createKnot(0,0)}}

	applyCommands(shortRope, commands )
	uniquesTailSteps := filterDuplicates(shortRope.knots[1].visited)
	

	longRope := &Rope{knots: []*Knot{}}
	for i := 0; i <= 9; i++ {
		longRope.knots = append(longRope.knots, createKnot(0,0))
	}
	applyCommands(longRope, commands2 )

	flatted := []Point{}
	for _, knot := range longRope.knots {
		flatted = append(flatted, knot.visited...)		
	}

	uniqueSteps := filterDuplicates(longRope.knots[9].visited)

	printMap(uniquesTailSteps, "X")
	fmt.Println("######################################")
	printMap(uniquesTailSteps, "O")

	fmt.Println("Part - 1 solution :", len(uniquesTailSteps)) // 6256
	fmt.Println("Part - 2 solution :", len(uniqueSteps)) // 2665
}

