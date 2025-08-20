package main

import "fmt"

type Car struct {
	Brand   string
	Year    int
	Electric bool
}

func (c Car) IsVintage() bool {
	return c.Year < 1990
}

func main() {
	c1 := Car{Brand: "Tesla", Year: 2022, Electric: true}
	c2 := Car{Brand: "Toyota", Year: 1985, Electric: false}

	fmt.Printf("%+v Vintage? %v\n", c1, c1.IsVintage())
	fmt.Printf("%+v Vintage? %v\n", c2, c2.IsVintage())
}
