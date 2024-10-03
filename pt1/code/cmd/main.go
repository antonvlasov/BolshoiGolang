package main

import (
	"fmt"
)

func main() {
	x := any(5)
	fmt.Println(x)
	x = any("string")
	fmt.Println(x)
}
