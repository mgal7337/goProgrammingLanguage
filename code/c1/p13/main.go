package main

import (
	"bufio"
	"fmt"
	"os"
)

type lineDetails struct {
	count    int
	occursIn map[string]interface{}
}

func main() {
	counts := make(map[string]*lineDetails)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {

		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n.count > 1 {
			fmt.Printf("Counted: %d\tString: %s\tIn Files: %s\n", n.count, line, n.occursIn)
		}
	}
}

func countLines(f *os.File, counts map[string]*lineDetails) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		nextLine := input.Text()
		if _, ok := counts[nextLine]; !ok {
			counts[nextLine] = &lineDetails{occursIn: make(map[string]interface{})}
		}
		counts[nextLine].count++
		counts[nextLine].occursIn[f.Name()] = struct{}{}
	}
}
