package main

import "fmt"

func removeAt[T any](s []T, i int) []T {
	return append(s[:i], s[i+1:]...)
}

func reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	names := []string{}
	names = append(names, "Alice", "Bob", "Charlie")
	fmt.Printf("len=%d cap=%d %v\n", len(names), cap(names), names)

	// remove "Bob"
	for i, v := range names {
		if v == "Bob" {
			names = removeAt(names, i)
			break
		}
	}
	fmt.Println("after remove Bob:", names)

	reverse(names)
	fmt.Println("after reverse:", names)
}
