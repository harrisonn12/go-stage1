package main

import (
	"flag"
	"fmt"

	"go-stage1/ex02_tempconv/conv"
)

func main() {
	c := flag.Float64("c", 0, "Temperature in Celsius")
	flag.Parse()

	f := conv.CToF(*c)
	fmt.Printf("%.2f°C = %.2f°F\n", *c, f)
}
