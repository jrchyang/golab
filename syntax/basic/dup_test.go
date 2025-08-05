package basic

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func countLines(f *os.File, counts map[string]int, fileMap map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		fileMap[input.Text()] = f.Name()
	}
}

func Test_dup2(t *testing.T) {
	if len(os.Args) < 6 {
		fmt.Printf("only have %d args for this test, just return", len(os.Args))
		return
	}

	counts := make(map[string]int)
	fileMap := make(map[string]string)
	files := os.Args[5:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, fileMap)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileMap)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, fileMap[line])
		}
	}
}
