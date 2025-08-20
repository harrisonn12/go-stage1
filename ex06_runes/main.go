package main

import (
	"fmt"
	"os"
)

func printRunes(s string) {
	for i, r := range s {
		fmt.Printf("index=%d char=%q codepoint=U+%04X\n", i, r, r)
	}
}

func main() {
	input := "Go 🐹"
	if len(os.Args) > 1 {
		input = os.Args[1]
	}
	printRunes(input)
}
