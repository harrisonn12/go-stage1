package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Hello, Go!")

	// var declaration
	var x int          // zero value 0
	var s string       // zero value ""
	var b bool         // zero value false
	var p Person       // zero value fields

	// short declaration
	y := 42
	fmt.Println("y:", y)

	fmt.Printf("Zero values -> int:%d string:%q bool:%v person:%+v\n", x, s, b, p)
}
