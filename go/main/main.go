package main

import (
	"fmt"
)

func main() {

	var test1 = state{0, 4, [64]piece{}}
	x := down(test1)
	fmt.Println(len(x))

}
