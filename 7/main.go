package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type File struct {
	name        string
	size        int
	isDirectory bool
}

func isCommand(s string) bool {
	if s[0] == "$"[0] {
		return true
	}
	return false
}

type Path struct {
	parent   *Path
	subs     map[string]*Path
	size     int
	name     string
	isFolder bool
}

func isCD(line string) bool {
	return string(line[2:4]) == "cd"
}

func isLS(line string) bool {
	return string(line[2:4]) == "ls"
}

func cd(p *Path, name string) *Path {
	if p == nil && name == "/" {
		return &Path{name: name, size: 0, parent: nil, subs: map[string]*Path{}, isFolder: true}
	} else if name == ".." {
		return p.parent
	} else {
		_, exists := p.subs[name]
		if !exists {
			p.subs[name] = &Path{name: name, size: 0, parent: p, subs: map[string]*Path{}, isFolder: true}
		}
		return p.subs[name]
	}
}

func parseStdOut(a []string) *Path {

	var pwd *Path = nil
	for _, v := range a {
		line := v
		if isCommand(line) {
			if isCD(line) {
				pwd = cd(pwd, string(line[5:]))
			} else if isLS(line) {
				// ignoring since folder and files will be handles as non command
			} else {
				panic("Unknown command")
			}
		} else {
			addToPath(pwd, line)
		}
	}

	return getRoot(pwd)
}

func getRoot(pwd *Path) *Path {
	if pwd.parent != nil {
		return getRoot(pwd.parent)
	}
	return pwd
}

func toUnsafeInt(s string) int {
	v, _ := strconv.ParseInt(s, 10, 32)
	return int(v)
}

func updateParentSize(p *Path, size int) {
	if p != nil {
		p.size += size
		updateParentSize(p.parent, size)
	}
}

func addToPath(pwd *Path, line string) {

	if string(line[0:3]) == "dir" {

	} else {
		pair := strings.Split(line, " ")
		filename := pair[1]
		size := toUnsafeInt(pair[0])
		_, exists := pwd.subs[filename]
		if !exists {
			pwd.subs[filename] = &Path{name: filename, size: size, parent: pwd, subs: nil, isFolder: false}
			updateParentSize(pwd, size)
		}
	}

}

func sum(p []*Path) int {
	total := 0
	for _, v := range p {
		total += v.size
	}
	return total
}

func sortBySize(files []*Path) []*Path {
	sort.Slice(files, func(i, j int) bool {
		return files[i].size < files[j].size
	})
	return files
}

func findFirstEqGr(sortedFiles []*Path, val int) *Path {
	for _, f := range sortedFiles {
		if f.size >= val {
			return f
		}
	}
	return nil
}

func flatten(p *Path, paths []*Path) []*Path {
	paths = append(paths, p)
	for _, v := range p.subs {
		if v.isFolder {
			paths = flatten(v, paths)
		}
	}

	return paths
}

func collectFoldersLessThen(p *Path, size int, path []*Path) []*Path {
	if p.size < size {
		path = append(path, p)

		for _, v := range p.subs {
			if v.isFolder {
				path = collectFoldersLessThen(v, size, path)
			}
		}
	} else {
		for _, v := range p.subs {
			if v.isFolder {
				path = collectFoldersLessThen(v, size, path)
			}
		}
	}

	return path
}

func createFileTree(path string) *Path {

	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	stdOutputs := []string{}
	for scanner.Scan() {
		stdOutputs = append(stdOutputs, scanner.Text())
	}

	root := parseStdOut(stdOutputs)
	return root
}

func main() {

	root := createFileTree("data/7/input.txt")
	sumOfSmallerFolders := sum(collectFoldersLessThen(root, 100000, []*Path{}))
	fmt.Printf("Part 1 solution %d\n", sumOfSmallerFolders)

	const DEVICE_SIZE = 70000000
	var UPDATE_SIZE = 30000000
	var FREE_SPACE = DEVICE_SIZE - root.size
	var REQUIRED_SPACE int = UPDATE_SIZE - FREE_SPACE

	files := flatten(root, []*Path{})
	dirToRemove := findFirstEqGr(sortBySize(files), REQUIRED_SPACE)
	if dirToRemove != nil {
		fmt.Printf("Part 2 solution %d\n", dirToRemove.size)
	}
}
