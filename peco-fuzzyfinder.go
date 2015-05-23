package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Sixeight/go-fuzzaldrin"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	buf := bufio.NewReader(os.Stdin)
	candidates := []string{}
	for {
		b, _, err := buf.ReadLine()
		if err != nil {
			break
		}
		candidates = append(candidates, string(b))
	}

	matches := fuzzaldrin.Filter(candidates, os.Args[1], -1)
	for _, line := range matches {
		fmt.Println(line)
	}
}
