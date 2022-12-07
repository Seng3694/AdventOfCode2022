package main

import (
	"aocutil"
	"fmt"
	"sort"
	"strings"
)

const (
	TYPE_FILE = iota
	TYPE_DIRECTORY
)

type directory_entry struct {
	size      int
	entryType int
	name      string
	parent    *directory_entry
	children  []*directory_entry
}

type context struct {
	current   *directory_entry
	root      *directory_entry
	dirLookup []*directory_entry
}

func initialize_context(ctx *context) {
	ctx.root = &directory_entry{
		size:      0,
		entryType: TYPE_DIRECTORY,
		name:      "/",
		parent:    nil,
		children:  make([]*directory_entry, 0, 8),
	}
	ctx.current = ctx.root
	ctx.dirLookup = make([]*directory_entry, 0, 32)
}

func parse_command(command string, ctx *context) {
	split := strings.Split(command, " ")

	switch split[1] {
	case "cd":
		switch split[2] {
		case "..":
			ctx.current = ctx.current.parent
		case "/":
			ctx.current = ctx.root
		default:
			for _, entry := range ctx.current.children {
				if entry.name == split[2] {
					ctx.current = entry
					break
				}
			}
		}
	case "ls": //ignore
	}
}

func parse_directory(directory string, ctx *context) {
	newEntry := &directory_entry{
		size:      0,
		entryType: TYPE_DIRECTORY,
		name:      directory[4:],
		parent:    ctx.current,
		children:  make([]*directory_entry, 0, 8),
	}
	ctx.current.children = append(ctx.current.children, newEntry)
	ctx.dirLookup = append(ctx.dirLookup, newEntry)
}

func parse_file(file string, ctx *context) {
	split := strings.Split(file, " ")
	newEntry := &directory_entry{
		size:      aocutil.Atoi(split[0]),
		entryType: TYPE_FILE,
		name:      split[1],
		parent:    ctx.current,
		children:  nil,
	}
	ctx.current.children = append(ctx.current.children, newEntry)
}

func parse_line(line string, ctx *context) {
	if strings.HasPrefix(line, "$") {
		parse_command(line, ctx)
	} else {
		if strings.HasPrefix(line, "dir") {
			parse_directory(line, ctx)
		} else {
			parse_file(line, ctx)
		}
	}
}

func get_entry_info(entry *directory_entry) string {
	switch entry.entryType {
	case TYPE_FILE:
		return fmt.Sprintf("file, size=%d", entry.size)
	case TYPE_DIRECTORY:
		return fmt.Sprintf("dir, size=%d", entry.size)
	default:
		return ""
	}
}

func print_tree(current *directory_entry, depth int) {
	fmt.Printf("%s- %s (%s)\n", strings.Repeat(" ", depth*2), current.name, get_entry_info(current))

	for _, c := range current.children {
		print_tree(c, depth+1)
	}
}

func update_directory_sizes(current *directory_entry) int {
	directorySize := 0
	for _, c := range current.children {
		directorySize += update_directory_sizes(c)
	}
	if current.entryType == TYPE_DIRECTORY {
		current.size = directorySize
	}
	return current.size
}

func part1(ctx *context) (solution int) {
	//assumes directories to be sorted by size in descending order
	//iterate from the back (lowest until size is bigger than 100000)
	for i := len(ctx.dirLookup) - 1; i >= 0; i-- {
		if ctx.dirLookup[i].size > 100000 {
			break
		}
		solution += ctx.dirLookup[i].size
	}
	return
}

func part2(ctx *context) (solution int) {
	totalSpace := 70000000
	spaceRequired := 30000000
	usedSpace := ctx.root.size
	spaceToFree := spaceRequired - (totalSpace - usedSpace)
	//assumes directories to be sorted by size in descending order
	for i, d := range ctx.dirLookup {
		if d.size < spaceToFree {
			solution = ctx.dirLookup[i-1].size
			break
		}
	}
	return
}

func main() {
	ctx := context{}
	initialize_context(&ctx)

	aocutil.FileReadAllLines("input.txt", func(s string) {
		parse_line(s, &ctx)
	})

	update_directory_sizes(ctx.root)

	//sort in descending order
	sort.Slice(ctx.dirLookup, func(i, j int) bool {
		return ctx.dirLookup[i].size > ctx.dirLookup[j].size
	})

	aocutil.AOCFinish(part1(&ctx), part2(&ctx))
}
