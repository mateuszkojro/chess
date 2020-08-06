package main

import (
	"fmt"
)

var allocSize int

func main() {
	allocSize = 16
	var test1 = state{0, 4, [64]piece{}}
	test1 = test1.emptyBoard()
	test1 = test1.set(king{true}, 0, 0)
	test1.show()
	x := up(test1)
	fmt.Println(len(x))

}
